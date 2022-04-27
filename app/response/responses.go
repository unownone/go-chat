package response

type Error struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

type Success struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
	Access  string `json:"access_token"`
}
