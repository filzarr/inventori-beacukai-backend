package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	m "inventori-beacukai-backend/internal/middleware"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *MasterHandler) getWarehousesStocks(c *fiber.Ctx) error {
	var (
		req = new(entity.GetWarehousesStocksReq)
		v   = adapter.Adapters.Validator
	)
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetWarehousesStocksReq - failed to parse query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetWarehousesStocksReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetWarehousesStocks(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getWarehousesStock(c *fiber.Ctx) error {
	var (
		req = new(entity.GetWarehousesStockReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetWarehousesStockReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetWarehousesStock(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) createWarehousesStock(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateWarehousesStockReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateWarehousesStockReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::CreateWarehousesStockReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateWarehousesStock(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) updateWarehousesStock(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateWarehousesStockReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateWarehousesStockReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::UpdateWarehousesStockReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.UpdateWarehousesStock(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}

func (h *MasterHandler) deleteWarehousesStock(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteWarehousesStockReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::DeleteWarehousesStockReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteWarehousesStock(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::DeleteWarehousesStockReq - failed to delete warehouses stock")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return c.JSON(response.Success(nil, ""))
}

func (h *MasterHandler) updateStockWarehouse(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateStockWarehousesReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::updateStockWarehouse - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::updateStockWarehouse - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.UpdateStockWarehouses(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}
