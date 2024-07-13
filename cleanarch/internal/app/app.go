package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"lessons/cleanarch/internal/adapter/sqllite"
	"lessons/cleanarch/internal/controller"
	"lessons/cleanarch/internal/repository/inmemory"
	"lessons/cleanarch/internal/repository/sqlite"
	service "lessons/cleanarch/internal/service/userservice"
)

type Config struct {
	InMemory bool
	DBPath   string
}

type App struct {
	router *mux.Router
	c      *controller.UserController
}

func New(config Config) *App {
	var svc *service.UserService

	c := controller.NewUserController(svc)
	if config.InMemory {
		// Инициализация репозитория в памяти
		repo := inmemory.NewInMemoryUserRepository()
		svc = service.NewUserService(repo)
	} else {
		// Инициализация репозитория SQLite
		dbPath := config.DBPath
		if dbPath == "" {
			dbPath = "orders.db"
		}
		adapter, err := sqllite.NewSQLLiteAdapter(dbPath)
		if err != nil {
			log.Fatalf("Failed to connect to SQLite: %v", err)
		}
		repo := sqlite.NewSQLiteUserRepository(adapter)
		svc = service.NewUserService(repo)
	}

	app := &App{
		router: mux.NewRouter(),
		c:      c,
	}

	app.initializeRoutes()
	return app
}

func (a *App) initializeRoutes() {
	a.router.HandleFunc("/users/{id}/change-password", a.c.ChangePassword).Methods("POST")
}

func (a *App) Run(addr string) {
	log.Printf("Server started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, a.router))
}
