package ports

import (
	"context"
	"inventori-beacukai-backend/internal/module/dashboard/entity"
)

type DashboardRepository interface {
	GetDashboardChart(ctx context.Context, req *entity.GetPenjualanChartReq) (*entity.GetPenjualanChartResp, error)
	GetTotalPenjualan(ctx context.Context, req *entity.GetTotalPenjualanReq) (*entity.GetTotalPenjualanResp, error)
	GetTotalPembelian(ctx context.Context, req *entity.GetTotalPembelianReq) (*entity.GetTotalPembelianResp, error)
	GetTotalWipToday(ctx context.Context, req *entity.GetTotalWipTodayReq) (*entity.GetTotalWipTodayResp, error)
	GetTotalProductMovementNotProcess(ctx context.Context, req *entity.GetTotalProductMovementNotProcessReq) (*entity.GetTotalProductMovementNotProcessResp, error)
}

type DashboardService interface {
	GetDashboardChart(ctx context.Context, req *entity.GetPenjualanChartReq) (*entity.GetPenjualanChartResp, error)
	GetTotalPenjualan(ctx context.Context, req *entity.GetTotalPenjualanReq) (*entity.GetTotalPenjualanResp, error)
	GetTotalPembelian(ctx context.Context, req *entity.GetTotalPembelianReq) (*entity.GetTotalPembelianResp, error)
	GetTotalWipToday(ctx context.Context, req *entity.GetTotalWipTodayReq) (*entity.GetTotalWipTodayResp, error)
	GetTotalProductMovementNotProcess(ctx context.Context, req *entity.GetTotalProductMovementNotProcessReq) (*entity.GetTotalProductMovementNotProcessResp, error)
}
