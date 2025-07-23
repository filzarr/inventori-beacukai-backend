package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	m "inventori-beacukai-backend/internal/middleware"
	"inventori-beacukai-backend/internal/module/user/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *UserHandler) getProfile(c *fiber.Ctx) error {
	var (
		req = new(entity.AuthListenReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.Id = l.GetUserId()
	log.Info().Any("user", req.Id).Msg("fetch user success")

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::changePassword - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}
	resp, err := h.service.GetProfile(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}

func (h *UserHandler) updateProfile(c *fiber.Ctx) error {
	var (
		req = new(entity.UpdateProfileReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.Id = l.GetUserId()
	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::updateProfile - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::updateprofile - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	err := h.service.UpdateProfile(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}
