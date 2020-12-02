package main

import(
	"github.com/gin-gonic/gin"
)

func main {
	cfg := readConfig()
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