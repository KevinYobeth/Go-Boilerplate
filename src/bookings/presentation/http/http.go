package http

import (
	"context"
	"go-boilerplate/shared/database"
	respond "go-boilerplate/shared/response"
	"go-boilerplate/shared/utils"
	"go-boilerplate/src/bookings/domain/bookings"
	"go-boilerplate/src/bookings/services"
	"go-boilerplate/src/bookings/services/command"
	"go-boilerplate/src/bookings/services/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTPTransport struct {
	app     *services.Application
	manager database.TransactionManager
}

func NewBookingsHTTPServer(app *services.Application) HTTPTransport {
	return HTTPTransport{app: app}
}

func (h HTTPTransport) TransactionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return h.manager.RunInTransaction(echoCtx.Request().Context(), func(c context.Context) error {
			ctx := echoCtx.Echo().NewContext(echoCtx.Request().WithContext(c), echoCtx.Response())
			ctx.SetParamNames(echoCtx.ParamNames()...)
			ctx.SetParamValues(echoCtx.ParamValues()...)

			return next(ctx)
		})
	}
}

func (h HTTPTransport) RegisterHTTPRoutes(r *echo.Group) {
	api := r.Group("/v1")
	RegisterHandlers(api, h)
}

func (h HTTPTransport) GetBookings(c echo.Context) error {
	bookingsObj, err := h.app.Queries.GetBookings.Execute(c.Request().Context())
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetBookingsResponse{
		Data:    TransformToHTTPBookings(bookingsObj),
		Message: "success get bookings",
	})
	return nil
}

func (h HTTPTransport) CreateBooking(c echo.Context) error {
	var request bookings.CreateBookingDto
	if err := c.Bind(&request); err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	err := h.app.Commands.CreateBooking.Execute(c.Request().Context(), command.CreateBookingParams{
		BookID:   request.BookID,
		DateFrom: request.DateFrom,
		DateTo:   request.DateTo,
	})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusCreated, MessageResponse{
		Message: "success create booking",
	})
	return nil
}

func (h HTTPTransport) GetBooking(c echo.Context, id string) error {
	parsedUUID, err := utils.ParseUUID(id)
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	booking, err := h.app.Queries.GetBooking.Execute(c.Request().Context(), query.GetBookingParams{ID: parsedUUID})
	if err != nil {
		respond.SendHTTP(c, err)
		return err
	}

	c.JSON(http.StatusOK, GetBookingResponse{
		Data:    TransformToHTTPBooking(booking),
		Message: "success get booking",
	})
	return nil
}
