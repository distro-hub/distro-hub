package main

import (
    "net/http"
    "distro-hub/views"
    "distro-hub/domain"
    "context"
)


func main() {
    mux := http.NewServeMux()

    distros := domain.Distros{
        {
            ID: 1,
            Name: "Ubuntu",
        },
        {
            ID: 2,
            Name: "Debian",
        },
        {
            ID: 3,
            Name: "Arch",
        },
    }

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "currentUser", "Radoje")
        views.HomePage("Distro Hub", distros).Render(ctx, w)
    })


    s := &http.Server {
        Addr: ":6969",
        Handler: mux,
    }

    s.ListenAndServe()
}
