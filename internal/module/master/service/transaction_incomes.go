package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetTransactionIncomes(ctx context.Context, req *entity.GetTransactionIncomesReq) (*entity.GetTransactionIncomesResp, error) {
	return s.repo.GetTransactionIncomes(ctx, req)
}

func (s *masterService) GetTransactionIncome(ctx context.Context, req *entity.GetTransactionIncomeReq) (*entity.GetTransactionIncomeResp, error) {
	return s.repo.GetTransactionIncome(ctx, req)
}

func (s *masterService) CreateTransactionIncome(ctx context.Context, req *entity.CreateTransactionIncomeReq) (*entity.CreateTransactionIncomeResp, error) {
	return s.repo.CreateTransactionIncome(ctx, req)
}

func (s *masterService) UpdateTransactionIncome(ctx context.Context, req *entity.UpdateTransactionIncomeReq) error {
	return s.repo.UpdateTransactionIncome(ctx, req)
}

func (s *masterService) DeleteTransactionIncome(ctx context.Context, req *entity.DeleteTransactionIncomeReq) error {
	return s.repo.DeleteTransactionIncome(ctx, req)
}
