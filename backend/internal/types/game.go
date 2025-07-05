package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Piece struct {
	Type  int  `bson:"type" json:"type"`
	Owner int  `bson:"owner" json:"owner"`
	Moved bool `bson:"moved" json:"moved"`
}

type Board struct {
	Width     int        `bson:"width" json:"width"`
	Height    int        `bson:"height" json:"height"`
	PlaceLine int        `bson:"placeLine" json:"placeLine"`
	Board     [][]*Piece `bson:"board" json:"board"`
}

type Game struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Board         Board              `bson:"board" json:"board"`
	WhiteID       string             `bson:"whiteID" json:"whiteID"`
	BlackID       string             `bson:"blackID" json:"blackID"`
	Mochigoma     [MochigomaSize]int `bson:"mochigoma" json:"mochigoma"` //turn 0=0-6  turn 1=7-13 | order 歩香桂銀金角飛
	Turn          int                `bson:"turn" json:"turn"`
	MoveCount     int                `bson:"moveCount" json:"moveCount"`
	HalfMoveCount int                `bson:"halfMoveCount" json:"halfMoveCount"`
	EnPassant     *Vec2              `bson:"enPassant" json:"enPassant"`
	CheckerJump   *Vec2              `bson:"checkerJump" json:"checkerJump"`
	Winner        *int               `bson:"winner" json:"winner"`
	State         int                `bson:"state" json:"state"` //0 place,1 move,2 over
	Time          [2]int             `bson:"time" json:"time"`
	LastMoveTime  time.Time          `bson:"lastMoveTime" json:"lastMoveTime"`
	Money         [2]int             `bson:"money" json:"money"`
}

/*
board
mochigoma
turn
moveCount
enPassant
CheckerJump
time
*/

type GameLogs struct {
	ID          primitive.ObjectID `bson:"_id.omitempty" json:"_id.omitempty"`
	BoardStates []string           `bson:"boardStates" json:"boardStates"`
}

const (
	MochigomaSize        = 14
	MochigomaBlackOffset = 7
)

type Move struct {
	Start   Vec2
	End     Vec2
	Promote *int
	Drop    *int
}

type Place struct {
	Pos  Vec2
	Type int
	Turn int
	Cost int
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
	White = iota
	Black
	Tie
)

const ( //states
	ConnectState = iota
	PlaceState
	MoveState
	OverState
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

var ShogiDropCharToPiece = map[byte]int{
	'P': Fu,
	'L': Kyou,
	'N': Kei,
	'S': Gin,
	'G': Kin,
	'B': Kaku,
	'R': Hi,
}

var ShogiDropPieceToChar = map[int]byte{
	Fu:   'P',
	Kyou: 'L',
	Kei:  'N',
	Gin:  'S',
	Kin:  'G',
	Kaku: 'B',
	Hi:   'R',
}

var ShogiDropPieceToMochiPiece = map[int]int{
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

var ShogiMochiPieceToDropPiece = map[int]int{
	MochiFu:   Fu,
	MochiKyou: Kyou,
	MochiKei:  Kei,
	MochiGin:  Gin,
	MochiKin:  Kin,
	MochiKaku: Kaku,
	MochiHi:   Hi,
}

var ShogiDropCharToMochiPiece = map[byte]int{
	'P': MochiFu,
	'L': MochiKyou,
	'N': MochiKei,
	'S': MochiGin,
	'G': MochiKin,
	'B': MochiKaku,
	'R': MochiHi,
}

var ShogiMochiPieceToChar = map[int]byte{
	MochiFu:   'P',
	MochiKyou: 'L',
	MochiKei:  'N',
	MochiGin:  'S',
	MochiKin:  'G',
	MochiKaku: 'B',
	MochiHi:   'R',
}

var ChessPromotePieceToChar = map[int]byte{
	Pawn:   'P',
	Knight: 'N',
	Bishop: 'B',
	Rook:   'R',
	Queen:  'Q',
}

var ChessPromoteCharToPiece = map[byte]int{
	'P': Pawn,
	'N': Knight,
	'B': Bishop,
	'R': Rook,
	'Q': Queen,
}

var FenPieceToString = map[int]string{
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

var PieceToCost = map[int]int{
	King: 50,
	Ou:   45,

	Pawn:   3,
	Knight: 20,
	Bishop: 25,
	Rook:   30,
	Queen:  50,

	Fu:   3,
	Kyou: 8,
	Kei:  12,
	Gin:  15,
	Kin:  20,
	Kaku: 28,
	Hi:   35,

	Checker: 10,
}
