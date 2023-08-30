package controller

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service/interfaces"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/request"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type SegmentController struct {
	ctx     context.Context
	service interfaces.SegmentService
}

func NewSegments(ctx context.Context, service interfaces.SegmentService) *SegmentController {
	return &SegmentController{
		ctx:     ctx,
		service: service,
	}
}

// Create TODO: add error handling: Invalid slug; Too long slug; Slug Already Exists; Internal Error;
func (c SegmentController) Create(ctx echo.Context) error {
	var segmentReq request.SegmentReq

	if err := ctx.Bind(&segmentReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}
	log.Println("creating new slug:", segmentReq.Slug)
	slug := segmentReq.Slug

	_, err := c.service.Create(slug)
	if err != nil {
		log.Println("err while creating new slug:", err.Error())
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "Slug already exists"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"slug": slug})
}

func (c SegmentController) Delete(ctx echo.Context) error {
	var segmentReq request.SegmentReq

	if err := ctx.Bind(&segmentReq); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	slug := segmentReq.Slug
	log.Println("deleting slug:", segmentReq.Slug)
	err := c.service.Delete(slug)
	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": err.Error()})
	}

	return ctx.NoContent(http.StatusOK)
}
