package main

import (
	"flag"
	"github.com/vitelabs/go-vite/config"
	"github.com/vitelabs/go-vite/vite"
	"github.com/vitelabs/vite-explorer-server"
	"github.com/vitelabs/go-vite/log15"
)

var (
	env = flag.String( "env", "dev", "env info")
)

func parseConfig() *config.Config {
	var globalConfig = config.GlobalConfig

	flag.StringVar(&globalConfig.Name, "name", globalConfig.Name, "boot name")
	flag.UintVar(&globalConfig.MaxPeers, "peers", globalConfig.MaxPeers, "max number of connections will be connected")
	flag.StringVar(&globalConfig.Addr, "addr", globalConfig.Addr, "will be listen by vite")
	flag.StringVar(&globalConfig.PrivateKey, "priv", globalConfig.PrivateKey, "hex encode of ed25519 privateKey, use for sign message")
	flag.StringVar(&globalConfig.DataDir, "dir", globalConfig.DataDir, "use for store all files")
	flag.UintVar(&globalConfig.NetID, "netid", globalConfig.NetID, "the network vite will connect")

	flag.Parse()

	globalConfig.P2P.Datadir = globalConfig.DataDir

	return globalConfig
}

func main() {

	mainLog := log15.New("module", "gvite/main")

	flag.Parse()

	parsedConfig := parseConfig()

	if s, e := parsedConfig.RunLogDirFile(); e == nil {
		log15.Root().SetHandler(
			log15.LvlFilterHandler(log15.LvlInfo, log15.Must.FileHandler(s, log15.TerminalFormat())),
		)
	}

	vite, err := vite.New(parsedConfig)

	if err != nil {
		mainLog.Crit("Start vite failed.", "err", err)
	}

	vite_explorer_server.StartUp(*env, vite)
}
