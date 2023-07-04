package model

import "time"

type Depression struct {
	ID               int
	DiseaseName      string
	DrationOfIllness time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
