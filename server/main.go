package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var Logger = zerolog.New(os.Stdout).With().Logger()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(10) + 1
		duration := time.Duration(random) * time.Second
		if duration == 5 {
			Logger.Error().Msg("BOOM!")
			panic("BOOM!")
		}
		time.Sleep(duration)
		queryParams := r.URL.Query()
		clientId := queryParams.Get("clientId")
		fmt.Fprintf(w, "SERVER GOT clientId %v - DONE PROCESSING AFTER %v \n", clientId, duration)
	})

	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		Logger.Error().Msg("Can't start server")
	}
}
