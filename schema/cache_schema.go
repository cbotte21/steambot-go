package schema

// CachedGame struct
type CachedGame struct { //Payload
	White      string `bson:"white,omitempty" json:"white,omitempty" redis:"whiteOne"`
	Black      string `bson:"black,omitempty" json:"black,omitempty" redis:"blackTwo"`
	Pawns      uint64 `bson:"pawns,omitempty" json:"pawns,omitempty" redis:"pawns"`
	Knights    uint64 `bson:"knights,omitempty" json:"knights,omitempty" redis:"knights"`
	Rooks      uint64 `bson:"rooks,omitempty" json:"rooks,omitempty" redis:"rooks"`
	Bishops    uint64 `bson:"bishops,omitempty" json:"bishops,omitempty" redis:"bishops"`
	Queens     uint64 `bson:"queens,omitempty" json:"queens,omitempty" redis:"queens"`
	Kings      uint64 `bson:"kings,omitempty" json:"kings,omitempty" redis:"kings"`
	P1BitBoard uint64 `bson:"p1_bitboard,omitempty" json:"p1_bitboard,omitempty" redis:"p1_bitboard"`
	Enpassants uint64 `bson:"enpassants,omitempty" json:"enpassants,omitempty" redis:"enpassants"`
	Ranked     bool   `bson:"ranked,omitempty" json:"ranked,omitempty" redis:"ranked"`
	Turn       bool   `bson:"turn,omitempty" json:"turn,omitempty" redis:"turn"`
}

func (game CachedGame) Database() string {
	return ""
}

func (game CachedGame) Collection() string {
	return ""
}

func (game CachedGame) Key() string {
	return game.White
}
