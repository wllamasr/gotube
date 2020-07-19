package models

type Upload struct {
	Model
	Title            string `form:"title" json:"title" validate:"required"`
	Description      string `form:"description" json:"description" gorm:"default:''"`
	FileName         string
	OriginalFileName string
	Visibility       string `form:"visibility" json:"visibility" gorm:"default:'pending';type:enum('public', 'unlisted','private', 'removed', 'pending')" validate:"oneof=public unlisted private removed pending"`
	Views            uint64 `json:"views" gorm:"default:0"`
	PublicURL        string `json:"public_url" gorm:"unique;not null"`
	UserID           uint   `json:"user_id" validate:"required"`
	User             User   `json:"user"`
}

func (upload *Upload) BeforeCreate() {
	upload.Views = 0
}

func (upload *Upload) BeforeSave() {
}
