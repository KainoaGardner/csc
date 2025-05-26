package engine

type Piece struct {
	Type  int
	Owner int
	Moved bool
}

type Board struct {
	Width  int       `bson:"width" json:"width"`
	Height int       `bson:"height" json:"height"`
	Board  [][]Piece `bson:"board" json:"board"`
}

type Game struct {
	ID        int                `bson:"_id.omitempty" json:"id"`
	Board     Board              `bson:"board" json:"board"`
	WhiteID   int                `bson:"white_id" json:"white_id"`
	BlackID   int                `bson:"black_id" json:"black_id"`
	Mochigoma [MochigomaSize]int `bson:"mochigoma" json:"mochigoma"` //turn 0=0-6  turn 1=7-13 | order 歩香桂銀金角飛
	Turn      int                `bson:"turn" json:"turn"`
	MoveCount int
	LastMove  Move
	Moves     []string `bson:"moves" json:"moves"`
}

type Move struct {
	Start      Vec2
	End        Vec2
	StartPiece Piece
	EndPiece   Piece
	TakenPiece Piece
	Promote    bool
	Drop       bool
}

type Vec2 struct {
	x int
	y int
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
