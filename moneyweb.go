package main

import "github.com/ishiikurisu/logeyweb/controller"

func main() {
    server := controller.CreateServer()
    server.Serve()
}
