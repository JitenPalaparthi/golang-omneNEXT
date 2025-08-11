package main

import (
	"demo/handlers"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var (
	PORT string
)

func main() {

	//handleMap := make(map[string]func(http.ResponseWriter, *http.Request))

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	println("Server started and running on port-->", PORT)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello OmneNEXT")
	})

	http.HandleFunc("/ping", handlers.Ping)

	http.HandleFunc("/health", handlers.Health)

	userHandler := handlers.NewUserHandler("users.dat")

	http.HandleFunc("/users", userHandler.Create)

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		println(err.Error())
		runtime.Goexit()
	}

	// http.ListenAndServeTLS()
	//

}
