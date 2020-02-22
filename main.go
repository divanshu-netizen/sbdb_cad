package main

import "news/internal/services"

func main() {
	neo := services.NewNeoService()
	neos, err := neo.FindNeoBy(services.NeoQueryOptions{})
	if err != nil {
		println(err)
	}
	println("ooh")
	print(len(neos))

	for _, neo := range neos {
		println(neo.FullName)
	}
}
