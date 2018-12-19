//go:generate swagger generate spec

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Recursive fib function
func fib_recurse(val int) (int, error) {
	if val < 0 {
		err := fmt.Errorf("cannot calculate sequence on negative value: %d", val)
		return 0, err
	}

	if val == 0 {
		err := fmt.Errorf("%d is not an acceptable value", val)
		return 0, err
	}

	if val == 1 || val == 2 {
		return 1, nil
	}

	if val == 3 {
		return 2, nil
	}

	val1, _ := fib_recurse(val - 1)
	val2, _ := fib_recurse(val - 2)
	return val1 + val2, nil
}

func main() {
	fmt.Println("starting the server")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Alive")
	})

	rand.Seed(time.Now().UTC().UnixNano())

	// swagger:route GET /ping ping
	//
	// Pings the service to check if it's alive.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Give me some math!"))
	})

	http.HandleFunc("/fibonacci/", func(w http.ResponseWriter, r *http.Request) {
		valStr := strings.Split(r.URL.Path, "/")[2]
		val, err := strconv.Atoi(valStr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Can't parse the number from that path!"))
		}

		ret, _ := fib_recurse(val)
		fmt.Println("Fibonacci recursively for %d", val)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("%d\n", ret)))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
