package types

type APIRespone struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PostMove struct {
	Move string `json:"move"`
	Turn int    `json:"turn"`
}

type PostMoveResponse struct {
	FEN  string `json:"fen"`
	Move string `json:"move"`
}

type PostPlace struct {
	Position string `json:"position"`
	Type     int    `json:"type"`
	Turn     int    `json:"turn"` //TEMP USE SENT ID INSTEAD
}

type DeletePlace struct {
	Position string `json:"position"`
	Turn     int    `json:"turn"` //TEMP USE SENT ID INSTEAD
}

type PlaceResponse struct {
	FEN      string `json:"fen"`
	Position string `json:"position"`
	Type     int    `json:"type"`
	Cost     int    `json:"cost"`
	Money    [2]int `json:"money"`
}

type PostGame struct {
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	Money     [2]int   `json:"money"`
	StartTime [2]int64 `json:"startTime"`
	PlaceLine int      `json:"placeLine"`
}

type PostGameResponse struct {
	ID        string `json:"id"`
	WhiteID   string `json:"whiteID"`
	BlackID   string `json:"blackID"`
	Color     string `json:"color"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Money     [2]int `json:"money"`
	PlaceLine int    `json:"placeLine"`
}

type PostState struct {
	State int `json:"state"`
}

type PostReady struct {
	Turn int `json:"turn"`
}
