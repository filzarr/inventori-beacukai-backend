package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/user/entity"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *UserHandler) registerUser(c *fiber.Ctx) error {
	var (
		req = new(entity.RegisterReq)
		v   = adapter.Adapters.Validator
	)
	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Any("req", req.Log()).Msg("handler::register - Invalid request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req.Log()).Msg("handler::register - Invalid request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	resp, err := h.service.RegisterUser(c.Context(), req)
	if err != nil {
		log.Error().Err(err).Any("req", req.Log()).Msg("handler::register - Service error")
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}
