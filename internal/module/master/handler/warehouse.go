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

func (h *MasterHandler) getWarehouses(c *fiber.Ctx) error {
	var (
		req = new(entity.GetWarehousesReq)
		v   = adapter.Adapters.Validator
	)
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetWarehousesReq - failed to parse query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetWarehousesReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetWarehouses(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getWarehouse(c *fiber.Ctx) error {
	var (
		req = new(entity.GetWarehouseReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetWarehouseReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetWarehouse(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) createWarehouse(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateWarehouseReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateWarehouseReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::CreateWarehouseReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateWarehouse(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) updateWarehouse(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateWarehouseReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateWarehouseReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::UpdateWarehouseReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.UpdateWarehouse(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}

func (h *MasterHandler) deleteWarehouse(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteWarehouseReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::DeleteWarehouseReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteWarehouse(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::DeleteWarehouseReq - failed to delete warehouse")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return c.JSON(response.Success(nil, ""))
}
