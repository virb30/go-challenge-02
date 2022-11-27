package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Message struct {
	ID  string
	Msg string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Invalid CEP")
	}
	cep := args[1]

	c1 := make(chan Message)
	c2 := make(chan Message)

	go requestApiCep(cep, c1)
	go requestViaCep(cep, c2)

	select {
	case msg := <-c1:
		fmt.Printf("API: %s\n%s", msg.ID, msg.Msg)
	case msg := <-c2:
		fmt.Printf("API: %s\n%s", msg.ID, msg.Msg)
	case <-time.After(time.Second):
		panic("Timeout")
	}
}

func requestViaCep(cep string, ch chan Message) {
	cep = formatCep(cep, false)
	url := "http://viacep.com.br/ws/" + cep + "/json/"
	response := makeRequest(url)
	msg := Message{"viacep", response}
	ch <- msg
}

func requestApiCep(cep string, ch chan Message) {
	cep = formatCep(cep, true)
	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	response := makeRequest(url)
	msg := Message{"apicep", response}
	ch <- msg
}

func formatCep(cep string, dot bool) string {
	r := regexp.MustCompile("[^0-9]")
	cep = r.ReplaceAllString(cep, "")
	if dot {
		cep = fmt.Sprintf("%s-%s", cep[0:5], cep[5:8])
	}
	return cep
}

func makeRequest(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
