package main

import (
	"kf-golang-api/routes"
)

func main() {
	r := routes.SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
