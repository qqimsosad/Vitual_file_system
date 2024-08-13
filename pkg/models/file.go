package models

import "time"

type File struct {
	Name        string
	Description string
	CreatedAt   time.Time
}
