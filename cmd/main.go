package main

import (
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/environment"
	"github.com/cbotte21/steambot-internal-go/internal"
	"github.com/cbotte21/steambot-internal-go/pb"
	"github.com/cbotte21/steambot-internal-go/schema"
	"github.com/doctype/steam"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	// Verify environment variables exist
	environment.VerifyEnvVariable("port")
	environment.VerifyEnvVariable("username")
	environment.VerifyEnvVariable("password")
	environment.VerifyEnvVariable("sharedSecret")

	port := environment.GetEnvVariable("port")

	// Setup tcp listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port: %s", port)
	}
	grpcServer := grpc.NewServer()

	pendingOffers := datastore.MongoClient[schema.PendingOffer]{}
	err = pendingOffers.Init()
	if err != nil {
		log.Fatalf("mongodb client initialization failed")
	}

	// Initialize hive
	steambot := internal.NewSteamBot(&pendingOffers, getSteamClient())
	pb.RegisterSteamBotServiceServer(grpcServer, &steambot)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf(err.Error())
	}
}

func getSteamClient() *steam.Session {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	timeTip, err := steam.GetTimeTip()
	if err != nil {
		log.Fatal(err)
	}

	timeDiff := time.Duration(timeTip.Time - time.Now().Unix())
	session := steam.NewSession(&http.Client{}, "")
	if err := session.Login(
		environment.GetEnvVariable("username"),
		environment.GetEnvVariable("password"),
		environment.GetEnvVariable("sharedSecret"),
		timeDiff,
	); err != nil {
		log.Fatal(err)
	}
	return session
}
