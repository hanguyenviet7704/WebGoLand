package po

import "time"

type Permissions struct {
	Id        int       `json:"id"`
	Name      string    `gorm:"size:255" json:"name"`
	Resource  string    `gorm:"size:255" json:"resource"`
	Action    string    `gorm:"size:255" json:"action"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
