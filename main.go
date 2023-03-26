package main

import "golang_microservices/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
