package schema

// RecordedGame struct
type RecordedGame struct { //Payload
	Viktor    string `bson:"viktor,omitempty" json:"viktor,omitempty"`
	Opponent  string `bson:"opponent,omitempty" json:"opponent,omitempty"`
	Timestamp int32  `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
}

func (game RecordedGame) Database() string {
	return ""
}

func (game RecordedGame) Collection() string {
	return ""
}

func (game RecordedGame) Key() string {
	return game.Viktor
}
