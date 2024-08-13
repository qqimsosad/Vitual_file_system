package models

import "time"

type Folder struct {
	Name        string
	Description string
	Files       map[string]*File
	CreatedAt   time.Time
}
