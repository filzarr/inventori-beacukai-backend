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

func (h *MasterHandler) getContractsBc(c *fiber.Ctx) error {
	var (
		req = new(entity.GetContractsBcReq)
		v   = adapter.Adapters.Validator
	)
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::GetContractsBcReq - failed to parse query")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetContractsBcReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetContractsBc(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getContractBc(c *fiber.Ctx) error {
	var (
		req = new(entity.GetContractBcReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::GetContractBcReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetContractBc(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) createContractBc(c *fiber.Ctx) error {
	var (
		req = new(entity.CreateContractBcReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::CreateContractBcReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::CreateContractBcReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateContractBc(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) updateContractBc(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateContractBcReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::UpdateContractBcReq - failed to parse body")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::UpdateContractBcReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.UpdateContractBc(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}

func (h *MasterHandler) deleteContractBc(c *fiber.Ctx) error {
	var (
		req = new(entity.DeleteContractBcReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	req.Id = c.Params("id")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::DeleteContractBcReq - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.DeleteContractBc(c.Context(), req)
	if err != nil {
		log.Warn().Err(err).Msg("handler::DeleteContractBcReq - failed to delete contract bc")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return c.JSON(response.Success(nil, ""))
}