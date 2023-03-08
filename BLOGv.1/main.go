package main

import (
	"GOproject/IM/config"
	"GOproject/IM/router"
)

func main() {
	config.Init()
	r := router.NewRoute()

	r.Run(config.HttpPort)
}
