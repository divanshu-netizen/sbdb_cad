package main

import (
	"news/internal/services"
)

func main() {
	neo := services.NewNeoService()
	_, err := neo.FindNeoBy(services.NeoQueryOptions{})
	if err != nil {
		println(err)
	}
}
