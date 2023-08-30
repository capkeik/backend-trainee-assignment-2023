package controller

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service/interfaces"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/request"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type UserController struct {
	ctx     context.Context
	service interfaces.UserService
}

func NewUsers(ctx context.Context, service interfaces.UserService) *UserController {
	return &UserController{
		ctx:     ctx,
		service: service,
	}
}

func (c UserController) Get(ctx echo.Context) error {
	userID := ctx.Param("id")

	id64, err := strconv.ParseInt(userID, 10, 32)
	id := int32(id64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user ID"))
	}
	slugs, err := c.service.GetUserSegments(id)
	if err != nil {
		switch {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
	}
	return ctx.JSON(http.StatusOK, response.Slugs{Slugs: *slugs, ID: id})
}

func (c UserController) Create(ctx echo.Context) error {
	var req request.UserReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	createUser, err := c.service.CreateUser(c.ctx, req.ID)
	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "User already exists"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"id": createUser.ID})

}
func (c UserController) UpdateSegments(ctx echo.Context) error {
	var req request.ChangeSegReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}
	changes, err := c.service.UpdateUserSegments(c.ctx, &req.ToAdd, &req.ToRemove, req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, changes)
}
