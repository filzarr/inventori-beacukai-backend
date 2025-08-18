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

func (h *MasterHandler) getLaporanMutasiPemasukan(c *fiber.Ctx) error {
	var (
		req = new(entity.GetLaporanMutasiPemasukanReq)
		v   = adapter.Adapters.Validator
	)

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getLaporanMutasiPemasukan - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()
	log.Info().Msg(req.KodeBarang)

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getLaporanMutasiPemasukan - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetLaporanMutasiPemasukan(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getLaporanWIP(c *fiber.Ctx) error {
	var (
		req = new(entity.GetLaporanMutasiWipReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getLaporanWIP - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getLaporanWIP - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetLaporanMutasiWip(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getLaporanMutasiJenisDokumen(c *fiber.Ctx) error {
	var (
		req = new(entity.GetLaporanMutasiJenisDokumenReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler:: getLaporanMutasiJenisDokumen - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getLaporanMutasiJenisDokumen - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetLaporanMutasiJenisDokumen(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))

}
