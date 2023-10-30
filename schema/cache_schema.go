package schema

import "strconv"

// PendingOffer struct
type PendingOffer struct { //Payload
	Id        uint64 `bson:"_id,omitempty" json:"_id,omitempty" redis:"_id"`
	ReturnUrl string `bson:"return_url,omitempty" json:"return_url,omitempty" redis:"return_url"`
}

func (pendingOffer PendingOffer) Database() string {
	return "steambot"
}

func (pendingOffer PendingOffer) Collection() string {
	return "pendingOffers"
}

func (pendingOffer PendingOffer) Key() string {
	return strconv.FormatUint(pendingOffer.Id, 10)
}
