package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/config"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/database"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/service/motos"
)

func main() {
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	/* 	defer db.Close()
	 */
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Intanciacion de servicio en base de datos
	service, _ := motos.New(db, cfg)
	httpService := motos.NewHTTPTransport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "este es el servicio de configuracion")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}
