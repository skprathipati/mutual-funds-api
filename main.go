package main

import "mutual-funds-api/routes"

func main() {
	r := routes.InitializeRoutes()

	r.Run(":8080")
}
