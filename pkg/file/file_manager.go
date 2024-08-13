package file

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"virtual-file-system/pkg/models"
	"virtual-file-system/pkg/user"
)

func CreateFile(username, foldername, filename, description string) error {
	usr, err := user.GetUser(username)
	if err != nil {
		return err
	}

	folder, exists := usr.Folders[strings.ToLower(foldername)]
	if !exists {
		return fmt.Errorf("the %s folder doesn't exist", foldername)
	}

	filename = strings.ToLower(filename)
	if err := models.ValidateName(foldername); err != nil {
		return err
	}
	if _, exists := folder.Files[filename]; exists {
		return fmt.Errorf("the %s has already existed", filename)
	}

	folder.Files[filename] = &models.File{
		Name:        filename,
		Description: description,
		CreatedAt:   time.Now(),
	}
	return nil
}

func DeleteFile(username, foldername, filename string) error {
	usr, err := user.GetUser(username)
	if err != nil {
		return err
	}

	foldername = strings.ToLower(foldername)
	filename = strings.ToLower(filename)

	folder, exists := usr.Folders[foldername]
	if !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}

	if _, exists := folder.Files[filename]; !exists {
		return fmt.Errorf("the %s doesn't exist ", filename)
	}

	delete(folder.Files, filename)

	return nil
}

func ListFiles(username, foldername, sortBy, order string) ([]*models.File, error) {
	usr, err := user.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("error: The %s doesn't exist", username)
	}

	foldername = strings.ToLower(foldername)
	folder, exists := usr.Folders[foldername]
	if !exists {
		return nil, fmt.Errorf("error: the %s doesn't exist", foldername)
	}

	fileList := make([]*models.File, 0, len(folder.Files))
	for _, file := range folder.Files {
		fileList = append(fileList, file)
	}

	switch sortBy {
	case "name":
		sort.Slice(fileList, func(i, j int) bool {
			if order == "desc" {
				return fileList[i].Name > fileList[j].Name
			}
			return fileList[i].Name < fileList[j].Name
		})
	case "created":
		sort.Slice(fileList, func(i, j int) bool {
			if order == "desc" {
				return fileList[i].CreatedAt.After(fileList[j].CreatedAt)
			}
			return fileList[i].CreatedAt.Before(fileList[j].CreatedAt)
		})
	default:
		sort.Slice(fileList, func(i, j int) bool {
			return fileList[i].Name < fileList[j].Name
		})
	}

	return fileList, nil
}
