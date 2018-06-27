package accountchain

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func BlockList (w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "accountchain.blocklist")
}

func Block (w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "accountchain.block")
}

