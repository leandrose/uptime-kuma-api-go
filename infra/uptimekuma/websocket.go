package uptimekuma

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/sirupsen/logrus"
	"nhooyr.io/websocket"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

var uptimeKumaWebServiceInstance *uptimekuma.IUptimeKumaWebSocket

const (
	PROTOCOL_OPEN    int = 0
	PROTOCOL_CLOSE   int = 1
	PROTOCOL_PING    int = 2
	PROTOCOL_PONG    int = 3
	PROTOCOL_MESSAGE int = 4
	PROTOCOL_UPGRADE int = 5
	PROTOCOL_NOOP    int = 6

	EXCHANGE_CONNECT       int = 0
	EXCHANGE_DISCONNECT    int = 1
	EXCHANGE_EVENT         int = 2
	EXCHANGE_ACK           int = 3
	EXCHANGE_CONNECT_ERROR int = 4
	EXCHANGE_BINARY_EVENT  int = 5
	EXCHANGE_BINARY_ACK    int = 6
)

type uptimeKumaWebSocket struct {
	conn            *websocket.Conn
	cancel          context.CancelFunc
	config          configs.UptimeKumaConfig
	handles         map[string]uptimekuma.Handle
	isAuthenticated bool
	isClosed        bool
	token           string
	messageCount    int
	log             *logrus.Entry

	chanAck map[int]chan entities.Event
}

func (c uptimeKumaWebSocket) IsAuthenticated() bool {
	return c.IsAuthenticated()
}

func (s *uptimeKumaWebSocket) Send() {

}

func (s *uptimeKumaWebSocket) Close() {
	s.isClosed = true
	_ = s.conn.Close(websocket.StatusNormalClosure, "finished")
}

func (s *uptimeKumaWebSocket) SetToken(token string) {
	s.token = token
	s.isAuthenticated = true
}

func (s *uptimeKumaWebSocket) Consume(ctx context.Context, cancel context.CancelFunc) {
	s.cancel = cancel

	var err error
	var newConn *websocket.Conn
	for {
		newConn, _, err = websocket.Dial(ctx, s.config.Uri, nil)
		if err == nil {
			break
		}

		s.log.Infof("Failed to reconnect: %s", err)
		time.Sleep(5 * time.Second)
	}
	s.conn = newConn
	s.conn.SetReadLimit(1024 * 1024 * 2)

	go s.run(ctx, cancel)
}

func (s *uptimeKumaWebSocket) run(ctx context.Context, cancelFunc context.CancelFunc) {
	for {
		msg, b, err := s.conn.Read(ctx)
		if err != nil {
			s.log.Errorf("Connection closed: %s %+v %+v", err, msg, b)
			break
		}
		s.log.Tracef("READ msg: %s b: %s", msg, string(b))

		// PROCESSAR
		bb, err := strconv.Atoi(string(b[0]))
		switch bb {
		case PROTOCOL_OPEN:
			// CONNECT
			_ = s.Write(ctx, websocket.MessageText, []byte("40"))
		case PROTOCOL_PING:
			// EVENT
			s.log.Tracef("PONG")
			_ = s.Write(ctx, websocket.MessageText, []byte("3"))
		case PROTOCOL_MESSAGE:
			str := string(b)

			re := regexp.MustCompile("^4(?P<EXCHANGE>[0-9]{0,1})(?P<SEQUENCE>[0-9]*)(?P<JSON>[[{].*)$")
			match := re.FindStringSubmatch(str)

			result := make(map[string]string)
			for i, name := range re.SubexpNames() {
				if i != 0 && name != "" {
					result[name] = match[i]
				}
			}

			// CONNECT
			exchange, _ := strconv.Atoi(result["EXCHANGE"])
			switch exchange {
			case EXCHANGE_CONNECT:
				var m map[string]interface{}
				err := json.Unmarshal([]byte(result["JSON"]), &m)
				if err == nil {
					k := reflect.ValueOf(m).MapKeys()
					if len(k) > 0 {
						kk := k[0].String()
						if _, ok := s.handles[kk]; ok {
							s.handles[kk](m[kk].(string))
						}
					}
				}
			case EXCHANGE_EVENT:
				var m []interface{}
				err := json.Unmarshal([]byte(result["JSON"]), &m)
				if err == nil {
					if reflect.TypeOf(m[0]).Kind() == reflect.String {
						mm := m[1:]
						if s.handles[m[0].(string)] != nil {
							s.handles[m[0].(string)](mm...)
						}
					}
				}
			case EXCHANGE_ACK:
				sequence, _ := strconv.Atoi(result["SEQUENCE"])
				if i, ok := s.chanAck[sequence]; ok {
					var event []entities.Event
					_ = json.Unmarshal([]byte(result["JSON"]), &event)
					i <- event[0]
				}
			}
		}
	}

	s.log.Tracef("Exited run")
	if !s.isClosed {
		go s.Consume(ctx, cancelFunc)
	}
}

func (s *uptimeKumaWebSocket) On(event string, handleFunc uptimekuma.Handle) {
	s.handles[event] = handleFunc
}

func (s *uptimeKumaWebSocket) WriteText(ctx context.Context, b []byte) (chan entities.Event, error) {
	id := s.messageCount
	s.messageCount = s.messageCount + 1
	s.chanAck[id] = make(chan entities.Event)

	b = append([]byte(fmt.Sprintf("%d%d%d", PROTOCOL_MESSAGE, EXCHANGE_EVENT, id)), b...)
	s.log.Debugf("Socket WriteText: %s", string(b))
	err := s.conn.Write(ctx, websocket.MessageText, b)
	return s.chanAck[id], err
}

func (s *uptimeKumaWebSocket) Write(ctx context.Context, messageType websocket.MessageType, b []byte) (err error) {
	s.log.Debugf("Socket Write: %s", string(b))
	return s.conn.Write(ctx, messageType, b)
}
