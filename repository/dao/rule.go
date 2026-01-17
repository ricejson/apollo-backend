package dao

// Rule 规则
type Rule struct {
	Id        string `json:"id"`
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value"`
}
