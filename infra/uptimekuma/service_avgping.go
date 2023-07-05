package uptimekuma

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (s *uptimekumaService) OnAvgPing(args ...interface{}) {
	// 42["avgPing",8,235]
	if s.avgPing == nil {
		s.avgPing = map[int]int{}
	}

	if monitorId, ok := args[0].(string); ok {
		if avgPing, ok := args[1].(float64); ok {
			monitorIdInt, err := strconv.Atoi(monitorId)
			if err == nil {
				logrus.Debugf("OnAvgPing: monitorID=%d avgPing=%d", monitorIdInt, int(avgPing))
				s.avgPing[monitorIdInt] = int(avgPing)
			}
		}
	}

	logrus.Tracef("OnAvgPing: %+v", s.monitors)
}

func (s *uptimekumaService) GetAvgPing(monitorID int) (*int, error) {
	if i, ok := s.avgPing[monitorID]; ok {
		return &i, nil
	}

	return nil, errors.New("avg ping not found")
}

func (s *uptimekumaService) GetAvgPings() map[int]int {
	if s.avgPing == nil {
		s.avgPing = map[int]int{}
	}

	return s.avgPing
}
