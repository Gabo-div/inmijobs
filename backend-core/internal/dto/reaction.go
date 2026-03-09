package dto

type ReactionRequest struct {
	UserID     string `json:"user_id"`
	ReactionID int    `json:"reaction_id"` 
}

type ReactionResponse struct {
	InteractionID int    `json:"interaction_id"` 
	UserID        string `json:"user_id"`
	ReactionID    int    `json:"reaction_id"`    
	Action        string `json:"action"`
}