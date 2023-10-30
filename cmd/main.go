package main

import (
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/environment"
	"github.com/cbotte21/steambot-internal-go/internal"
	"github.com/cbotte21/steambot-internal-go/schema"
	"github.com/doctype/steam"
	"log"
	"net/http"
	"time"
)

func main() {
	// Verify environment variables exist
	environment.VerifyEnvVariable("username")
	environment.VerifyEnvVariable("password")
	environment.VerifyEnvVariable("sharedSecret")
	environment.VerifyEnvVariable("tradeChannel")

	pendingOffers := datastore.MongoClient[schema.PendingOffer]{}
	err := pendingOffers.Init()
	if err != nil {
		log.Fatalf("mongodb client initialization failed")
	}

	tradeRequests := datastore.RedisClient[schema.TradeRequest]{}
	tradeRequests.Init()

	// Initialize hive
	steambot := internal.NewSteamBot(&pendingOffers, &tradeRequests, environment.GetEnvVariable("tradeChannel"), getSteamClient())
	steambot.Listen()
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
