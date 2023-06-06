package main

import "github.com/dwdarm/go-url-shortener/cmd"

func main() {
	r := cmd.Init()

	r.Run()
}
