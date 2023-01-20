package server

import (
	"fmt"
	"net/http"
)

type Mux struct {
	Mux *http.ServeMux
}

func NewMux() *Mux {
	fmt.Printf("Starting multiplexer...\n")

	mux := http.NewServeMux()

	return &Mux{
		Mux: mux,
	}
}
