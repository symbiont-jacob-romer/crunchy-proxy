package main

import (
	"github.com/symbiont-jacob-romer/crunchy-proxy/common"
	"github.com/symbiont-jacob-romer/crunchy-proxy/config"
	"github.com/symbiont-jacob-romer/crunchy-proxy/server"
	"github.com/symbiont-jacob-romer/crunchy-proxy/util/log"
)

func main() {
	log.SetLevel("info")

	serverConfig := config.ServerConfig{
		Admin: config.AdminConfig{HostPort: "127.0.0.1:5434"},
		Proxy: config.ProxyConfig{HostPort: "127.0.0.1:5433"},
	}
	creds := common.Credentials{
		Username: "txe",
		Password: "password",
		Database: "state_db",
		SSL:      common.SSLConfig{
			Enable: false,
			SSLMode: "disable",
		},
	}
	nodes := map[string]common.Node{
		"master": {
			HostPort: "127.0.0.1:5432",
		},
	}

	config.Set("Credentials", creds)
	config.Set("Nodes", nodes)
	config.Set("Server", serverConfig)
	config.Set("HealthCheck", common.HealthCheckConfig{
		Delay: 10000,
	})
	config.SetConfig()
	log.Infof("conf creds %#v", config.GetConfig())

	s := server.NewServer()
	s.Start()
}
