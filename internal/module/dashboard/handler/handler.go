package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	m "inventori-beacukai-backend/internal/middleware"
	"inventori-beacukai-backend/internal/module/dashboard/entity"
	"inventori-beacukai-backend/internal/module/dashboard/ports"
	"inventori-beacukai-backend/internal/module/dashboard/repository"
	"inventori-beacukai-backend/internal/module/dashboard/service"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type DashboardHandler struct {
	service ports.DashboardService
}

func NewDashboardHandler() *DashboardHandler {
	var (
		repo    = repository.NewDashboardRepository()
		svc     = service.NewDashboardService(repo)
		handler = new(DashboardHandler)
	)
	handler.service = svc

	return handler
}

func (h *DashboardHandler) Register(router fiber.Router) {
	// router.Get("/logs", m.AuthBearer, h.getLogs)
	router.Get("/chart-penjualan", m.AuthBearer, h.getDashboardChart)
	router.Get("/total-penjualan", m.AuthBearer, h.getTotalPenjualan)
	router.Get("/total-pembelian", m.AuthBearer, h.getTotalPembelian)
	router.Get("/total-wip", m.AuthBearer, h.getTotalWipToday)
	router.Get("/total-product-movement-not-progress", m.AuthBearer, h.getTotalProductMovementNotProcess)
	router.Get("/stock-minimum", m.AuthBearer, h.getTotalStockMiminum)
}

func (h *DashboardHandler) getDashboardChart(c *fiber.Ctx) error {
	var (
		req = new(entity.GetPenjualanChartReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getDashboardChart - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getDashboardChart - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetDashboardChart(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *DashboardHandler) getTotalPenjualan(c *fiber.Ctx) error {
	var (
		req = new(entity.GetTotalPenjualanReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getTotalPenjualan - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetTotalPenjualan(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}
func (h *DashboardHandler) getTotalPembelian(c *fiber.Ctx) error {
	var (
		req = new(entity.GetTotalPembelianReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getTotalPembelian - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetTotalPembelian(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *DashboardHandler) getTotalWipToday(c *fiber.Ctx) error {
	var (
		req = new(entity.GetTotalWipTodayReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getTotalWipToday - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetTotalWipToday(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *DashboardHandler) getTotalProductMovementNotProcess(c *fiber.Ctx) error {
	var (
		req = new(entity.GetTotalProductMovementNotProcessReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getTotalProductMovementNotProcess - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetTotalProductMovementNotProcess(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *DashboardHandler) getTotalStockMiminum(c *fiber.Ctx) error {
	var (
		req = new(entity.GetTotalStockMiminumReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getTotalStockMiminum - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetTotalStockMinimum - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetTotalStockMiminum(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}
