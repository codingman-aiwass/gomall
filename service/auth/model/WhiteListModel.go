package model

type WhiteListModel struct {
	Id   uint32 `gorm:"size:32;primaryKey;autoIncrement;column:id" json:"id"`
	Path string `gorm:"size:255;not null;column:path" json:"path"`
}
