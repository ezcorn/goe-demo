package action

import (
	"fmt"
	"net/http"
)

func HelloAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}
