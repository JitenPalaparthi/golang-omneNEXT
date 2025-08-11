package main

import (
	"demo/handlers"
	"fmt"
	"net/http"
	"runtime"
)

func main() {

	//handleMap := make(map[string]func(http.ResponseWriter, *http.Request))

	println("Server started and running on port 8081")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello OmneNEXT")
	})

	http.HandleFunc("/ping", handlers.Ping)

	http.HandleFunc("/health", handlers.Health)

	userHandler := handlers.NewUserHandler("users.dat")
	http.HandleFunc("/users", userHandler.Create)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		println(err.Error())
		runtime.Goexit()
	}

	// http.ListenAndServeTLS()
	//
}
