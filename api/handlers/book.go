package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/husanmusa/pro-book-service/api/http"
	"github.com/husanmusa/pro-book-service/genproto/book_service"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body book_service.CreateUserRequest true "CreateUserRequestBody"
// @Success 201 {object} book_service.User "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user book_service.CreateUserRequest

	err := c.BodyParser(&user)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookService().CreateUser(
		c.Context(),
		&user,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetUser godoc
// @ID get_user
// @Router /user/{id} [GET]
// @Summary Get User
// @Description Get User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} book_service.User "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().GetUser(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// ListUser	godoc
// @ID list_user
// @Router /user [GET]
// @Summary List User
// @Description List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} book_service.ListUsersResponse "ListUserResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ListUser(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	resp, err := h.services.BookService().ListUsers(
		c.Context(),
		&book_service.ListUsersRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
		},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// FindBookFromUsers godoc
// @ID find_book
// @Router /user/find_book/{id} [GET]
// @Summary List User has a book
// @Description List User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} book_service.ListUsersResponse "ListUserResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) FindBookFromUsers(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id, "dddddddd")
	resp, err := h.services.BookService().FindBookFromUsers(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateUser godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body book_service.UpdateUserRequest true "UpdateUserRequestBody"
// @Success 200 {object} book_service.User "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var user book_service.UpdateUserRequest

	err := c.BodyParser(&user)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookService().UpdateUser(
		c.Context(),
		&user,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} http.Response{data=string} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().DeleteUser(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// CreateBook godoc
// @ID create_book
// @Router /book [POST]
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body book_service.CreateBookRequest true "CreateBookRequestBody"
// @Success 201 {object} book_service.Book "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateBook(c *fiber.Ctx) error {
	var book book_service.CreateBookRequest

	err := c.BodyParser(&book)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookService().CreateBook(
		c.Context(),
		&book,
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetBook godoc
// @ID get_book
// @Router /book/{id} [GET]
// @Summary Get Book
// @Description Get Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} book_service.Book "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().GetBook(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// ListBook godoc
// @ID list_book
// @Router /book [GET]
// @Summary List Book
// @Description List Book
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} book_service.ListBooksResponse "ListBookResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ListBook(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	resp, err := h.services.BookService().ListBooks(
		c.Context(),
		&book_service.ListBooksRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
		},
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// ListBooksByUserId godoc
// @ID list_book_by_user_id
// @Router /book/user/{id} [GET]
// @Summary List Book User
// @Description List Book User
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string false "id"
// @Success 200 {object} book_service.ListBooksResponse "ListBookResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ListBooksByUserId(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().ListBooksByUserId(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateBook godoc
// @ID update_book
// @Router /book/{id} [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body book_service.UpdateBookRequest true "UpdateBookRequestBody"
// @Success 200 {object} book_service.Book "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	var book book_service.UpdateBookRequest

	err := c.BodyParser(&book)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	id := c.Params("id")
	book.Id = id
	resp, err := h.services.BookService().UpdateBook(
		c.Context(),
		&book,
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// DeleteBook godoc
// @ID delete_book
// @Router /book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} http.Response{data=string} "Book data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().DeleteBook(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// CreateSaleBook godoc
// @ID create_sale_book
// @Router /sale [POST]
// @Summary Create SaleBook
// @Description Create SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param sale body book_service.CreateSaleBookRequest true "CreateSaleBookRequestBody"
// @Success 201 {object} book_service.SaleBook "SaleBook data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSaleBook(c *fiber.Ctx) error {
	var sale book_service.CreateSaleBookRequest

	err := c.BodyParser(&sale)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	resp, err := h.services.BookService().CreateSaleBook(
		c.Context(),
		&sale,
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.Created, resp)
}

// GetSaleBook godoc
// @ID get_sale_book
// @Router /sale/{id} [GET]
// @Summary Get SaleBook
// @Description Get SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Success 200 {object} book_service.SaleBook "SaleBook data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSaleBook(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().GetSaleBook(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// ListSaleBook godoc
// @ID list_sale_book
// @Router /sale [GET]
// @Summary List SaleBook
// @Description List SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Success 200 {object} book_service.ListSaleBooksResponse "ListSaleBookResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ListSaleBook(c *fiber.Ctx) error {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		return h.handleResponse(c, http.InvalidArgument, err.Error())
	}

	resp, err := h.services.BookService().ListSaleBooks(
		c.Context(),
		&book_service.ListSaleBooksRequest{
			Offset: int64(offset),
			Limit:  int64(limit),
		},
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// UpdateSaleBook godoc
// @ID update_sale_book
// @Router /sale/{id} [PUT]
// @Summary Update SaleBook
// @Description Update SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Param sale body book_service.UpdateSaleBookRequest true "UpdateSaleBookRequestBody"
// @Success 200 {object} book_service.SaleBook "SaleBook data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSaleBook(c *fiber.Ctx) error {
	var sale book_service.UpdateSaleBookRequest

	err := c.BodyParser(&sale)
	if err != nil {
		return h.handleResponse(c, http.BadRequest, err.Error())
	}

	id := c.Params("id")
	sale.Id = id

	resp, err := h.services.BookService().UpdateSaleBook(
		c.Context(),
		&sale,
	)
	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}

// DeleteSaleBook godoc
// @ID delete_sale_book
// @Router /sale/{id} [DELETE]
// @Summary Delete SaleBook
// @Description Delete SaleBook
// @Tags SaleBook
// @Accept json
// @Produce json
// @Param id path string true "SaleBook ID"
// @Success 200 {object} http.Response{data=string} "SaleBook data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSaleBook(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.services.BookService().DeleteSaleBook(
		c.Context(),
		&book_service.ByIdReq{Id: id},
	)

	if err != nil {
		return h.handleResponse(c, http.GRPCError, err.Error())
	}

	return h.handleResponse(c, http.OK, resp)
}
