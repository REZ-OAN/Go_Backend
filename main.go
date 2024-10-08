package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/REZ-OAN/simplebank/utils"

	db "github.com/REZ-OAN/simplebank/database/sqlc"

	"github.com/REZ-OAN/simplebank/api"

	_ "github.com/lib/pq"
)

func main() {
	path := os.Getenv("path")
	name := os.Getenv("name")
	ext := os.Getenv("ext")

	if path == "" || name == "" || ext == "" {
		log.Fatalln("All arguments should be given!!!")
		return
	}
	config, err := utils.LoadConfig(path, name, ext)

	if err != nil {
		log.Fatalf("cannot get configs: %v", err)
	}

	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server ", err)
		return
	}
	SERVER_URL := fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT)
	err = server.Start(SERVER_URL)

	if err != nil {
		log.Fatalf("Cannont Connect to the server : %v", err)
	}
}
