package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/config"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/database"
	"github.com/matiasgandara/GoLang/tree/main/Entregable_Go_Motos/internal/service/motos"
)

func main() {
	cfg := readConfig()

	// Instanciacion de db
	db, err := database.NewDatabase(cfg)
	defer db.Close()

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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS  motos (
    id integer NOT NULL CONSTRAINT motos_pk PRIMARY KEY AUTOINCREMENT,
	patente varchar(6) NOT NULL,
    modelo integer NOT NULL,
	nombre varchar(30) NOT NULL,
    cilindrada integer NOT NULL,
	color varchar(30) NOT NULL,
    );`
	// execute query on server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertMoto := `INSERT INTO motos (patente, modelo, nombre, cilindrada, color) VALUES(?,?,?,?,?)`
	s := fmt.Sprintf("ABC123", 2015, "Tornado", 250, "azul")
	db.MustExec(insertMoto, s)
	return nil
}
