package main

import (
	hanlders "distro-hub/handlers"
	"distro-hub/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("POST /distro", func(w http.ResponseWriter, r *http.Request) {
	// 	id = id + 1
	// 	newDistro := domain.Distro{
	// 		ID:   id,
	// 		Name: r.FormValue("distro_name"),
	// 	}
	// 	fmt.Println(id)
	// 	distros = append(distros, newDistro)
	// 	views.DistroList(distros).Render(r.Context(), w)
	// })

	h := hanlders.NewHandler()
	// auth
	mux.Handle("/auth/", h.AuthMux())
	mux.HandleFunc("/", h.HomePage)

	// distro stuff
	mux.Handle("/distro/", h.DistroMux())

	s := &http.Server{
		Addr:    ":6969",
		Handler: middleware.Auth(mux),
	}

	s.ListenAndServe()
}
