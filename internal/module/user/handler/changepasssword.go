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

func (h *UserHandler) changePassword(c *fiber.Ctx) error {
	var (
		req = new(entity.ChangePasswordReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	log.Info().Any("user", c.Get("Authorization")).Msg("fetch user success")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::changePassword - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.UserId = l.GetUserId()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::changePassword - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}
	err := h.service.ChangePassword(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(nil, ""))
}
