package domain

type Node struct {
	Port     int  `json:"port,omitempty"`
	IsActive bool `json:"is_active,omitempty"`
}
