package engine

import (
	"github.com/KainoaGardner/csc/internal/types"
)

type Piece struct {
	Type  int
	Owner int
	Moved bool
}

type Board struct {
	Width  int        `bson:"width" json:"width"`
	Height int        `bson:"height" json:"height"`
	Board  [][]*Piece `bson:"board" json:"board"`
}

type Game struct {
	ID            int                `bson:"_id.omitempty" json:"id"`
	Board         Board              `bson:"board" json:"board"`
	WhiteID       int                `bson:"white_id" json:"white_id"`
	BlackID       int                `bson:"black_id" json:"black_id"`
	Mochigoma     [MochigomaSize]int `bson:"mochigoma" json:"mochigoma"` //turn 0=0-6  turn 1=7-13 | order 歩香桂銀金角飛
	Turn          int                `bson:"turn" json:"turn"`
	MoveCount     int
	HalfMoveCount int
	Moves         []string `bson:"moves" json:"moves"`
	EnPassant     *types.Vec2
	CheckerJump   bool
}

type Move struct {
	Start   types.Vec2
	End     types.Vec2
	Promote *int
	Drop    *int
}

const (
	MochiFu = iota
	MochiKyou
	MochiKei
	MochiGin
	MochiKin
	MochiKaku
	MochiHi
)

const (
	MochigomaSize        = 14
	MochigomaBlackOffset = 7
)

const (
	Empty = iota
	Pawn
	Knight
	Bishop
	Rook
	Queen
	King
	Fu
	Kyou
	Kei
	Gin
	Kin
	Kaku
	Hi
	Ou
	To
	NariKyou
	NariKei
	NariGin
	Uma
	Ryuu
	Checker
	CheckerKing
)

var shogiDropCharToPiece = map[byte]int{
	'P': Fu,
	'L': Kyou,
	'N': Kei,
	'S': Gin,
	'G': Kin,
	'B': Kaku,
	'R': Hi,
}

var shogiDropPieceToChar = map[int]byte{
	Fu:   'P',
	Kyou: 'L',
	Kei:  'N',
	Gin:  'S',
	Kin:  'G',
	Kaku: 'B',
	Hi:   'R',
}

var shogiDropPieceToMochiPiece = map[int]int{
	Fu:       MochiFu,
	Kyou:     MochiKyou,
	Kei:      MochiKei,
	Gin:      MochiGin,
	Kin:      MochiKin,
	Kaku:     MochiKaku,
	Hi:       MochiHi,
	To:       MochiFu,
	NariKyou: MochiKyou,
	NariKei:  MochiKei,
	NariGin:  MochiGin,
	Uma:      MochiKaku,
	Ryuu:     MochiHi,
}

var shogiMochiPieceToDropPiece = map[int]int{
	MochiFu:   Fu,
	MochiKyou: Kyou,
	MochiKei:  Kei,
	MochiGin:  Gin,
	MochiKin:  Kin,
	MochiKaku: Kaku,
	MochiHi:   Hi,
}

var shogiDropCharToMochiPiece = map[byte]int{
	'P': MochiFu,
	'L': MochiKyou,
	'N': MochiKei,
	'S': MochiGin,
	'G': MochiKin,
	'B': MochiKaku,
	'R': MochiHi,
}

var shogiMochiPieceToChar = map[int]byte{
	MochiFu:   'P',
	MochiKyou: 'L',
	MochiKei:  'N',
	MochiGin:  'S',
	MochiKin:  'G',
	MochiKaku: 'B',
	MochiHi:   'R',
}

var chessPromotePieceToChar = map[int]byte{
	Pawn:   'P',
	Knight: 'N',
	Bishop: 'B',
	Rook:   'R',
	Queen:  'Q',
}

var chessPromoteCharToPiece = map[byte]int{
	'P': Pawn,
	'N': Knight,
	'B': Bishop,
	'R': Rook,
	'Q': Queen,
}

var fenPieceToString = map[int]string{
	Pawn:        "CP",
	Knight:      "CN",
	Bishop:      "CB",
	Rook:        "CR",
	Queen:       "CQ",
	King:        "CK",
	Fu:          "SP",
	Kyou:        "SL",
	Kei:         "SN",
	Gin:         "SG",
	Kin:         "SC",
	Kaku:        "SB",
	Hi:          "SR",
	Ou:          "SK",
	To:          "NP",
	NariKyou:    "NL",
	NariKei:     "NN",
	NariGin:     "NG",
	Uma:         "NB",
	Ryuu:        "NR",
	Checker:     "KC",
	CheckerKing: "KK",
}
