package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	controllerAccount "vite-explorer-server/controller/account"
	controllerAccountchain "vite-explorer-server/controller/accountchain"
	controllerSnapshotchain "vite-explorer-server/controller/snapshotchain"

	controllerToken "vite-explorer-server/controller/token"
)

var (
	port = "8081"
)

func registerAccountRouter()  {
	router := httprouter.New()

	router.GET("/detail", controllerAccount.Detail)

	http.Handle("/api/account/", http.StripPrefix("/api/account", router))

}

func registerAccountChainRouter()  {
	router := httprouter.New()

	router.POST("/blocklist", controllerAccountchain.BlockList)
	router.GET("/block", controllerAccountchain.Block)

	http.Handle("/api/accountchain/",  http.StripPrefix("/api/accountchain", router))
}

func registerSnapshotChainRouter()  {
	router := httprouter.New()

	router.POST("/blocklist", controllerSnapshotchain.BlockList)
	router.GET("/block", controllerSnapshotchain.Block)

	http.Handle("/api/snapshotchain/", http.StripPrefix("/api/snapshotchain", router))
}

func registerTokenRouter () {
	router := httprouter.New()

	router.POST("/list", controllerToken.List)

	http.Handle("/api/token/", http.StripPrefix("/api/token", router))
}


func main ()  {
	registerAccountRouter()

	registerAccountChainRouter()

	registerSnapshotChainRouter()

	registerTokenRouter()

	log.Println("Server listen in " + port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
