package internal

import (
	"context"
	"fmt"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/steambot-internal-go/pb"
	"github.com/cbotte21/steambot-internal-go/schema"
	"github.com/doctype/steam"
	"github.com/google/uuid"
	"log"
	"time"
)

const HandleOffersIntervalSeconds = 15

type SteamBot struct {
	PendingOffers *datastore.MongoClient[schema.PendingOffer]
	SteamClient   *steam.Session
	pb.UnimplementedSteamBotServiceServer
}

func NewSteamBot(pendingOffers *datastore.MongoClient[schema.PendingOffer], steamClient *steam.Session) SteamBot {
	// Create thread for handling trades
	go HandleIncomingTrades(pendingOffers, steamClient)
	return SteamBot{PendingOffers: pendingOffers, SteamClient: steamClient}
}

func (steambot *SteamBot) Create(context context.Context, createRequest *pb.CreateRequest) (*pb.CreateResponse, error) {
	confirmationCode := uuid.NewString()

	offer := steam.TradeOffer{ // TODO: Implement items
		RecvItems: nil,
		SendItems: nil,
		Message:   "Confirmation Code: " + confirmationCode,
	}

	err := steambot.SteamClient.SendTradeOffer(&offer, steam.SteamID(createRequest.Recipient), "")
	if err != nil {
		return &pb.CreateResponse{Status: false}, err
	}

	err = steambot.PendingOffers.Create(schema.PendingOffer{
		Id:        offer.ID,
		ReturnUrl: createRequest.Response,
	})
	if err != nil {
		// Cancel trade
		_ = steambot.SteamClient.CancelTradeOffer(offer.ID)
		return &pb.CreateResponse{Status: false}, err
	}

	return &pb.CreateResponse{Status: true, Confirmation: confirmationCode}, nil
}

func HandleIncomingTrades(pendingOffers *datastore.MongoClient[schema.PendingOffer], steamClient *steam.Session) {
	resp, err := steamClient.GetTradeOffers(
		steam.TradeFilterRecvOffers,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, offer := range resp.ReceivedOffers {
		fmt.Println(offer)
		// If exists accept, otherwise decline
		_, err := pendingOffers.Find(schema.PendingOffer{Id: offer.ID})
		if err == nil {
			err := steamClient.AcceptTradeOffer(offer.ID)
			if err == nil {
				_ = pendingOffers.Delete(schema.PendingOffer{Id: offer.ID})
			}
		} else {
			_ = steamClient.DeclineTradeOffer(offer.ID)
		}
	}

	time.Sleep(time.Second * HandleOffersIntervalSeconds)
	HandleIncomingTrades(pendingOffers, steamClient)
}
