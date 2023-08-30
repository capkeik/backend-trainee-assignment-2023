package controller

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/repository/pg"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

// ChangeSegReq TODO Move to internal/web/request
type ChangeSegReq struct {
	ToAdd    []string `json:"to_add"`
	ToRemove []string `json:"to_remove"`
	Id       int32    `json:"id"`
}

// UserReq TODO Move to internal/web/request
type UserReq struct {
	ID int32 `json:"id"`
}

type SlugsResp struct {
	ID    int32    `json:"id"`
	Slugs []string `json:"slugs"`
}
type UserController struct {
	ctx context.Context
	rep *pg.UserRepo
}

func NewUsers(ctx context.Context, rep *pg.UserRepo) *UserController {
	return &UserController{
		ctx: ctx,
		rep: rep,
	}
}

func (c UserController) Get(ctx echo.Context) error {
	userID := ctx.Param("id")

	id, err := strconv.ParseInt(userID, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "could not parse user ID"))
	}
	user, err := c.rep.GetUser(int32(id))
	slugs, err := c.rep.GetUserSegments(int32(id))
	if err != nil {
		switch {
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, errors.Wrap(err, "could not get user"))
		}
	}
	return ctx.JSON(http.StatusOK, SlugsResp{Slugs: *slugs, ID: user.ID})
}

func (c UserController) Create(ctx echo.Context) error {
	var req UserReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	createUser, err := c.rep.CreateUser(c.ctx, req.ID)
	if err != nil {
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "User already exists"})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"id": createUser.ID})

}
func (c UserController) UpdateSegments(ctx echo.Context) error {
	var req ChangeSegReq

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	user, err := c.rep.GetUser(req.Id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": "User not found"})
	}

	var oldSeg []string

	for _, s := range user.Segments {
		oldSeg = append(oldSeg, s.Slug)
	}

	segments, err := c.rep.UpdateUserSegments(c.ctx, &req.ToAdd, &req.ToRemove, req.Id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, segments)
}
