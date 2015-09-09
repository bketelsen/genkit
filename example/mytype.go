//go:generate genkit $GOFILE
package example

import "time"

// @service
type User struct {
	ID                  int64
	Name                string
	Website             string
	PrimaryContact      string
	PrimaryContactEmail string
	PrimaryContactPhone string
	CreatedAt           time.Time
}
