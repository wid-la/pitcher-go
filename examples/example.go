package main

import "github.com/wid-la/pitcher-go"

func main() {
	vtURL := "https://api.versatile.la"
	c := pitcher.NewClient(vtURL)

	c.Signin("", "")
}
