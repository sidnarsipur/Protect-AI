package main

import (
	"fmt"
	"strings"

	"github.com/yterajima/go-sitemap"
)

type Directory struct {
	Folders []Folder
}

type Folder struct {
	FolderName string
	Folders    []Folder
	Files      []string
}

var root string

func GetSitemap(url string) {
	root = url

	sitemap, err := sitemap.Get(url, nil)
	if err != nil {
		panic(err)
	}

	BuildDirectory(sitemap)
}

func BuildDirectory(sitemap sitemap.Sitemap) {
	rootFolder := Folder{
		FolderName: root,
	}

	for _, url := range sitemap.URL {
		components := strings.Split(url.Loc, "/")[3:] // Extract relevant components after "https:", "", and "openai.com"
		currentFolder := &rootFolder

		for _, component := range components[:len(components)-1] {
			subFolder := findFolderByName(currentFolder.Folders, component)

			if subFolder == nil {
				subFolder = &Folder{
					FolderName: component,
				}
				currentFolder.Folders = append(currentFolder.Folders, *subFolder)
			}

			currentFolder = subFolder
		}

		file := components[len(components)-1]
		currentFolder.Files = append(currentFolder.Files, file)
	}

	directory := Directory{
		Folders: []Folder{rootFolder},
	}

	for _, folder := range directory.Folders {
		printFolder(folder)
		fmt.Println(" ")
	}
}

func findFolderByName(folders []Folder, name string) *Folder {
	for i := range folders {
		if folders[i].FolderName == name {
			return &folders[i]
		}
	}
	return nil
}

func printFolder(folder Folder) {
	fmt.Println(folder.FolderName)
	for _, file := range folder.Files {
		fmt.Println(file)
	}
	for _, subFolder := range folder.Folders {
		printFolder(subFolder)
	}
}
