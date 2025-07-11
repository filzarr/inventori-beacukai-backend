package handler

import (
	"inventori-beacukai-backend/internal/adapter"
	"inventori-beacukai-backend/internal/module/master/entity"
	errmsg "inventori-beacukai-backend/pkg/errrmsg"
	"inventori-beacukai-backend/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func (h *MasterHandler) getProductsMovement(c *fiber.Ctx) error {
	req := new(entity.GetProductsMovementReq)
	v := adapter.Adapters.Validator

	if err := c.QueryParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::getProductsMovement - parse failed")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.SetDefault()
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getProductsMovement - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProductsMovement(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) getProductsMovementByID(c *fiber.Ctx) error {
	req := &entity.GetProductsMovementReqID{Id: c.Params("id")}
	v := adapter.Adapters.Validator
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::getProductsMovementByID - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.GetProductsMovementByID(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.JSON(response.Success(resp, ""))
}

func (h *MasterHandler) createProductsMovement(c *fiber.Ctx) error {
	req := new(entity.CreateProductsMovementReq)
	v := adapter.Adapters.Validator

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::createProductsMovement - parse failed")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::createProductsMovement - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	resp, err := h.service.CreateProductsMovement(c.Context(), req)
	if err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.Status(fiber.StatusCreated).JSON(response.Success(resp, ""))
}

func (h *MasterHandler) updateProductsMovement(c *fiber.Ctx) error {
	req := new(entity.UpdateProductsMovementReq)
	v := adapter.Adapters.Validator
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::updateProductsMovement - parse failed")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::updateProductsMovement - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := h.service.UpdateProductsMovement(c.Context(), req); err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.JSON(response.Success(nil, ""))
}

func (h *MasterHandler) deleteProductsMovement(c *fiber.Ctx) error {
	req := &entity.DeleteProductsMovementReq{Id: c.Params("id")}
	v := adapter.Adapters.Validator
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::deleteProductsMovement - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := h.service.DeleteProductsMovement(c.Context(), req); err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.JSON(response.Success(nil, ""))
}
func (h *MasterHandler) updateStatusProductsMovement(c *fiber.Ctx) error {
	req := new(entity.UpdateStatusProductsMoveMentReq)
	v := adapter.Adapters.Validator
	req.Id = c.Params("id")

	if err := c.BodyParser(req); err != nil {
		log.Warn().Err(err).Msg("handler::updateProductsMovement - parse failed")
		return c.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := v.Validate(req); err != nil {
		log.Warn().Err(err).Any("req", req).Msg("handler::updateProductsMovement - invalid")
		code, errs := errmsg.Errors(err, req)
		return c.Status(code).JSON(response.Error(errs))
	}

	if err := h.service.UpdateStatusProductsMovement(c.Context(), req); err != nil {
		code, errs := errmsg.Errors[error](err)
		return c.Status(code).JSON(response.Error(errs))
	}
	return c.JSON(response.Success(nil, ""))
}
