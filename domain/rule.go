package domain

import "github.com/ricejson/apollo-backend/repository/dao"

// Rule 规则
type Rule struct {
	Id        string `json:"id"`
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}

func RuleDao2Domain(rule dao.Rule) Rule {
	return Rule{
		Id:        rule.Id,
		Attribute: rule.Attribute,
		Operator:  rule.Operator,
		Value:     rule.Value,
	}
}

func RuleDomain2Dao(rule Rule) dao.Rule {
	return dao.Rule{
		Id:        rule.Id,
		Attribute: rule.Attribute,
		Operator:  rule.Operator,
		Value:     rule.Value,
	}
}
