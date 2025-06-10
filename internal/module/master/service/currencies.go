package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetCurrencies(ctx context.Context, req *entity.GetCurrenciesReq) (*entity.GetCurrenciesResp, error) {
	return s.repo.GetCurrencies(ctx, req)
}

func (s *masterService) GetCurrency(ctx context.Context, req *entity.GetCurrencyReq) (*entity.GetCurrencyResp, error) {
	return s.repo.GetCurrency(ctx, req)
}

func (s *masterService) CreateCurrency(ctx context.Context, req *entity.CreateCurrencyReq) (*entity.CreateCurrencyResp, error) {
	return s.repo.CreateCurrency(ctx, req)
}

func (s *masterService) UpdateCurrency(ctx context.Context, req *entity.UpdateCurrencyReq) error {
	return s.repo.UpdateCurrency(ctx, req)
}

func (s *masterService) DeleteCurrency(ctx context.Context, req *entity.DeleteCurrencyReq) error {
	return s.repo.DeleteCurrency(ctx, req)
}
