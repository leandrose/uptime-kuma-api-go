package uptimekuma

import (
	"encoding/json"
	"errors"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"time"
)

func (s *uptimekumaService) getTags() (*[]entities.Tag, error) {
	// 427["getTags"]
	// 437[{"ok":true,"tags":[{"id":1,"name":"Teste","color":"#DC2626"}]}]
	b, _ := json.Marshal([]interface{}{
		"getTags",
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		return nil, err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			if ok.Tags != nil {
				return ok.Tags, nil
			}
		}
	case <-time.After(5 * time.Second):
		return nil, errors.New("expired requests")
	}

	return nil, errors.New("error occurred")
}

func (s *uptimekumaService) GetTag(tagID int) (*entities.Tag, error) {
	tags, err := s.getTags()
	if err != nil {
		return nil, err
	}

	for _, v := range *tags {
		if *v.ID == tagID {
			return &v, nil
		}
	}

	return nil, errors.New("tag not found")
}

func (s *uptimekumaService) GetTags() (*[]entities.Tag, error) {
	return s.getTags()
}

func (s *uptimekumaService) CreateTag(tag entities.Tag) (*entities.Tag, error) {
	// 4212["addTag",{"id":null,"name":"sdf","color":"#DC2626"}]
	// 4312[{"ok":true,"tag":{"id":2,"name":"sdf","color":"#DC2626"}}]
	b, _ := json.Marshal([]interface{}{
		"addTag",
		tag,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		return nil, err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			if ok.Tag != nil {
				return ok.Tag, nil
			}
		}
	case <-time.After(5 * time.Second):
		return nil, errors.New("expired request")
	}

	return nil, errors.New("error occurred")
}

func (s *uptimekumaService) UpdateTag(tag entities.Tag) (*entities.Tag, error) {
	// 4215["editTag",{"id":2,"name":"sdf","color":"#DC2626"}]
	// 4315[{"ok":true,"msg":"Saved","tag":{"id":2,"name":"sdf","color":"#DC2626"}}]
	b, _ := json.Marshal([]interface{}{
		"editTag",
		tag,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		return nil, err
	}

	select {
	case ok := <-c:
		if ok.Ok && ok.Tag != nil {
			return ok.Tag, nil
		}
	case <-time.After(5 * time.Second):
		return nil, errors.New("expired request")
	}

	return nil, errors.New("error occurred")
}

func (s *uptimekumaService) DeleteTag(tabID int) error {
	// 4218["deleteTag",2]
	// 4318[{"ok":true,"msg":"Deleted Successfully."}]
	b, _ := json.Marshal([]interface{}{
		"deleteTag",
		tabID,
	})
	c, err := s.conn.WriteText(s.ctx, b)
	if err != nil {
		return err
	}

	select {
	case ok := <-c:
		if ok.Ok {
			return nil
		}
	case <-time.After(5 * time.Second):
		return errors.New("expired request")
	}

	return errors.New("error occurred")
}
