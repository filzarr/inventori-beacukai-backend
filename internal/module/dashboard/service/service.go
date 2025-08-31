package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/dashboard/entity"
	"inventori-beacukai-backend/internal/module/dashboard/ports"
)

var _ ports.DashboardService = &dashboardService{}

type dashboardService struct {
	repo ports.DashboardService
}

func NewDashboardService(repo ports.DashboardRepository) *dashboardService {
	return &dashboardService{
		repo: repo,
	}
}

func (s *dashboardService) GetDashboardChart(ctx context.Context, req *entity.GetPenjualanChartReq) (*entity.GetPenjualanChartResp, error) {
	return s.repo.GetDashboardChart(ctx, req)
}

func (s *dashboardService) GetTotalPenjualan(ctx context.Context, req *entity.GetTotalPenjualanReq) (*entity.GetTotalPenjualanResp, error) {
	return s.repo.GetTotalPenjualan(ctx, req)
}
func (s *dashboardService) GetTotalPembelian(ctx context.Context, req *entity.GetTotalPembelianReq) (*entity.GetTotalPembelianResp, error) {
	return s.repo.GetTotalPembelian(ctx, req)
}
func (s *dashboardService) GetTotalWipToday(ctx context.Context, req *entity.GetTotalWipTodayReq) (*entity.GetTotalWipTodayResp, error) {
	return s.repo.GetTotalWipToday(ctx, req)
}
func (s *dashboardService) GetTotalProductMovementNotProcess(ctx context.Context, req *entity.GetTotalProductMovementNotProcessReq) (*entity.GetTotalProductMovementNotProcessResp, error) {
	return s.repo.GetTotalProductMovementNotProcess(ctx, req)
}
func (s *dashboardService) GetTotalStockMiminum(ctx context.Context, req *entity.GetTotalStockMiminumReq) (*entity.GetTotalStockMiminumResp, error) {
	return s.repo.GetTotalStockMiminum(ctx, req)
}
