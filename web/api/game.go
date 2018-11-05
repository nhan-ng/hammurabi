package api

// Game represents a unique Hammurabi game for a specific user.
type Game struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
