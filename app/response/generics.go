package response

type Error struct {
	Message string `json:"message"`
	Error   bool   `json:"error" default:"true"`
}

type Success struct {
	Message string `json:"message"`
	Error   bool   `json:"error" default:"false"`
	Access  string `json:"access_token,omitempty"`
	Refresh string `json:"refresh_token,omitempty"`
}

type TokenData struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
