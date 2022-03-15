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

	"github.com/friendsofgo/graphiql"
	"github.com/go-chi/chi/v5"
	"github.com/teal-finance/garcon"

	"github.com/teal-finance/rainbow/pkg/provider"
	"github.com/teal-finance/rainbow/pkg/rainbow"
	"github.com/teal-finance/rainbow/pkg/rainbow/api/apigraphql"
	"github.com/teal-finance/rainbow/pkg/rainbow/storage/dbram"
)

func main() {
	parseFlags()

	// Start the service in background
	service := rainbow.NewService(
		provider.AllProviders(),
		dbram.NewDB())
	go service.Run()

	g, err := garcon.New(
		garcon.WithURLs(*mainAddr),
		garcon.WithDocURL("/doc"),
		garcon.WithServerHeader("Rainbow-v0"),
		garcon.WithLimiter(*reqBurst, *reqPerMinute),
		garcon.WithProm(*expPort),
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

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	r.With(g.JWT.Set).Mount("/", WebHandler(g, *wwwDir))
	r.With(g.JWT.Chk).Mount("/graphql", apigraphql.Handler(s))
	r.With(g.JWT.Chk).Mount("/graphiql", graphiqlHandler)

	return r
}
