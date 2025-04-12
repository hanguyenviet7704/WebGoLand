package po

import "time"

type User struct {
	Id          int           `gorm:"primary_key:auto_increment" json:"id"`
	Name        string        `gorm:"size:255;not null;unique" json:"name"`
	Email       string        `gorm:"size:255;not null;unique"  json:"email"`
	Password    string        `json:"password"`
	Is_active   bool          `json:"is_active"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Roles       []Roles       `gorm:"many2many:user_roles"`
	Tokens      []Tokens      `gorm:"foreignkey:UserID"`
	Permissions []Permissions `gorm:"many2many:user_permissions"`
}
