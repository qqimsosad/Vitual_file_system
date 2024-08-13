package models

type User struct {
	Username string
	Folders  map[string]*Folder
}
