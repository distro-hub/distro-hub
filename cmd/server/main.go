package main

import (
	"distro-hub/internal/infrastructure/db/psql"
	"distro-hub/internal/interface/web/handler"
	"distro-hub/internal/services"
	"log"
	"net/http"
)

func main() {
	db, err := psql.NewDB("172.17.0.2", "5432", "radoje", "12345", "distro_hub")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	distroRepo := psql.NewDistroRepository(db)
	userRepo := psql.NewUserRepository(db)

	// Initialize services
	distroService := services.NewDistroService(distroRepo)
	userService := services.NewUserService(userRepo)

	webHandler := handler.NewWebHandler(distroService, userService)
	mux := setupRoutes(webHandler)

	log.Fatal(http.ListenAndServe(":6969", mux))
}

func setupRoutes(webHandler *handler.WebHandler) http.Handler {
	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/", webHandler.Home)
	mux.HandleFunc("/auth", webHandler.Auth)
	mux.HandleFunc("/auth/login-form", webHandler.LoginForm)
	mux.HandleFunc("/auth/register-form", webHandler.RegisterForm)
	// mux.HandleFunc("/auth/login", webHandler.Login)
	// mux.HandleFunc("/auth/register", webHandler.Register)

	// Serve static files
	// fs := http.FileServer(http.Dir("./web/static"))
	// mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
