package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"product-service/internal/domain/book"
	"product-service/internal/service/library"
	"product-service/pkg/server/response"
	"product-service/pkg/store"
)

type BookHandler struct {
	libraryService *library.Service
}

func NewBookHandler(s *library.Service) *BookHandler {
	return &BookHandler{libraryService: s}
}

func (h *BookHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", h.list)
	r.Post("/", h.add)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.get)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})

	return r
}

// List of books from the database
//
//	@Summary	List of books from the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Success	200		{array}		response.Object
//	@Failure	500		{object}	response.Object
//	@Router		/books 	[get]
func (h *BookHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.libraryService.ListBooks(r.Context())
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Add a new book to the database
//
//	@Summary	Add a new book to the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		request	body		book.Request	true	"body param"
//	@Success	200		{object}	response.Object
//	@Failure	400		{object}	response.Object
//	@Failure	500		{object}	response.Object
//	@Router		/books [post]
func (h *BookHandler) add(w http.ResponseWriter, r *http.Request) {
	req := book.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	res, err := h.libraryService.AddBook(r.Context(), req)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Read the book from the database
//
//	@Summary	Read the book from the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"path param"
//	@Success	200	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/books/{id} [get]
func (h *BookHandler) get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.libraryService.GetBook(r.Context(), id)
	if err != nil && err != store.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == store.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}

	response.OK(w, r, res)
}

// Update the book in the database
//
//	@Summary	Update the book in the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		id		path	int				true	"path param"
//	@Param		request	body	book.Request	true	"body param"
//	@Success	200
//	@Failure	400	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/books/{id} [put]
func (h *BookHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	req := book.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	err := h.libraryService.UpdateBook(r.Context(), id, req)
	if err != nil && err != store.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == store.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}

// Delete the book from the database
//
//	@Summary	Delete the book from the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		id	path	int	true	"path param"
//	@Success	200
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/books/{id} [delete]
func (h *BookHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.libraryService.DeleteBook(r.Context(), id)
	if err != nil && err != store.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == store.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}
}

// List of authors from the database
//
//	@Summary	List of authors from the database
//	@Tags		books
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"path param"
//	@Success	200	{array}		response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/books/{id}/authors [get]
func (h *BookHandler) listAuthors(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	res, err := h.libraryService.ListBookAuthors(r.Context(), id)
	if err != nil && err != store.ErrorNotFound {
		response.InternalServerError(w, r, err)
		return
	}

	if err == store.ErrorNotFound {
		response.NotFound(w, r, err)
		return
	}

	response.OK(w, r, res)
}
