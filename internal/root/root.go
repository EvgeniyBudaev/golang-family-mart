package root

import (
	"io"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is my website!\n")
}
