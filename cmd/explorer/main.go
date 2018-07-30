package main

import (
	"flag"
	"github.com/vitelabs/go-vite/config"
	"github.com/vitelabs/go-vite/vite"
	"log"
	"github.com/vitelabs/vite-explorer-server"
)

var (
	nameFlag  = flag.String("name", "", "boot name")
	sigFlag   = flag.String("sig", "", "boot sig")
	maxPeers  = flag.Uint("maxpeers", 0, "max number of connections will be connected")
	passRatio = flag.Uint("passration", 0, "max passive connections will be connected")

	minerFlag     = flag.Bool("miner", false, "boot miner")
	minerInterval = flag.Int("minerInterval", 6, "miner interval(unit sec).")
	coinbaseFlag  = flag.String("coinbaseAddress", "", "boot coinbaseAddress")
	env = flag.String( "env", "dev", "env info")
)

func main() {

	flag.Parse()

	globalConfig := config.GlobalConfig

	globalConfig.P2P = config.MergeP2PConfig(&config.P2P{
		Name:                 *nameFlag,
		Sig:                  *sigFlag,
		MaxPeers:             uint32(*maxPeers),
		MaxPassivePeersRatio: uint32(*passRatio),
	})
	globalConfig.P2P.Datadir = globalConfig.DataDir

	globalConfig.Miner = config.MergeMinerConfig(&config.Miner{
		Miner:         *minerFlag,
		Coinbase:      *coinbaseFlag,
		MinerInterval: *minerInterval,
	})

	vite, err := vite.New(globalConfig)

	if err != nil {
		log.Fatalf("Start vue failed. Error is %v\n", err)
	}

	vite_explorer_server.StartUp(*env, vite)
}
