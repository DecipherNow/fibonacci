package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("starting the server")

	client := &http.Client{}

	rand.Seed(time.Now().UTC().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if port == "" {
		port = "localhost"
	}

	allRoutes := []string{
		"/fibonacci/0",
		"/fibonacci/1",
		"/fibonacci/2",
		"/fibonacci/3",
		"/fibonacci/4",
		"/fibonacci/5",
		"/fibonacci/6",
		"/fibonacci/",
		"/fibonacci/1",
		"/fibonacci/1",
	}

	route := allRoutes[0]

	for {
		randRoute := rand.Intn(45)
		route = fmt.Sprintf("/fibonacci/%d", randRoute)
		fmt.Printf(fmt.Sprintf("http://%s:%s%s : ", host, port, route))
		resp, _ := client.Get(fmt.Sprintf("http://%s:%s%s", host, port, route))
		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		time.Sleep(time.Second / 2)
	}
}
