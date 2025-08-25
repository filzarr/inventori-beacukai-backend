package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	m "inventori-beacukai-backend/internal/middleware"
	"inventori-beacukai-backend/internal/module/log/entity"
	"inventori-beacukai-backend/internal/module/log/ports"
	"inventori-beacukai-backend/internal/module/log/repository"
	"inventori-beacukai-backend/internal/module/log/service"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type LogHandler struct {
	service ports.LogService
}

func NewLogHandler() *LogHandler {
	var (
		repo    = repository.NewLogRepository()
		svc     = service.NewLogService(repo)
		handler = new(LogHandler)
	)
	handler.service = svc

	return handler
}

func (h *LogHandler) Register(router fiber.Router) {
	router.Get("/logs", m.AuthBearer, h.getLogs)
}

func (h *LogHandler) getLogs(c *fiber.Ctx) error {
	var (
		req = new(entity.GetLogReq)
		v   = adapter.Adapters.Validator
		l   = m.GetLocals(c)
	)
	req.UserId = l.GetUserId()
	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getLog - failed to parse request")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}

	req.SetDefault()

	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getLogs - invalid request")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetLogs(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}

	return c.JSON(response.Success(resp, ""))
}
