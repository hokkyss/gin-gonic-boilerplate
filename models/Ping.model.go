package models

type PingResponse struct {
	Pong string `json:"pong" bson:"pong"`
}
