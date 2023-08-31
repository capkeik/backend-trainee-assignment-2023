package controller

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service/interfaces"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/request"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type SegmentController struct {
	ctx      context.Context
	service  interfaces.SegmentService
	validate *validator.Validate
}

func NewSegments(ctx context.Context, service interfaces.SegmentService) *SegmentController {
	return &SegmentController{
		ctx:      ctx,
		service:  service,
		validate: validator.New(),
	}
}

// Create TODO: add error handling: Invalid slug; Too long slug; Slug Already Exists; Internal Error;
func (c SegmentController) Create(ctx echo.Context) error {
	var req request.SegmentReq

	if err := ctx.Bind(&req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	if err := c.validate.Struct(req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}
	slug := req.Slug

	_, err := c.service.Create(slug)
	if err != nil {
		log.Println("err while creating new slug:", err.Error())
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "Slug already exists"})
	}
	log.Println("SegmentController:", "creating new slug:", req.Slug)
	return ctx.JSON(http.StatusOK, map[string]interface{}{"slug": slug})
}

func (c SegmentController) Delete(ctx echo.Context) error {
	var req request.SegmentReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	if err := c.validate.Struct(req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	slug := req.Slug
	log.Println("SegmentController:", "Deleting slug:", req.Slug)
	err := c.service.Delete(slug)
	if err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "InternalServerError"})
	}

	return ctx.NoContent(http.StatusOK)
}
