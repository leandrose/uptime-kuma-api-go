package uptimekuma

import (
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

func (s *uptimekumaService) OnHeartbeatList(args ...interface{}) {
	// 42["heartbeatList","3",[{"id":56004,"important":0,"monitor_id":3,"status":1,"msg":"200 - OK","time":"2023-07-06 16:06:03.734","ping":1107,"duration":61,"down_count":0}],false]
	log := logrus.WithField("func", "OnHeartbeatList")
	if s.heartbeatList == nil {
		s.heartbeatList = map[int][]entities.Heartbeat{}
	}

	monitorID, err := strconv.Atoi(args[0].(string))
	if err != nil {
		log.Errorf("error: %s", err)
		return
	}

	v := reflect.ValueOf(args[1])
	if v.Kind() == reflect.Slice {
		var result []entities.Heartbeat
		for i := 0; i < v.Len(); i++ {
			n := entities.Heartbeat{}
			_ = s.fill(v.Index(i).Interface(), &n)
			result = append(result, n)
		}
		s.heartbeatList[monitorID] = result
	}

	log.Debugf("count heartbeatList(%d): %d", monitorID, len(s.heartbeatList[monitorID]))
	log.Tracef("heartbeatList(%d): %+v", monitorID, s.heartbeatList[monitorID])
}

func (s *uptimekumaService) GetMonitorHeartbets(monitorID int) []entities.Heartbeat {
	if s.heartbeatList != nil {
		if _, ok := s.heartbeatList[monitorID]; ok {
			return s.heartbeatList[monitorID]
		}
	}
	return []entities.Heartbeat{}
}
