package uptimekuma

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/sirupsen/logrus"
	"time"
)

func (s *uptimekumaService) OnMonitorList(args ...interface{}) {
	_ = s.fill(args[0], &s.monitors)
	logrus.Debugf("OnMonitorList count monitors: %d", len(s.monitors))
	logrus.Tracef("OnMonitorList: %+v", s.monitors)
}

func (s *uptimekumaService) GetMonitors() []entities.Monitor {
	monitors := []entities.Monitor{}
	for _, m := range s.monitors {
		monitors = append(monitors, m)
	}

	return monitors
}

func (s *uptimekumaService) GetMonitorById(id int) (*entities.Monitor, error) {
	if _, ok := s.monitors[id]; ok {
		m := s.monitors[id]
		return &m, nil
	}
	return nil, errors.New("monitor not found")
}

func (s *uptimekumaService) CreateMonitor(monitor entities.Monitor) (*entities.Monitor, error) {
	// 422["add",{"type":"http","name":"Teste","url":"https://www.uol.com.br","method":"GET","interval":60,"retryInterval":60,"resendInterval":0,"maxretries":1,"notificationIDList":{},"ignoreTls":false,"upsideDown":false,"packetSize":56,"expiryNotification":false,"maxredirects":10,"accepted_statuscodes":["200-299"],"dns_resolve_type":"A","dns_resolve_server":"1.1.1.1","docker_container":"","docker_host":null,"proxyId":null,"mqttUsername":"","mqttPassword":"","mqttTopic":"","mqttSuccessMessage":"","authMethod":null,"httpBodyEncoding":"json"}]
	// 432[{"ok":true,"msg":"Added Successfully.","monitorID":8}]
	log := logrus.WithField("func", "CreateMonitor")
	m := []interface{}{
		"add",
		monitor,
	}
	log.Debugf("monitor=%+v", monitor)
	b, _ := json.Marshal(m)
	c, _ := s.conn.WriteText(s.ctx, b)

	select {
	case ok := <-c:
		log.Debugf("success: %+v", ok)
		if !ok.Ok {
			return nil, errors.New(*ok.Msg)
		} else if ok.Ok && *ok.MonitorId > 0 {
			return s.GetMonitorById(*ok.MonitorId)
		}
	case <-time.After(10 * time.Second):
		return nil, errors.New("there was no answer")
	}

	return nil, errors.New("error occurred")
}

func (s *uptimekumaService) DeleteMonitor(monitorId int) error {
	// 4218["deleteMonitor",7]
	// 4318[{"ok":true,"msg":"Deleted Successfully."}]
	log := logrus.WithField("func", "DeleteMonitor")
	m := []interface{}{
		"deleteMonitor",
		monitorId,
	}
	log.Debugf("monitorId=%d", monitorId)
	b, _ := json.Marshal(m)
	c, _ := s.conn.WriteText(s.ctx, b)

	select {
	case ok := <-c:
		if !ok.Ok {
			return errors.New(fmt.Sprintf("failed on delete monitor: %s", ok.Msg))
		}
	case <-time.After(5 * time.Second):
		return errors.New("there was no answer")
	}

	if _, ok := s.monitors[monitorId]; ok {
		delete(s.monitors, monitorId)
	}
	if _, ok := s.avgPing[monitorId]; ok {
		delete(s.avgPing, monitorId)
	}
	if _, ok := s.uptimes[monitorId]; ok {
		delete(s.uptimes, monitorId)
	}
	return nil
}

func (s *uptimekumaService) EditMonitor(monitor entities.Monitor) (*entities.Monitor, error) {
	// 4222["editMonitor",{"id":6,"name":"web.1.receitanet.net","description":null,"url":null,"method":"GET","hostname":"web.1.receitanet.net","port":53,"maxretries":1,"weight":2000,"active":1,"type":"ping","interval":60,"retryInterval":60,"resendInterval":0,"keyword":null,"expiryNotification":false,"ignoreTls":false,"upsideDown":false,"packetSize":56,"maxredirects":10,"accepted_statuscodes":["200-299"],"dns_resolve_type":"A","dns_resolve_server":"1.1.1.1","dns_last_result":null,"docker_container":"","docker_host":null,"proxyId":null,"notificationIDList":{},"tags":[],"maintenance":false,"mqttTopic":null,"mqttSuccessMessage":null,"databaseQuery":null,"authMethod":"","grpcUrl":null,"grpcProtobuf":null,"grpcMethod":null,"grpcServiceName":null,"grpcEnableTls":false,"radiusCalledStationId":null,"radiusCallingStationId":null,"game":null,"httpBodyEncoding":null,"headers":null,"body":null,"grpcBody":null,"grpcMetadata":null,"basic_auth_user":null,"basic_auth_pass":null,"pushToken":null,"databaseConnectionString":null,"radiusUsername":null,"radiusPassword":null,"radiusSecret":null,"mqttUsername":null,"mqttPassword":null,"authWorkstation":null,"authDomain":null,"tlsCa":null,"tlsCert":null,"tlsKey":null,"includeSensitiveData":true}]
	// 4322[{"ok":true,"msg":"Saved.","monitorID":6}]
	log := logrus.WithField("func", "DeleteMonitor")
	m := []interface{}{
		"editMonitor",
		monitor,
	}
	log.Debugf("monitor=%+v", monitor)
	b, _ := json.Marshal(m)
	c, _ := s.conn.WriteText(s.ctx, b)

	var mm *entities.Monitor
	select {
	case ok := <-c:
		if !ok.Ok {
			return nil, errors.New(fmt.Sprintf("failed on edit monitor: %s", ok.Msg))
		}
		if *ok.MonitorId > 0 {
			return s.GetMonitorById(*ok.MonitorId)
		}
	case <-time.After(5 * time.Second):
		return nil, errors.New("there was no answer")
	}

	if mm != nil {
		return nil, errors.New("monitor not found")
	}
	return mm, nil
}
