package controller

import (
	"context"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/service/interfaces"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/request"
	"github.com/capkeik/backend-trainee-assignment-2023/internal/web/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"math"
	"net/http"
	"strconv"
)

type UserController struct {
	ctx      context.Context
	service  interfaces.UserService
	validate *validator.Validate
}

func NewUsers(ctx context.Context, service interfaces.UserService) *UserController {
	return &UserController{
		ctx:      ctx,
		service:  service,
		validate: validator.New(),
	}
}

func (c UserController) Get(ctx echo.Context) error {
	userID := ctx.Param("id")
	log.Println("UserController: ", "Getting usr by id ", userID)
	id64, err := strconv.ParseInt(userID, 10, 32)

	if id64 <= 0 || id64 >= math.MaxInt32 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
	}
	id := int32(id64)

	if err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid ID"})
	}

	slugs, err := c.service.GetUserSegments(id)
	if err != nil {
		log.Println("Error:" + err.Error())
		switch {
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Internal server error"})
		}
	}
	log.Println("UserController: ", "Responding user id ", userID)
	return ctx.JSON(http.StatusOK, response.Slugs{Slugs: *slugs, ID: id})
}

func (c UserController) Create(ctx echo.Context) error {
	var req request.UserReq
	log.Println("UserController: ", "Creating new user ", req.ID)
	if err := ctx.Bind(&req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	if err := c.validate.Struct(req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	createUser, err := c.service.CreateUser(c.ctx, req.ID)
	if err != nil {
		log.Println("Error:", "User already exists: ", req.ID)
		return ctx.JSON(http.StatusConflict, map[string]interface{}{"error": "User already exists"})
	}
	log.Println("UserController: ", "User Successfully created with id ", createUser.ID)

	return ctx.JSON(http.StatusOK, map[string]interface{}{"id": createUser.ID})
}
func (c UserController) UpdateSegments(ctx echo.Context) error {
	var req request.ChangeSegReq
	log.Println("UserController: ", "Updating user segments by id ", req.ID)
	if err := ctx.Bind(&req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	if err := c.validate.Struct(req); err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid JSON"})
	}

	changes, err := c.service.UpdateUserSegments(c.ctx, &req.ToAdd, &req.ToRemove, req.ID)
	if err != nil {
		log.Println("Error:" + err.Error())
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "Internal server error"})
	}
	log.Println("UserController: ", "Updated user segments by id ", req.ID)
	log.Println("UserController: ", "Removed slugs: ", changes.Removed)
	log.Println("UserController: ", "Added slugs: ", changes.Added)

	return ctx.JSON(http.StatusOK, changes)
}
