package main

import (
	"fmt"
	auth "todo/Auth"
	"todo/pck/dbconnection"
	"todo/pck/handlers"
	"todo/pck/middleware"
	"todo/pck/repository"
	"todo/pck/server"
	"todo/pck/service"
	"todo/redis"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("connecting to database")
	db, err := dbconnection.Dbconnect()
	log.Info().Msg("CONNECTED TO DATABASE")

	if err != nil {
		return
	}
	repo, err := repository.Newrepository(db)
	if err != nil {
		return
	}
	// redis connection
	rdb, err := dbconnection.ConnectRedis()
	if err != nil {
		fmt.Println("error when connecting the redis")
	}
	rdb1, err := redis.NewRediers(rdb)
	if err != nil {
		return
	}

	auth, err := auth.NewAuth(auth.SingNature)
	if err != nil {
		return
	}

	mid, err := middleware.NewMiddleware(auth)
	if err != nil {
		return
	}
	service, err := service.NewService(repo, auth, rdb1)
	if err != nil {
		return
	}
	handlers, err := handlers.NewHandler(service)
	if err != nil {
		return
	}

	server.StartingServer(handlers, mid)
}
