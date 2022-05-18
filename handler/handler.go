package handler

import (
	"context"
	"repo-pattern-w-trx-management/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type usecase interface {
	GetUserList(ctx context.Context) (u []model.User, err error)
	CreateUser(ctx context.Context, name string, balance int) (err error)
	Transact(ctx context.Context, senderId, receiverId, amount int) (err error)
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

func (h *handler) HandleTransact(c echo.Context) error {
	p := new(TransactRequest)
	if err := c.Bind(p); err != nil {
		h.log.Error(err)
		return sendBadRequestResponse(c, nil, "invalid request param")
	}

	err := h.uc.Transact(c.Request().Context(), p.SenderId, p.ReceiverId, p.Amount)
	if err != nil {
		h.log.Error(err)
		return sendInternalErrorResponse(c, nil, "server error")
	}

	return sendSuccessResponse(c, p)
}
