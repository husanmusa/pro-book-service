package api

import (
	"encoding/json"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/husanmusa/pro-book-service/api/docs"
	"github.com/husanmusa/pro-book-service/api/handlers"
	"github.com/husanmusa/pro-book-service/config"
)

// SetupRouter godoc
// @description This is api gateway
// @termsOfService https://t.me/husanmusa
func SetupRouter(h handlers.Handler, cfg config.Config) *fiber.App {
	r := fiber.New(fiber.Config{JSONEncoder: json.Marshal, BodyLimit: 100 * 1024 * 1024})
	r.Use(logger.New(), cors.New())
	r.Use(MaxAllowed(5000))

	r.Get("/swagger/*", swagger.HandlerDefault)

	r.Get("/ping", h.Ping)
	r.Get("/config", h.GetConfig)

	// CLIENT SERVICE
	r.Post("/user", h.CreateUser)
	r.Get("/user/:id", h.GetUser)
	r.Get("/user", h.ListUser)
	r.Get("/user/find_book/:id", h.FindBookFromUsers)
	r.Put("/user/:id", h.UpdateUser)
	r.Delete("/user/:id", h.DeleteUser)

	r.Post("/book", h.CreateBook)
	r.Get("/book/:id", h.GetBook)
	r.Get("/book", h.ListBook)
	r.Get("/book/user/:id", h.ListBooksByUserId)
	r.Put("/book/:id", h.UpdateBook)
	r.Delete("/book/:id", h.DeleteBook)

	r.Post("/sale", h.CreateSaleBook)
	r.Get("/sale/:id", h.GetSaleBook)
	r.Get("/sale", h.ListSaleBook)
	r.Put("/sale/:id", h.UpdateSaleBook)
	r.Delete("/sale/:id", h.DeleteSaleBook)

	return r
}

func MaxAllowed(n int) fiber.Handler {
	var countReq int64
	sem := make(chan struct{}, n)
	acquire := func() {
		sem <- struct{}{}
		countReq++
		fmt.Println("countRequest: ", countReq)
	}

	release := func() {
		select {
		case <-sem:
		default:
		}
		countReq--
		fmt.Println("countResp: ", countReq)
	}

	return func(c *fiber.Ctx) error {
		acquire()       // before request
		defer release() // after request

		return c.Next()
	}
}
