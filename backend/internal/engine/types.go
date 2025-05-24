package engine

type Board struct {
	Width  int     `bson:"width" json:"width"`
	Height int     `bson:"height" json:"height"`
	Board  [][]int `bson:"board" json:"board"`
}

type Game struct {
	ID        int                `bson:"_id.omitempty" json:"id"`
	Board     Board              `bson:"board" json:"board"`
	WhiteID   int                `bson:"white_id" json:"white_id"`
	BlackID   int                `bson:"black_id" json:"black_id"`
	Turn      int                `bson:"turn" json:"turn"`
	Mochigoma [MochigomaSize]int `bson:"mochigoma" json:"mochigoma"` //turn 0=0-6  turn 1=7-13 | order 歩香桂銀金角飛
	Moves     []string           `bson:"moves" json:"moves"`
}

type Move struct {
	Start        [2]int
	End          [2]int
	MovePiece    int
	TakenPiece   int
	Promote      bool
	PromotePiece int
	Drop         bool
	DropPiece    int
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
	'P': MochiFu,
	'L': MochiKyou,
	'N': MochiKei,
	'S': MochiGin,
	'G': MochiKin,
	'B': MochiKaku,
	'R': MochiHi,
}

var shogiDropPieceToChar = map[int]byte{
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
