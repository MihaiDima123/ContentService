package tenant_entities

import "time"

type Tenant struct {
	ID        int       `json:"id" gorm:"id"`
	TenantRef string    `json:"tenantRef" gorm:"tenantRef"`
	Channel   string    `json:"channel" gorm:"channel"`
	Language  string    `json:"language" gorm:"language"`
	CreatedAt time.Time `json:"createdAt" gorm:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"updatedAt"`
}
