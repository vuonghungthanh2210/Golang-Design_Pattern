package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

type Folder struct {
	children []*Folder
	files    []*File
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, file := range f.files {
		file.print(indentation + indentation)
	}
	for _, folder := range f.children {
		folder.print(indentation + indentation)
	}
}

// Manual cloning for Folder and File
func cloneFolder(folder *Folder) *Folder {
	// Clone the folder itself
	clone := &Folder{name: folder.name + "_clone"}

	// Clone the files in the folder
	for _, file := range folder.files {
		clone.files = append(clone.files, &File{name: file.name + "_clone"})
	}

	// Recursively clone child folders
	for _, child := range folder.children {
		clone.children = append(clone.children, cloneFolder(child))
	}

	return clone
}

func main() {
	// Create some files and folders
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}

	folder1 := &Folder{
		files:    []*File{file1},
		name:     "Folder1",
		children: []*Folder{},
	}

	folder2 := &Folder{
		files:    []*File{file2, file3},
		children: []*Folder{folder1},
		name:     "Folder2",
	}

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	// Clone the folder manually
	cloneFolder2 := cloneFolder(folder2)
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder2.print("  ")
}
