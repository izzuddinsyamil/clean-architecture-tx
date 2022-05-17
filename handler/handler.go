package handler

import (
	"context"
	errCode "repo-pattern-w-trx-management/error"
	"repo-pattern-w-trx-management/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

type usecase interface {
	GetUserList(ctx context.Context) (u []model.User, err error)
	GetUserById(ctx context.Context, id int) (u model.User, err error)
	CreateUser(ctx context.Context, name string, balance int) (err error)
}

type handler struct {
	log *logrus.Logger
	uc  usecase
}

func NewHandler(l *logrus.Logger, uc usecase) *handler {
	return &handler{
		log: l,
		uc:  uc,
	}
}

func (h *handler) HandleGetUser(c echo.Context) error {
	user, err := h.uc.GetUserList(c.Request().Context())
	if err != nil {
		return sendInternalErrorResponse(c, nil, "server error")
	}

	return sendSuccessResponse(c, user)
}

func (h *handler) HandleGetUserById(c echo.Context) error {
	param := new(GetUserRequest)
	if err := c.Bind(param); err != nil {
		h.log.Error(err)
		return sendBadRequestResponse(c, nil, "invalid request param")
	}

	userId, err := strconv.Atoi(param.Id)
	if err != nil {
		h.log.Error(err)
		return sendBadRequestResponse(c, nil, err.Error())
	}

	user, err := h.uc.GetUserById(c.Request().Context(), userId)
	if err != nil {
		h.log.Error(err)
		if stacktrace.GetCode(err) == errCode.EcodeNotFound {
			return sendBadRequestResponse(c, nil, "user not found")
		}

		return sendInternalErrorResponse(c, nil, "server error")
	}

	return sendSuccessResponse(c, user)
}

func (h *handler) HandleCreateUser(c echo.Context) error {
	param := new(CreateUserRequest)
	if err := c.Bind(param); err != nil {
		h.log.Error(err)
		return sendBadRequestResponse(c, nil, "invalid request param")
	}

	err := h.uc.CreateUser(c.Request().Context(), param.Name, param.Balance)
	if err != nil {
		h.log.Error(err)
		return sendInternalErrorResponse(c, nil, "server error")
	}

	return sendCreatedResponse(c, nil)
}
