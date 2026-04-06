package service

type AuthentikUserInfo struct {
	ID                string   `json:"sub"`
	Email             string   `json:"email"`
	EmailVerified     bool     `json:"emailVerified"`
	Name              string   `json:"name"`
	DisplayName       string   `json:"givenName"`
	PreferredUsername string   `json:"preferredUsername"`
	Nickname          string   `json:"nickname"`
	Groups            []string `json:"groups"`
}
