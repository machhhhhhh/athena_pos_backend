package testcontroller

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
} // @name LoginResponse
