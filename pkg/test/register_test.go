package pkg_test

import (
	"testing"
	"virtual-file-system/pkg/file"
	"virtual-file-system/pkg/folder"
	"virtual-file-system/pkg/user"
)

// Test User Creation
func TestUserCreation(t *testing.T) {
	username := "testuser"
	err := user.Register(username)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	_, err = user.GetUser(username)
	if err != nil {
		t.Fatalf("User %s does not exist after creation", username)
	}
}

// Test Folder Creation
func TestFolderCreation(t *testing.T) {
	username := "testuser"
	folderName := "testfolder"
	description := "Test Folder"

	// Ensure user exists
	err := user.Register(username)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Create folder
	err = folder.CreateFolder(username, folderName, description)
	if err != nil {
		t.Fatalf("Failed to create folder: %v", err)
	}

	usr, _ := user.GetUser(username)
	if _, exists := usr.Folders[folderName]; !exists {
		t.Fatalf("Folder %s does not exist after creation", folderName)
	}
}

// Test Folder Update (Rename)
func TestFolderRename(t *testing.T) {
	user.Register("testuser")
	username := "testuser"
	folderName := "testfolder"
	folder.CreateFolder(username, folderName, "Test Folder")
	newFolderName := "renamedfolder"

	// Rename folder
	err := folder.RenameFolder(username, folderName, newFolderName)
	if err != nil {
		t.Fatalf("Failed to rename folder: %v", err)
	}

	usr, _ := user.GetUser(username)
	if _, exists := usr.Folders[newFolderName]; !exists {
		t.Fatalf("Folder %s does not exist after renaming", newFolderName)
	}
	if _, exists := usr.Folders[folderName]; exists {
		t.Fatalf("Old folder name %s still exists after renaming", folderName)
	}
}

// Test Folder Deletion
func TestFolderDeletion(t *testing.T) {
	user.Register("testuser")
	username := "testuser"
	folderName := "dele_folder"
	folder.CreateFolder(username, folderName, "Test_dele_Folder")
	// Delete folder
	err := folder.DeleteFolder(username, folderName)
	if err != nil {
		t.Fatalf("Failed to delete folder: %v", err)
	}

	usr, _ := user.GetUser(username)
	if _, exists := usr.Folders[folderName]; exists {
		t.Fatalf("Folder %s still exists after deletion", folderName)
	}
}

// Test List Folders
func TestListFolders(t *testing.T) {
	username := "testuser"
	user.Register(username)
	folderNames := []string{"folder1", "folder2", "folder3"}

	// Create folders
	for _, name := range folderNames {
		folder.CreateFolder(username, name, "test_list_ "+name)
	}

	// List folders
	folders, err := folder.ListFolders(username, "name", "asc")
	if err != nil {
		t.Fatalf("Failed to list folders: %v", err)
	}
	// 驗證結果
	if len(folders) != len(folderNames) {
		t.Fatalf("Expected %d folders, but found %d", len(folderNames), len(folders))
	}

	// 檢查每個資料夾是否存在於返回的列表中
	folderMap := make(map[string]bool)
	for _, f := range folders {
		folderMap[f.Name] = true
	}

	for _, name := range folderNames {
		if !folderMap[name] {
			t.Fatalf("Folder %s not found in the list", name)
		}
	}
}

// Test File Creation
func TestFileCreation(t *testing.T) {
	username := "testuser"
	foldername := "folder1"
	fileName := "testfile"
	description := "Test create file"
	user.Register(username)
	folder.CreateFolder(username, foldername, "folder create")

	// Create file
	err := file.CreateFile(username, foldername, fileName, description)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	usr, _ := user.GetUser(username)
	folder := usr.Folders[foldername]
	if _, exists := folder.Files[fileName]; !exists {
		t.Fatalf("File %s does not exist after creation", fileName)
	}
}

// Test File Deletion
func TestFileDeletion(t *testing.T) {
	username := "testuser"
	foldername := "folder1"
	filename := "testfile"
	user.Register(username)
	folder.CreateFolder(username, foldername, "folder create")
	file.CreateFile(username, foldername, filename, "Test delete file")

	// Delete file
	err := file.DeleteFile(username, foldername, filename)
	if err != nil {
		t.Fatalf("Failed to delete file: %v", err)
	}

	usr, _ := user.GetUser(username)
	folder := usr.Folders[foldername]
	if _, exists := folder.Files[filename]; exists {
		t.Fatalf("File %s still exists after deletion", filename)
	}
}

// Test List Files
func TestListFiles(t *testing.T) {
	username := "testuser"
	folderName := "folder1"
	fileNames := []string{"file1", "file2", "file3"}
	user.Register(username)
	folder.CreateFolder(username, folderName, "folder create")

	// Create files
	for _, name := range fileNames {
		file.CreateFile(username, folderName, name, "Description for "+name)
	}

	// List files
	files, err := file.ListFiles(username, folderName, "name", "asc")
	if err != nil {
		t.Fatalf("Failed to list files: %v", err)
	}

	if len(files) != len(fileNames) {
		t.Fatalf("Expected %d files, but found %d", len(fileNames), len(files))
	}

	for i, name := range fileNames {
		if files[i].Name != name {
			t.Errorf("Expected file %s, but got %s", name, files[i].Name)
		}
	}
}
