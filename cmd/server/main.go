package server

import (
	"log"
	"net/http"

	"github.com/tflorimo/habit-tracker/internal/config"
)

func main() {
	cfg := config.NewConfig()

	http.ListenAndServe(":"+cfg.HttpPort, nil)
	log.Printf("Server started on port %s", cfg.HttpPort)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		log.Println("Request received")
	})
}
