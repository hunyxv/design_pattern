package composite

import "testing"

func TestComposite(t *testing.T) {
	root := NewComponent(DirNode, "/")
	folder1 := NewComponent(DirNode, "folder1")
	folder2 := NewComponent(DirNode, "folder2")
	folder3 := NewComponent(DirNode, "folder3")

	file1 := NewComponent(FileNode, "file1")
	file2 := NewComponent(FileNode, "file2")
	file3 := NewComponent(FileNode, "file3")

	root.AddChild(folder1)
	root.AddChild(folder2)
	folder2.AddChild(folder3)

	folder1.AddChild(file1)
	folder2.AddChild(file2)
	folder3.AddChild(file3)

	root.Print("-")
}