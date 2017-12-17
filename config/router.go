package config

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	classDeliver "github.com/goweb4/class/delivery"
	classRepo "github.com/goweb4/class/repository/psql"
	classUsecase "github.com/goweb4/class/usecase"
	productDeliver "github.com/goweb4/product/delivery/http"
	productRepo "github.com/goweb4/product/repository/psql"
	productCase "github.com/goweb4/product/usecase"
	"github.com/goweb4/services/cache"
	userDeliver "github.com/goweb4/user/delivery/http"
	userRepo "github.com/goweb4/user/repository/psql"
	userCase "github.com/goweb4/user/usecase"
)

// Router func
func Router(db *sql.DB, c cache.Cache) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	addUserRoutes(r, db, c)
	addProductRoutes(r, db, c)
	addClassRoutes(r, db, c)
	return r
}

func addUserRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := userRepo.NewUserRepository(db)
	uc := userCase.NewUserUsecase(repo)
	userDeliver.NewUserController(r, uc, c)
}

func addProductRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := productRepo.NewProductRepository(db)
	uc := productCase.NewProductUsecase(repo)
	productDeliver.NewProductController(r, uc, c)
}

func addClassRoutes(r *chi.Mux, db *sql.DB, c cache.Cache) {
	repo := classRepo.NewClassRepository(db)
	uc := classUsecase.NewClassUsecase(repo)
	classDeliver.NewClassController(r, uc, c)
}
