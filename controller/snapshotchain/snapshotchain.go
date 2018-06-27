package snapshotchain

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func BlockList (w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "snapshotchain.blocklist")
}

func Block (w http.ResponseWriter, r *http.Request, _ httprouter.Params)  {
	fmt.Fprint(w, "snapshotchain.block")
}
