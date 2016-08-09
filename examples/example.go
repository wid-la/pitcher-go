package main

import "github.com/wid-la/pitcher-go"

func main() {
	vtURL := "https://api.test.org"
	c := pitcher.NewClient(vtURL)
	c.Signin("name@test.org", "password")

	session := c.CurrentSession()
	println("Token: ", session.Token)
}
