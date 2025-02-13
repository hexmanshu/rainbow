// Copyright (c) 2021 Teal.Finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/garcon"
	"github.com/teal-finance/garcon/webserver"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	parseFlags()

	var providers []rainbow.Provider
	switch len(*alert) {
	case 0:
		providers = provider.AllProviders()
	default:
		providers = provider.AllProvidersWithAlert(
			provider.NewOracle(*alert),
		)
	}

	// Start the service in background
	service := rainbow.NewService(
		providers,
		dbram.NewDB())
	go service.Run()

	g, err := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDocURL("/doc"),
		garcon.WithServerHeader("Rainbow-v0"),
		garcon.WithLimiter(*reqBurst, *reqPerMinute),
		garcon.WithProm(*expPort, *mainAddr),
		garcon.WithDev(*dev))
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:              listenAddr,
		Handler:           g.Middlewares.Then(handler(&service, g)),
		ReadTimeout:       time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Minute, // Garcon.Limiter postpones response, attacker should wait long time.
		IdleTimeout:       time.Second,
		ConnState:         g.ConnState,
		ErrorLog:          log.Default(),
	}

	log.Print("Server listening on http://localhost", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func handler(s *rainbow.Service, g *garcon.Garcon) http.Handler {
	r := chi.NewRouter()

	// Static website: set the cookie only when visiting index.html
	web := webserver.WebServer{Dir: *wwwDir, ResErr: g.ResErr}
	r.With(g.JWT.Set).NotFound(web.ServeFile("index.html", "text/html; charset=utf-8")) // catch index.html and other Vue sub-folders
	r.Get("/favicon.ico", web.ServeFile("favicon.ico", "image/x-icon"))
	r.Get("/favicon.png", web.ServeFile("favicon.png", "image/png"))
	r.Get("/preview.jpg", web.ServeFile("preview.jpg", "image/jpeg"))
	r.With(g.JWT.Chk).Get("/js/*", web.ServeDir("text/javascript; charset=utf-8"))
	r.With(g.JWT.Chk).Get("/assets/*", web.ServeAssets())

	r.Route("/v0", func(r chi.Router) {
		h := api.Handler{Service: s}

		// HTTP API
		r.With(g.JWT.Vet).Route("/options", func(r chi.Router) {
			r.HandleFunc("/", h.Options)
			r.HandleFunc("/{asset}", h.Options)
			r.HandleFunc("/{asset}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}", h.Options)
			r.HandleFunc("/{asset}/{expiry}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}/", h.Options)
			r.HandleFunc("/{asset}/{expiry}/{provider}/{format}", h.Options)
		})

		r.With(g.JWT.Chk).Get("/bff/cp", h.CallPut)

		// GraphQL API (and interactive API in developer mode)
		r.With(g.JWT.Chk).Mount("/graphql", h.GraphQLHandler())
		if *dev {
			r.Mount("/graphiql", api.InteractiveGQLHandler("/v0/graphql"))
		}
	})

	return r
}
