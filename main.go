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
	http.HandleFunc("/ping", healthHandler)
	http.HandleFunc("/error", errorHandler)
	log.Printf("Server start up port: %s\n", port)
	log.Println("All API:\n /test\n /ping\n /error")

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

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[ACCESS LOG]: /ping")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[ACCESS LOG]: /error")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("ERROR"))
}
