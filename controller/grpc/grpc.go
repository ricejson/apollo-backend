package grpc

import (
	"context"
	"errors"
	"github.com/ricejson/apollo-backend/domain"
	toggle2 "github.com/ricejson/apollo-backend/service/toggle"
	"github.com/ricejson/apollo-idl-go/proto"
)

type GRPCToggleServerImpl struct {
	proto.UnimplementedRPCToggleServiceServer // 必须嵌入这个
	toggleService                             toggle2.ToggleService
}

func NewGRPCToggleServerImpl(toggleService toggle2.ToggleService) *GRPCToggleServerImpl {
	return &GRPCToggleServerImpl{
		toggleService: toggleService,
	}
}

func (s *GRPCToggleServerImpl) FindAll(ctx context.Context, req *proto.FindAllReq) (*proto.FindAllResp, error) {
	toggles, err := s.toggleService.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	resToggles := make([]*proto.Toggle, len(toggles))
	for i, t := range toggles {
		resToggles[i] = convertToggle(&t)
	}
	return &proto.FindAllResp{
		Toggles: resToggles,
	}, nil
}

func (s *GRPCToggleServerImpl) InsertOne(ctx context.Context, req *proto.InsertOneReq) (*proto.InsertOneResp, error) {
	if req == nil || req.Toggle == nil {
		return nil, errors.New("invalid request")
	}
	res, err := s.toggleService.InsertToggle(ctx, convert2DomainToggle(req.Toggle))
	if err != nil {
		return nil, err
	}
	return &proto.InsertOneResp{
		Result: res,
	}, nil
}

func convertToggle(t *domain.Toggle) *proto.Toggle {
	audiences := make([]*proto.Audience, len(t.Audiences))
	for i, a := range t.Audiences {
		audiences[i] = convertAudience(&a)
	}
	return &proto.Toggle{
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

func convert2DomainToggle(t *proto.Toggle) domain.Toggle {
	audiences := make([]domain.Audience, len(t.Audiences))
	for i, a := range t.Audiences {
		audiences[i] = convert2DomainAudience(a)
	}
	return domain.Toggle{
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

func convertAudience(a *domain.Audience) *proto.Audience {
	rules := make([]*proto.Rule, len(a.Rules))
	for i, r := range a.Rules {
		rules[i] = convertRule(&r)
	}
	return &proto.Audience{
		Id:    a.Id,
		Name:  a.Name,
		Rules: rules,
	}
}

func convert2DomainAudience(a *proto.Audience) domain.Audience {
	rules := make([]domain.Rule, len(a.Rules))
	for i, r := range a.Rules {
		rules[i] = convert2DomainRule(r)
	}
	return domain.Audience{
		Id:    a.Id,
		Name:  a.Name,
		Rules: rules,
	}
}

func convertRule(r *domain.Rule) *proto.Rule {
	return &proto.Rule{
		Id:        r.Id,
		Attribute: r.Attribute,
		Operator:  r.Operator,
		Value:     r.Value,
	}
}

func convert2DomainRule(r *proto.Rule) domain.Rule {
	return domain.Rule{
		Id:        r.Id,
		Attribute: r.Attribute,
		Operator:  r.Operator,
		Value:     r.Value,
	}
}

func (s *GRPCToggleServerImpl) mustEmbedUnimplementedRPCToggleServiceServer() {

}
