package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func main() {

	l := slog.New(slog.Default().Handler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		l.Info("insight root handler")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		l.Info("insight hello handler")
		body, err := io.ReadAll(r.Body)
		// err = errors.New("something went wrong")
		if err != nil {
			// w.Write([]byte(err.Error()))
			// w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// fmt.Println(string(body))
		// fmt.Println(r.UserAgent())
		res := fmt.Sprintf("Hello %s", string(body))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
		return
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		l.Info("insight goodbye handler")
	})

	if err := http.ListenAndServe(":9090", nil); err != nil {

		log.Fatalln(err)
	}
}
