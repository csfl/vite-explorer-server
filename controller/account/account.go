package account

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func Detail(w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "account.detail")
}