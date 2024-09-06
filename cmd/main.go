package cmd

import (
	"log"

	"github.com/tejasvi541/GoColl/internal/server"
)

func main() {
	if err:=server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
}