package models

type commentStatus string

const (
	published commentStatus = "public"
	deleted   commentStatus = "removed"
)

type Comment struct {
	User     User          `json:"user"`
	UserID   uint          `validate:"required"`
	Upload   Upload        `json:"upload"`
	UploadID uint          `validate:"required"`
	Text     string        `json:"text"`
	Status   commentStatus `gorm:"default:'public'"`
}
