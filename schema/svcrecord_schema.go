package schema

// SVCRecord struct
type SVCRecord struct { //Payload
	Player string `bson:"player,omitempty" json:"player,omitempty" redis:"player"`
	Game   string `bson:"game,omitempty" json:"game,omitempty" redis:"game"`
}

func (svcRecord SVCRecord) Database() string {
	return ""
}

func (svcRecord SVCRecord) Collection() string {
	return ""
}

func (svcRecord SVCRecord) Key() string {
	return svcRecord.Player
}
