package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/pro-book-service/api/http"
	"github.com/husanmusa/pro-book-service/config"
	"github.com/husanmusa/pro-book-service/grpc/client"
	"github.com/saidamir98/udevs_pkg/logger"
	"strconv"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *fiber.Ctx, status http.Status, data interface{}) error {
	return c.Status(status.Code).JSON(http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *fiber.Ctx) (int, error) {
	offset := c.Query("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offset)
}

func (h *Handler) getLimitParam(c *fiber.Ctx) (int, error) {
	limit := c.Query("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(limit)
}
