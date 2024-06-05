package main

import (
	"Ntt_DATA/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
