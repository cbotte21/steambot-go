package internal

import (
	"fmt"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/steambot-internal-go/schema"
	"github.com/doctype/steam"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"time"
)

const HandleOffersIntervalSeconds = 15

type SteamBot struct {
	PendingOffers *datastore.MongoClient[schema.PendingOffer]
	TradeRequests *datastore.RedisClient[schema.TradeRequest]
	SteamClient   *steam.Session
	TradeChannel  string
}

func NewSteamBot(pendingOffers *datastore.MongoClient[schema.PendingOffer], tradeRequests *datastore.RedisClient[schema.TradeRequest], tradeChannel string, steamClient *steam.Session) SteamBot {
	return SteamBot{PendingOffers: pendingOffers, TradeRequests: tradeRequests, TradeChannel: tradeChannel, SteamClient: steamClient}
}

func (steambot *SteamBot) Listen() {
	go steambot.HandleIncomingTrades()
	steambot.HandleOutgoingTrades()
}

func (steambot *SteamBot) HandleOutgoingTrades() {
	sub := steambot.TradeRequests.Subscribe(steambot.TradeChannel)
	ch := sub.Channel()
	for request := range ch { // Create trade offer
		var payload schema.TradeRequest
		err := bson.Unmarshal([]byte(request.Payload), &payload)
		if err != nil {
			return
		}

		confirmationCode := uuid.NewString()
		offer := steambot.createOffer(confirmationCode)

		// Send offer
		err = steambot.SteamClient.SendTradeOffer(&offer, steam.SteamID(payload.Recipient), "")
		if err != nil {
			continue
		}

		// Create record
		err = steambot.PendingOffers.Create(schema.PendingOffer{
			Id:        offer.ID,
			ReturnUrl: payload.Response,
		})
		if err != nil { // Destroy offer if failed to save record
			// Cancel trade
			_ = steambot.SteamClient.CancelTradeOffer(offer.ID)
		}

		// Publish the confirmation code
		err = steambot.TradeRequests.Publish(strconv.FormatInt(payload.Recipient, 10), confirmationCode)
		if err != nil { // Delete request if failed to publish
			_ = steambot.SteamClient.CancelTradeOffer(offer.ID)
		}
	}
}

func (steambot *SteamBot) createOffer(confirmationCode string) steam.TradeOffer {
	return steam.TradeOffer{ // TODO: Implement items
		RecvItems: nil,
		SendItems: nil,
		Message:   "Confirmation Code: " + confirmationCode,
	}
}

func (steambot *SteamBot) HandleIncomingTrades() {
	resp, err := steambot.SteamClient.GetTradeOffers(
		steam.TradeFilterRecvOffers,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, offer := range resp.ReceivedOffers {
		fmt.Println(offer)
		// If exists accept, otherwise decline
		_, err := steambot.PendingOffers.Find(schema.PendingOffer{Id: offer.ID})
		if err == nil {
			err := steambot.SteamClient.AcceptTradeOffer(offer.ID)
			if err == nil {
				_ = steambot.PendingOffers.Delete(schema.PendingOffer{Id: offer.ID})
			}
		} else {
			_ = steambot.SteamClient.DeclineTradeOffer(offer.ID)
		}
	}

	time.Sleep(time.Second * HandleOffersIntervalSeconds)
	steambot.HandleIncomingTrades()
}
