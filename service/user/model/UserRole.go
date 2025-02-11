package model

type UserRole struct {
	ID     uint32 `gorm:"primaryKey;autoIncrement;column:id"`
	UserID uint32 `gorm:"not null;index;column:user_id"`
	RoleID uint32 `gorm:"not null;index;column:role_id"`

	User UserModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role RoleModel `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}
