package proto

import (
	"context"
	"github.com/ricejson/apollo-backend/domain"
	toggle2 "github.com/ricejson/apollo-backend/service/toggle"
)

type GRPCToggleServerImpl struct {
	toggleService toggle2.ToggleService
}

func NewGRPCToggleServerImpl(toggleService toggle2.ToggleService) *GRPCToggleServerImpl {
	return &GRPCToggleServerImpl{
		toggleService: toggleService,
	}
}

func (s *GRPCToggleServerImpl) FindAll(ctx context.Context, req *FindAllReq) (*FindAllResp, error) {
	toggles, err := s.toggleService.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	resToggles := make([]*Toggle, len(toggles))
	for i, t := range toggles {
		resToggles[i] = convertToggle(&t)
	}
	return &FindAllResp{
		Toggles: resToggles,
	}, nil
}

func convertToggle(t *domain.Toggle) *Toggle {
	audiences := make([]*Audience, len(t.Audiences))
	for i, a := range t.Audiences {
		audiences[i] = convertAudience(&a)
	}
	return &Toggle{
		Id:          t.Id,
		Key:         t.Key,
		Name:        t.Name,
		Description: t.Description,
		Status:      t.Status,
		CreateAt:    t.CreateAt,
		UpdateAt:    t.UpdateAt,
		Audiences:   audiences,
	}
}

func convertAudience(a *domain.Audience) *Audience {
	rules := make([]*Rule, len(a.Rules))
	for i, r := range a.Rules {
		rules[i] = convertRule(&r)
	}
	return &Audience{
		Id:    a.Id,
		Name:  a.Name,
		Rules: rules,
	}
}

func convertRule(r *domain.Rule) *Rule {
	return &Rule{
		Id:        r.Id,
		Attribute: r.Attribute,
		Operator:  r.Operator,
		Value:     r.Value,
	}
}

func (s *GRPCToggleServerImpl) mustEmbedUnimplementedRPCToggleServiceServer() {

}
