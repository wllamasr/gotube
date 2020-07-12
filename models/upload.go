package models

import "github.com/jinzhu/gorm"

type fileType string
type visibility string

const (
	video   fileType = "video"
	image   fileType = "image"
	audio   fileType = "audio"
	unknown fileType = "unknown"
	convert fileType = "convert"
)
const (
	public   visibility = "public"
	unlisted visibility = "unlisted"
	private  visibility = "private"
	removed  visibility = "removed"
	pending  visibility = "pending"
)

type Upload struct {
	gorm.Model
	Title            string    `json:"title" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	OriginalFileName string    `json:"original_file_name"`
	FileExtension    string    `json:"file_extension"`
	FileType         fileType `json:"file_type"`
	UploaderID       uint      `json:"uploader_id" validate:"required"`
}
