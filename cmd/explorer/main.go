package main

import (
	"flag"
	"github.com/vitelabs/go-vite/config"
	"github.com/vitelabs/go-vite/vite"
	"github.com/vitelabs/vite-explorer-server"
	"github.com/vitelabs/go-vite/log15"
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

	mainLog := log15.New("module", "gvite/main")

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

	if s, e := config.GlobalConfig.RunLogDirFile(); e == nil {
		log15.Root().SetHandler(
			log15.LvlFilterHandler(log15.LvlInfo, log15.Must.FileHandler(s, log15.TerminalFormat())),
		)
	}

	vite, err := vite.New(globalConfig)

	if err != nil {
		mainLog.Crit("Start vite failed.", "err", err)
	}

	vite_explorer_server.StartUp(*env, vite)
}
