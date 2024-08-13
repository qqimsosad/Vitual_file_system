package folder

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"virtual-file-system/pkg/models"
	"virtual-file-system/pkg/user"
)

func CreateFolder(username, foldername, description string) error {
	usr, err := user.GetUser(username)
	if err != nil {
		return err
	}

	foldername = strings.ToLower(foldername)
	if err := models.ValidateName(foldername); err != nil {
		return err
	}

	if _, exists := usr.Folders[foldername]; exists {
		return fmt.Errorf("the %s has already existed", foldername)
	}

	usr.Folders[foldername] = &models.Folder{
		Name:        foldername,
		Description: description,
		Files:       make(map[string]*models.File),
		CreatedAt:   time.Now(),
	}
	return nil
}

func DeleteFolder(username, foldername string) error {
	usr, err := user.GetUser(username)
	if err != nil {
		return err
	}
	foldername = strings.ToLower((foldername))
	if _, exists := usr.Folders[foldername]; !exists {
		return fmt.Errorf("the %s doesn't exist", foldername)
	}
	delete(usr.Folders, foldername)
	return nil
}

// ListFolders 列出指定用戶的所有資料夾，並根據排序標準進行排序
func ListFolders(username string, sortBy string, order string) ([]*models.Folder, error) {
	usr, err := user.GetUser(username)
	if err != nil {
		return nil, err
	}

	// 收集所有資料夾
	folderList := make([]*models.Folder, 0, len(usr.Folders))
	for _, folder := range usr.Folders {
		folderList = append(folderList, folder)
	}

	switch sortBy {
	case "name":
		sort.Slice(folderList, func(i, j int) bool {
			if order == "desc" {
				return folderList[i].Name > folderList[j].Name
			}
			return folderList[i].Name < folderList[j].Name
		})
	case "created":
		sort.Slice(folderList, func(i, j int) bool {
			if order == "desc" {
				return folderList[i].CreatedAt.After(folderList[j].CreatedAt)
			}
			return folderList[i].CreatedAt.Before(folderList[j].CreatedAt)
		})
	default:
		sort.Slice(folderList, func(i, j int) bool {
			return folderList[i].Name < folderList[j].Name
		})
	}

	return folderList, nil
}

func RenameFolder(username, foldername, newFoldername string) error {
	usr, err := user.GetUser(username)
	if err != nil {
		return err
	}

	foldername = strings.ToLower(foldername)
	newFoldername = strings.ToLower(newFoldername)
	if err := models.ValidateName(newFoldername); err != nil {
		return err
	}

	if _, exists := usr.Folders[foldername]; !exists {
		return fmt.Errorf("the folder %s does not exist", foldername)
	}

	if _, exists := usr.Folders[newFoldername]; exists {
		return fmt.Errorf("the folder %s already exists", newFoldername)
	}

	usr.Folders[newFoldername] = usr.Folders[foldername]
	usr.Folders[newFoldername].Name = newFoldername
	delete(usr.Folders, foldername)

	return nil
}
