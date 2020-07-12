package models

type commentStatus string

const (
	published commentStatus = "public"
	deleted   commentStatus = "removed"
)

type Comment struct {
	Commenter   User          `json:"user"`
	CommenterID uint          `validate:"required"`
	Upload      Upload        `json:"upload"`
	UploadID    uint          `validate:"required"`
	Text        string        `json:"text"`
	Status      commentStatus `gorm:"default:'public'"`
}
