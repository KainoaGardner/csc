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

type PostGame struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Money     [2]int `json:"money"`
	StartTime [2]int `json:"startTime"`
	PlaceLine int    `json:"placeLine"`
}

type PostState struct {
	State int `json:"state"`
}
