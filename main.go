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

	if err := handler(*port); err != nil {
		log.Println("Failed server start up")
		log.Fatalln(err)
	}
}

func handler(port string) error {
	http.HandleFunc("/test", testHandler)
	log.Printf("Server start up port: %s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		return err
	}
	return nil
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	log.Println("[ACCESS LOG]: /test")
	res := testResponse{
		Test: "test",
		Time: now,
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
