package model

type AddressModel struct {
	Model
	UserId        uint32 `gorm:"size:32;not null;column:user_id" json:"user_id"`
	StreetAddress string `gorm:"size:32;not null;column:street_address" json:"street_address"`
	City          string `gorm:"size:32;not null;column:city" json:"city"`
	State         string `gorm:"size:32;column:state" json:"state"`
	Country       string `gorm:"size:32;not null;column:country" json:"country"`
	ZipCode       int32  `gorm:"size:32;default:0;not null;column:zipcode" json:"zipcode"`
}
