package dependencyinjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, str string) {
	fmt.Fprintf(writer, "Hello, %s", str)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
