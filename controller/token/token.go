package token

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func List (w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "token.list")
}

