package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type UsdbrlOut struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var cotacao UsdbrlOut
	json.Unmarshal(body, &cotacao)

	f, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	_, err = f.Write([]byte("Dólar:" + cotacao.Bid + "\n"))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	f.Close()
}
