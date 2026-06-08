package service

import "context"

type IdentityService struct {
}

func NewIdentityService() *IdentityService {
	return &IdentityService{}
}

func (s *IdentityService) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}
