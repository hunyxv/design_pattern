package command

import (
	"os"
	"testing"
)

var (
	createFile = &CreateFile{SrcFile: "file1.txt", Text: "test"}
	renameFile = &RenameFile{SrcFile: "file1.txt", DestFile: "file2.txt"}
	removeFile = &RemoveFile{SrcFile: "file2.txt"}
)

func TestExecute(t *testing.T) {
	err := Do(createFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file1.txt"); err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}

	err = Do(renameFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file2.txt"); err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}

	err = Do(removeFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file2.txt"); err != nil && os.IsExist(err) {
		t.Fatal(err)
	}
	if _, err := os.Stat(".trash/file2.txt"); err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}
}

func TestUndo(t *testing.T) {
	err := Undo(removeFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file2.txt"); err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}
	if _, err := os.Stat(".trash/file2.txt"); err != nil && os.IsExist(err) {
		t.Fatal(err)
	}

	err = Undo(renameFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file1.txt"); err != nil && !os.IsExist(err) {
		t.Fatal(err)
	}

	err = Undo(createFile)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat("file1.txt"); err == nil {
		t.Fatal(err)
	}
}
