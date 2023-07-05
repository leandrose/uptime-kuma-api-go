package uptimekuma

import (
	"errors"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/sirupsen/logrus"
	"strconv"
)

func (s *uptimekumaService) OnUptime(args ...interface{}) {
	// 42["uptime",6,24,1]
	// 42["uptime",10,720,1]
	if s.uptimes == nil {
		s.uptimes = map[int]map[int]float64{}
	}

	if monitorId, ok := args[0].(string); ok {
		if duration, ok := args[1].(float64); ok {
			if uptime, ok := args[2].(float64); ok {
				monitorIdInt, err := strconv.Atoi(monitorId)
				if err == nil {
					if s.uptimes[monitorIdInt] == nil {
						s.uptimes[monitorIdInt] = map[int]float64{}
					}
					logrus.Debugf("OnUptime: monitorID=%d duration=%d uptime=%f", monitorIdInt, int(duration), uptime)
					s.uptimes[monitorIdInt][int(duration)] = uptime
				}
			}
		}
	}

	logrus.Tracef("OnAvgPing: %+v", s.monitors)
}

func (s *uptimekumaService) GetUptime(monitorID int) (*[]entities.Uptime, error) {
	if _, ok := s.uptimes[monitorID]; ok {
		uptimes := []entities.Uptime{}
		for k, v := range s.uptimes[monitorID] {
			uptimes = append(uptimes, entities.Uptime{
				ID:       monitorID,
				Duration: k,
				Uptime:   v,
			})
		}
		return &uptimes, nil
	}

	return nil, errors.New("uptime not found")
}

func (s *uptimekumaService) GetUptimes() []entities.Uptime {
	if s.uptimes == nil {
		s.uptimes = map[int]map[int]float64{}
	}

	uptimes := []entities.Uptime{}
	for mk, mv := range s.uptimes {
		for dk, dv := range mv {
			uptimes = append(uptimes, entities.Uptime{
				ID:       mk,
				Duration: dk,
				Uptime:   dv,
			})
		}
	}

	return uptimes
}
