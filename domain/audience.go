package domain

import "github.com/ricejson/apollo-backend/repository/dao"

// Audience 人群
type Audience struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}

func AudienceDao2Domain(audience dao.Audience) Audience {
	rules := make([]Rule, len(audience.Rules))
	for i, rule := range audience.Rules {
		rules[i] = RuleDao2Domain(rule)
	}
	return Audience{
		Id:    audience.Id,
		Name:  audience.Name,
		Rules: rules,
	}
}
