package main

import (
	"Tp_Final_GO/internal/config"
	"Tp_Final_GO/internal/database"
	"Tp_Final_GO/internal/service/chat"
	"flag"
	"fmt"
	"os"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	service, _ := chat.New(db, cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}

}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "este es el servicio de config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}
