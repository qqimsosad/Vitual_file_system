package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"virtual-file-system/pkg/file"
	"virtual-file-system/pkg/folder"
	"virtual-file-system/pkg/user"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		switch command {
		case "register":
			HandleRegister(args)
		case "create-folder":
			HandleCreateFolder(args)
		case "delete-folder":
			HandleDeleteFolder(args)
		case "rename-folder":
			HandleRenameFolder(args)
		case "list-folders":
			HandleListFolders(args)
		case "create-file":
			HandleCreateFile(args)
		case "delete-file":
			HandleDeleteFile(args)
		case "list-files":
			HandleListFiles(args)
		case "exit":
			return
		default:
			fmt.Fprintln(os.Stderr, "Unrecognized command")
		}
	}
}

func HandleRegister(args []string) {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: register [username]")
		return
	}
	err := user.Register(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Add", args[0], "successfully.")
	}
}

func HandleCreateFolder(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: create-folder [username] [foldername] [description]?")
		return
	}
	description := ""
	if len(args) > 2 {
		description = strings.Join(args[2:], " ")
	}
	err := folder.CreateFolder(args[0], args[1], description)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Create", args[1], "successfully.")
	}
}

func HandleDeleteFolder(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: delete-folder [username] [foldername]?")
		return
	}
	err := folder.DeleteFolder(args[0], args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Delete", args[1], "successfully.")
	}
}

func HandleRenameFolder(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: rename-folder [username] [foldername] [new-foldername]")
		return
	}
	err := folder.RenameFolder(args[0], args[1], args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Rename", args[1], "to", args[2], "successfully.")
	}
}

func HandleListFolders(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
		return
	}
	username := args[0]
	sortBy := "name"
	order := "asc"
	if len(args) > 1 {
		for _, arg := range args[1:] {
			switch arg {
			case "--sort-name":
				sortBy = "name"
			case "--sort-created":
				sortBy = "created"
			case "asc":
				order = "asc"
			case "desc":
				order = "desc"
			default:
				fmt.Fprintln(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]")
				return
			}
		}
	}

	folders, err := folder.ListFolders(username, sortBy, order)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		return
	}

	if len(folders) == 0 {
		fmt.Fprintln(os.Stderr, "Warning:", "The", username, "doesn't have any folders.")
	} else {
		for _, folder := range folders {
			fmt.Printf("%s\t%s\t%s\t%s\n", folder.Name, folder.Description, folder.CreatedAt.Format(time.RFC1123), username)
		}
	}
}

func HandleCreateFile(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: create-file [username] [foldername] [filename] [description]?")
		return
	}
	description := ""
	if len(args) > 3 {
		description = strings.Join(args[3:], " ")
	}
	err := file.CreateFile(args[0], args[1], args[2], description)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Create", args[2], "in", args[0]+"/"+args[1], "successfully.")
	}
}

func HandleDeleteFile(args []string) {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: delete-file [username] [foldername] [filename]")
		return
	}
	err := file.DeleteFile(args[0], args[1], args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		fmt.Println("Delete", args[2], "in", args[0]+"/"+args[1], "successfully.")
	}
}

func HandleListFiles(args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
		return
	}

	username := args[0]
	foldername := args[1]
	sortBy := "name"
	order := "asc"

	if len(args) > 2 {
		for i := 2; i < len(args); i++ {
			switch args[i] {
			case "--sort-name":
				sortBy = "name"
			case "--sort-created":
				sortBy = "created"
			case "asc":
				order = "asc"
			case "desc":
				order = "desc"
			default:
				fmt.Fprintln(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
				return
			}
		}
	}

	files, err := file.ListFiles(username, foldername, sortBy, order)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(files) == 0 {
		fmt.Println("Warning: The folder is empty.")
	} else {
		for _, file := range files {
			fmt.Printf("%s\t%s\t%s\t%s\n", file.Name, file.Description, file.CreatedAt.Format(time.RFC1123), foldername)
		}
		fmt.Println(username)
	}
}
