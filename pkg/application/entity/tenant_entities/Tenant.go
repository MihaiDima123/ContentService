package tenant_entities

import "time"

type Tenant struct {
	ID        int
	tenantRef string
	channel   string
	language  string
	createdAt string
	updatedAt time.Time
}
