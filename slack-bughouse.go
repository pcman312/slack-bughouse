package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var t test_struct
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Test)
}

func randomizer() {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond())) // no shuffling without this line

	for i := len(players) - 1; i > 0; i-- {
		j := rand.Intn(i)
		players[i], players[j] = players[j], players[i]
	}
}

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
