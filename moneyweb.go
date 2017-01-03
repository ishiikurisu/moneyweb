package main

import "github.com/ishiikurisu/moneyweb/controller"

func main() {
    server := controller.CreateServer()
    server.Serve()
}
