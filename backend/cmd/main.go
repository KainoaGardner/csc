package main

import (
	"context"
	"fmt"
	"github.com/KainoaGardner/csc/internal/api"
	"github.com/KainoaGardner/csc/internal/config"
	"github.com/KainoaGardner/csc/internal/db"
	"log"
	"time"
)

func main() {
	config := config.LoadConfig()

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := db.ConnectToDB(context, config.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context); err != nil {
			log.Fatal(err)
		}
	}()

	server := api.NewAPIServer(fmt.Sprintf("%s:%s", config.PublicHost, config.Port))
	if err := server.Run(client, config); err != nil {
		log.Fatal(err)
	}
}
