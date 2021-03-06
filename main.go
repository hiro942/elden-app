package main

import (
	"github.com/hiro942/elden-app/global"
	"github.com/hiro942/elden-app/router"
	"github.com/hiro942/elden-app/utils"
	"log"
	"time"
)

func main() {
	time.Local = time.FixedZone("CST", 8*3600)

	// create gRPC connection to the PEER
	clientConnection := utils.NewGrpcConnection(global.TlsCertPath, global.GatewayPeer, global.PeerEndpoint)
	defer clientConnection.Close()

	// get gateway
	gateway, err := utils.NewGateway(clientConnection)
	if err != nil {
		log.Panicln(err)
	}
	defer gateway.Close()

	// get contract from the FABRIC NET
	global.Contract = utils.GetContract(gateway)

	// router start
	sysRouter := router.Routers()
	sysRouter.Run(":8080")
}
