package dto

type ReactionRequest struct {
	UserID     string `json:"userId"`
	ReactionID int    `json:"reactionId"` // Cambialo a int
}

type ReactionResponse struct {
	InteractionID int    `json:"interactionId"` // Cambialo a int
	UserID        string `json:"userId"`
	ReactionID    int    `json:"reactionId"`    // Cambialo a int
	Action        string `json:"action"`
}