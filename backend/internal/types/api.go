package types

type PostMove struct {
	Move string `json:"move"`
}

type PostPlace struct {
	Position string `json:"position"`
	Type     int    `json:"type"`
}

type PostGame struct {
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Money     [2]int   `json:"money"`
	StartTime [2]int64 `json:"startTime"`
	PlaceLine int      `json:"placeLine"`
}
