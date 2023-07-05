package uptimekuma

import (
	"context"
	"encoding/json"
	"github.com/leandrose/uptime-kuma-api-go/config"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities/configs"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/sirupsen/logrus"
	"time"
)

type uptimekumaService struct {
	conn   uptimekuma.IUptimeKumaWebSocket
	ctx    context.Context
	cancel context.CancelFunc
	config configs.UptimeKumaConfig

	info     entities.UptimeKumaInfo
	sid      string
	monitors map[int]entities.Monitor
	avgPing  map[int]int
	uptimes  map[int]map[int]float64

	chanAuth chan bool
}

func NewUptimeKumaService() *uptimekuma.IUptimeKumaService {
	configUptimeKuma := config.GetConfig().UptimeKuma
	ctx, cancel := context.WithCancel(context.Background())

	var instance *uptimekumaService
	instance = &uptimekumaService{
		ctx:    ctx,
		cancel: cancel,
		config: configUptimeKuma,
	}

	var socket *uptimeKumaWebSocket
	socket = &uptimeKumaWebSocket{
		config: configUptimeKuma,
		handles: map[string]uptimekuma.Handle{
			"info":        instance.OnInfo,
			"sid":         instance.OnSid,
			"avgPing":     instance.OnAvgPing,
			"monitorList": instance.OnMonitorList,
			"uptime":      instance.OnUptime,
		},
		cancel: cancel,
		log: logrus.WithFields(logrus.Fields{
			"service": "uptimeKumaWebSocket",
		}),
		chanAck: map[int]chan entities.Event{},
	}
	var socket2 uptimekuma.IUptimeKumaWebSocket = socket
	instance.conn = socket2

	var service uptimekuma.IUptimeKumaService = instance
	socket.Consume(ctx, cancel)
	return &service
}

func (s *uptimekumaService) fill(old interface{}, new interface{}) error {
	b, err := json.Marshal(old)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &new)
	return nil
}

func (s *uptimekumaService) OnInfo(args ...interface{}) {
	err := s.fill(args[0], &s.info)
	logrus.Debugf("OnInfo: %+v", s.info)
	if err != nil {
		logrus.Errorf("OnInfo error: %s", err)
		return
	}

	go s.Auth()
}

func (s *uptimekumaService) OnSid(args ...interface{}) {
	s.sid = args[0].(string)
	logrus.Debugf("OnSid: %+v", s.sid)
}

func (s *uptimekumaService) Auth() bool {
	s.chanAuth = make(chan bool)

	login := []interface{}{
		"login",
		entities.Login{
			Username: s.config.Username,
			Password: s.config.Password,
		},
	}
	b, _ := json.Marshal(login)
	c, _ := s.conn.WriteText(s.ctx, b)

	select {
	case ok := <-c:
		logrus.Debugf("Authenticated: %+v", ok)
		if ok.Token != nil {
			s.conn.SetToken(*ok.Token)
		}
		return ok.Ok
	case <-time.After(5 * time.Second):
		break
	}

	return false
}
