package main

import "net/http"

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("index"))
	})

	http.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("users"))
	})

	http.ListenAndServe(":8080", nil)
}
