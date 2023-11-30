package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var Logger = zerolog.New(os.Stdout).With().Logger()

type reply struct {
	success string
	err     error
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan reply, 100)
	// set timeout accordingly
	client := http.Client{Timeout: 60 * time.Second}

	for id := 0; id < 100; id++ {
		wg.Add(1)
		go send(id, &wg, ch, client)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result.success)
	}
}

func send(clientId int, wg *sync.WaitGroup, ch chan<- reply, client http.Client) {
	defer wg.Done()

	Logger.Info().Msgf("Sending clientId %n", clientId)

	request, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://localhost:8080/?clientId="+strconv.Itoa(clientId), nil)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		Logger.Error().Msg("error client.Do()")
		return
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		Logger.Error().Msg("error reading body")
	}

	r := reply{
		success: string(body),
		err:     err,
	}
	ch <- r
}
