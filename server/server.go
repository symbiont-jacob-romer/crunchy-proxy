/*
Copyright 2017 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/symbiont-jacob-romer/crunchy-proxy/common"
	"github.com/symbiont-jacob-romer/crunchy-proxy/config"
	"github.com/symbiont-jacob-romer/crunchy-proxy/util/log"
)

type Server struct {
	proxy     *ProxyServer
	waitGroup *sync.WaitGroup
}

func NewServer() *Server {
	s := &Server{
		waitGroup: &sync.WaitGroup{},
	}

	s.proxy = NewProxyServer(s)

	return s
}

func (s *Server) Start() {
	proxyConfig := config.GetProxyConfig()

	log.Info("Proxy Server Starting...")
	proxyListener, err := net.Listen("tcp", proxyConfig.HostPort)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	s.waitGroup.Add(1)
	go s.proxy.Serve(proxyListener)

	s.waitGroup.Wait()

	log.Info("Server Exiting...")
}

func StartServer(host, port, user, pass, database, sslMode, ca, proxyPort string) {
	log.SetLevel("debug")

	serverConfig := config.ServerConfig{
		Proxy: config.ProxyConfig{HostPort: fmt.Sprintf("127.0.0.1:%s", proxyPort)},
	}
	creds := common.Credentials{
		Username: user,
		Password: pass,
		Database: database,
		SSL: common.SSLConfig{
			SSLMode:   sslMode,
			SSLRootCA: ca,
		},
	}
	nodes := map[string]common.Node{
		"master": {
			HostPort: fmt.Sprintf("%s:%s", host, port),
			Role: "master",
		},
	}

	config.Set("Credentials", creds)
	config.Set("Nodes", nodes)
	config.Set("Server", serverConfig)
	config.Set("Pool", config.PoolConfig{Capacity: 2})
	config.Set("HealthCheck", common.HealthCheckConfig{Delay: 10000})
	config.SetConfig()
	log.Infof("conf creds %#v", config.GetConfig())

	s := NewServer()
	s.Start()
}
