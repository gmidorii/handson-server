package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type testResponse struct {
	Test string    `json:"test"`
	Time time.Time `json:"time"`
}

func main() {
	port := flag.String("p", "8080", "port")
	flag.Parse()

	handler(*port)
}

func handler(port string) error {
	http.HandleFunc("/test", testHandler)
	log.Printf("Server runnint port: %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}
	return nil
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	res := testResponse{
		Test: "test",
		Time: time.Now(),
	}

	str, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}
