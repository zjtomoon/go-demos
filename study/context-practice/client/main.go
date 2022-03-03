package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", respBytes)
}
