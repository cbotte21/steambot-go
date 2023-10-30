package schema

import "strconv"

type Item struct {
	AssetID    string `bson:"assetId,omitempty" json:"assetId,omitempty" redis:"assetId"`
	InstanceID int64  `bson:"instanceId,omitempty" json:"instanceId,omitempty" redis:"instanceId"`
	ClassID    int64  `bson:"classId,omitempty" json:"classId,omitempty" redis:"classId"`
	AppID      int32  `bson:"appId,omitempty" json:"appId,omitempty" redis:"appId"`
	ContextID  int64  `bson:"contextId,omitempty" json:"contextId,omitempty" redis:"contextId"`
	Amount     uint32 `bson:"amount" json:"amount,omitempty" redis:"amount"`
	Missing    int32  `bson:"missing,omitempty" json:"missing,omitempty" redis:"missing"`
}

// TradeRequest struct
type TradeRequest struct { //Payload
	Id          int64  `bson:"_id,omitempty" json:"_id,omitempty" redis:"_id"`
	PlayerItems []Item `bson:"playerItems,omitempty" json:"playerItems,omitempty" redis:"playerItems"`
	BotItems    []Item `bson:"botItems,omitempty" json:"botItems,omitempty" redis:"botItems"`
	Recipient   int64  `bson:"recipient,omitempty" json:"recipient,omitempty" redis:"recipient"`
	Response    string `bson:"response,omitempty" json:"response,omitempty" redis:"response"`
}

func (tradeRequest TradeRequest) Database() string {
	return "steambot"
}

func (tradeRequest TradeRequest) Collection() string {
	return "tradeRequest"
}

func (tradeRequest TradeRequest) Key() string {
	return strconv.FormatInt(tradeRequest.Id, 10)
}
