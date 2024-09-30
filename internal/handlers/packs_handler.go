package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pr1sonmike/golang-packing-challenge/internal/service"
)

type PacksHandler struct {
	Service service.PacksService
}

func NewPacksHandler(service service.PacksService) *PacksHandler {
	return &PacksHandler{Service: service}
}

type CalculatePacksRequest struct {
	Number int   `json:"number" validate:"required,gt=0"`
	Packs  []int `json:"packs" validate:"required,dive,gte=0"`
}

type CalculatePacksResponse struct {
	Packs map[int]int `json:"packs"`
}

func (h *PacksHandler) CalculatePacks(c echo.Context) error {
	var req CalculatePacksRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request format"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	packs, err := h.Service.FulfillItems(c.Request().Context(), req.Number, req.Packs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to calculate packs"})
	}

	return c.JSON(http.StatusOK, CalculatePacksResponse{Packs: packs})
}
