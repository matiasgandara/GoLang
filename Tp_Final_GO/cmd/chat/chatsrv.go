package main

import (
	"fmt"
	"os"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.DB.Version)
}
