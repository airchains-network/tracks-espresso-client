package types

type Response struct {
	Status      bool        `json:"status"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Data        interface{} `json:"data ,omitempty"`
}
