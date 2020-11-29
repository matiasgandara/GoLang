package main

import (
	"Tp_Final_GO/internal/config"
	"Tp_Final_GO/internal/database"
	"Tp_Final_GO/internal/service/chat"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	/* if err := createSchema(db); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} */

	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTrnasport(service)

	r := gin.Default()
	httpService.Register(r)
	r.Run()
	/* for _, m := range service.FindAll() {
		fmt.Println(m)
	} */

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

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar);`
	// execute query on server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	//or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES(?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}
