package main

import (
	"github.com/hal-iosk/hal-cinema/router"
)

func main() {
	r := router.GetRouter()
	r.Run(":2000")
}
