package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GabrielL915/Api-Rest-Go/db"
	"github.com/GabrielL915/Api-Rest-Go/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var itemId = "id"

func products(router chi.Router) {
	router.Get("/", getAllProducts)
	router.Post("/", createProduct)
	router.Route(fmt.Sprintf("/{%s}", itemId), func(router chi.Router) {
		router.Use(productContext)
		router.Get("/", getProduct)
		router.Put("/", updateProduct)
		router.Delete("/", deleteProduct)
	})
}

func productContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemId := chi.URLParam(r, itemId)
		if itemId == "" {
			render.Render(w, r, ErrRender(fmt.Errorf("products ID is required")))
			return
		}
		id, err := strconv.Atoi(itemId)
		if err != nil {
			render.Render(w, r, ErrRender(fmt.Errorf("invalid product ID")))
		}
		ctx := context.WithValue(r.Context(), itemId, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	data := &models.Product{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.CreateProduct(data); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
	if err := render.Render(w, r, data); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := dbInstance.GetAllProducts()
	if err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
	if err := render.Render(w, r, products); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemId).(string)
	product, err := dbInstance.GetProductById(itemId)

	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrRender(err))
		}
		return
	}
	if err := render.Render(w, r, product); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemId).(string)
	data := &models.Product{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	product, err := dbInstance.UpdateProduct(itemId, *data)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}
	if err := render.Render(w, r, &product); err != nil {
		render.Render(w, r, ServerErrorRender(err))
		return
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	itemId := r.Context().Value(itemId).(string)
	if err := dbInstance.DeleteProduct(itemId); err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRender(err))
		}
		return
	}
}
