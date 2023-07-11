package uptimekuma

import (
	"encoding/json"
	"errors"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/sirupsen/logrus"
	"time"
)

func (s *uptimekumaService) OnStatusPageList(args ...interface{}) {
	// 42["statusPageList",{"2":{"id":2,"slug":"f72307ae-53b4-4d1b-a114-f79af92071b3","title":"Descricao informaca pelo Cliente","description":"","icon":"/icon.svg","theme":"light","published":true,"showTags":false,"domainNameList":[],"customCSS":"body {\n  \n}\n","footerText":"","showPoweredBy":false,"googleAnalyticsId":null}}]
	if s.statusPages == nil {
		s.statusPages = map[int]entities.StatusPage{}
	}

	_ = s.fill(args[0], &s.statusPages)
	log := logrus.WithField("func", "OnStatusPageList")
	log.Debugf("count statusPages: %d", len(s.statusPages))
	log.Tracef("statusPages: %+v", s.statusPages)
}

func (s *uptimekumaService) getStatusPage(slug string) (*entities.StatusPage, error) {
	log := logrus.WithField("func", "getStatusPage")
	b, _ := json.Marshal([]interface{}{
		"getStatusPage",
		slug,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Errorf("error in WriteText: %s", err)
		return nil, err
	}

	select {
	case ok := <-c:
		log.Debugf("statuspages: %+v", ok)
		log.Debugf("statuspages Config: %+v", ok.Config)
		if ok.Config != nil {
			s.statusPages[*ok.Config.Id] = *ok.Config
			return ok.Config, nil
		}
		return nil, errors.New(*ok.Msg)
	case <-time.After(5 * time.Second):
		log.Errorf("expired request")
		return nil, errors.New("expired request")
	}
}

func (s *uptimekumaService) GetStatusPages() []entities.StatusPage {
	if s.statusPages == nil {
		return []entities.StatusPage{}
	}

	items := []entities.StatusPage{}
	for _, v := range s.statusPages {
		items = append(items, v)
	}
	return items
}

func (s *uptimekumaService) GetStatusPage(slug string) (*entities.StatusPage, error) {
	if s.statusPages != nil {
		for _, v := range s.statusPages {
			if v.Slug == slug {
				return &v, nil
			}
		}
	}
	return nil, errors.New("statuspage not found")
}

func (s *uptimekumaService) CreateStatusPage(statuspage entities.StatusPage) (*entities.StatusPage, error) {
	//421["addStatusPage","name","slug"]
	//431[{"ok":true,"msg":"OK!"}]
	log := logrus.WithField("func", "CreateStatusPage")
	b, _ := json.Marshal([]interface{}{
		"addStatusPage",
		statuspage.Title,
		statuspage.Slug,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Errorf("error WriteText: %s", err)
		return nil, err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			return s.getStatusPage(statuspage.Slug)
		} else {
			return nil, errors.New(*ok.Msg)
		}
	case <-time.After(5 * time.Second):
		log.Error("expired request")
		return nil, errors.New("expired request")
	}
}

func (s *uptimekumaService) DeleteStatusPage(slug string) error {
	// 422["deleteStatusPage","f72307ae-53b4-4d1b-a114-f79af92071b3"]
	// 432[{"ok":true}]
	log := logrus.WithField("func", "DeleteStatusPage")
	sp, _ := s.getStatusPage(slug)
	b, _ := json.Marshal([]interface{}{
		"deleteStatusPage",
		slug,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Errorf("error WriteText: %s", err)
		return err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			if sp != nil {
				if _, ok2 := s.statusPages[*sp.Id]; ok2 {
					delete(s.statusPages, *sp.Id)
				}
			}
			return nil
		} else {
			return errors.New(*ok.Msg)
		}
	case <-time.After(5 * time.Second):
		log.Error("expired request")
		return errors.New("expired request")
	}
}

func (s *uptimekumaService) UpdateStatusPage(slug string, statusPage entities.StatusPage, monitors []entities.PublicGroupList) (*entities.StatusPage, error) {
	// 422["saveStatusPage","slug",{"id":3,"slug":"slug","title":"name","description":"DescriÃ§Ã£o","icon":"/icon.svg","theme":"light","published":true,"showTags":false,"domainNameList":["example.com","uol.com"],"customCSS":"body {\n  \n}\n","footerText":"Footer Text","showPoweredBy":true,"googleAnalyticsId":"Google Analytics ID"},"/icon.svg",[]]
	// 432[{"ok":true,"publicGroupList":[]}]
	log := logrus.WithField("func", "DeleteStatusPage")
	b, _ := json.Marshal([]interface{}{
		"saveStatusPage",
		slug,
		statusPage,
		statusPage.Icon,
		monitors,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		log.Errorf("error WriteText: %s", err)
		return nil, err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			sp, _ := s.getStatusPage(slug)
			if sp != nil {
				if _, ok2 := s.statusPages[*sp.Id]; ok2 {
					s.statusPages[*sp.Id] = *sp
				}
			}
			return sp, nil
		} else {
			return nil, errors.New(*ok.Msg)
		}
	case <-time.After(5 * time.Second):
		log.Error("expired request")
		return nil, errors.New("expired request")
	}
}
