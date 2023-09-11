package main

import (
	dependencyinjection "dependency_injection"
	"log"
	"net/http"
	"os"
)

func main() {
	dependencyinjection.Greet(os.Stdout, "main")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencyinjection.MyGreetHandler)))
}
