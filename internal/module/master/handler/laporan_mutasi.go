package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *MasterHandler) getLaporanMutasi(c *fiber.Ctx) error {
	var (
		req = new(entity.GetLaporanMutasiReq)
		v   = adapter.Adapters.Validator
	)

	// Parse query params
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getLaporanMutasi - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	// Validasi request
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getLaporanMutasi - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	// Call service
	resp, err := h.service.GetLaporanMutasi(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	// Return success
	return c.JSON(response.Success(resp, ""))
}
