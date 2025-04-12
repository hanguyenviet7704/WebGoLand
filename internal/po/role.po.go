package po

import "time"

type Roles struct {
	ID          int           `gorm:"primary_key" json:"id"`
	Name        string        `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Permissions []Permissions `gorm:"many2many:roles_permissions;" json:"permissions"`
}
