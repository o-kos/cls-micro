package main

import (
	"fmt"
	"log"

	"github.com/o-kos/cls-micro/pkg"
)

func main() {
	cfg, err := config.New("cls")
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println(cfg)
}
