package schema

// CachedGame struct
type CachedGame struct { //Payload
	PlayerOne string `bson:"playerOne,omitempty" json:"playerOne,omitempty" redis:"playerOne"`
	PlayerTwo string `bson:"playerTwo,omitempty" json:"playerTwo,omitempty" redis:"playerTwo"`
	Board     string `bson:"board,omitempty" json:"board,omitempty" redis:"board"`
	Ranked    bool   `bson:"ranked,omitempty" json:"ranked,omitempty" redis:"ranked"`
	Turn      bool   `bson:"turn,omitempty" json:"turn,omitempty" redis:"turn"`
}

func (game CachedGame) Database() string {
	return ""
}

func (game CachedGame) Collection() string {
	return ""
}

func (game CachedGame) Key() string {
	return game.PlayerOne
}
