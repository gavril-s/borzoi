package api

type Error struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
}
