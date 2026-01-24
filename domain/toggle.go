package domain

import "github.com/ricejson/apollo-backend/repository/dao"

// Toggle 开关
type Toggle struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Key         string     `json:"key"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreateAt    int64      `json:"createAt"`
	UpdateAt    int64      `json:"updateAt"`
	Audiences   []Audience `json:"audiences"`
}

func ToggleDao2Domain(toggle dao.Toggle) Toggle {
	audiences := make([]Audience, len(toggle.Audiences))
	for i, audience := range toggle.Audiences {
		audiences[i] = AudienceDao2Domain(audience)
	}
	return Toggle{
		Id:          toggle.Id,
		Name:        toggle.Name,
		Key:         toggle.Key,
		Description: toggle.Description,
		Status:      toggle.Status,
		CreateAt:    toggle.CreateAt,
		UpdateAt:    toggle.UpdateAt,
		Audiences:   audiences,
	}
}

func ToggleDomain2Dao(toggle Toggle) dao.Toggle {
	audiences := make([]dao.Audience, len(toggle.Audiences))
	for i, audience := range toggle.Audiences {
		audiences[i] = AudienceDDomain2Dao(audience)
	}
	return dao.Toggle{
		Id:          toggle.Id,
		Name:        toggle.Name,
		Key:         toggle.Key,
		Description: toggle.Description,
		Status:      toggle.Status,
		CreateAt:    toggle.CreateAt,
		UpdateAt:    toggle.UpdateAt,
		Audiences:   audiences,
	}
}
