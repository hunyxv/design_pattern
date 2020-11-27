package command

import (
	"fmt"
	"os"
	"path/filepath"
)

type Executer interface {
	Execute() error
	Undo() error
}

type RenameFile struct {
	SrcFile  string
	DestFile string
}

func (r *RenameFile) Execute() error {
	return os.Rename(r.SrcFile, r.DestFile)
}

func (r *RenameFile) Undo() error {
	return os.Rename(r.DestFile, r.SrcFile)
}

type CreateFile struct {
	SrcFile string
	Text    string
}

func (c *CreateFile) Execute() (err error) {
	f, err := os.Create(c.SrcFile)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.WriteString(c.Text)
	return
}

func (c *CreateFile) Undo() error {
	return os.Remove(c.SrcFile)
}

type RemoveFile struct {
	SrcFile string
}

func (r *RemoveFile) Execute() (err error) {
	_, err = os.Stat(".trash")
	if err != nil {
		if os.IsExist(err) {
			fmt.Println(os.IsExist(err))
			return
		}
		err = os.Mkdir(".trash", os.ModeDir)
		
		if err != nil {
			return
		}
	}

	return os.Rename(r.SrcFile, fmt.Sprintf(".trash/%s", filepath.Base(r.SrcFile)))
}

func (r *RemoveFile) Undo() (err error) {
	return os.Rename(fmt.Sprintf(".trash/%s", filepath.Base(r.SrcFile)), r.SrcFile)
}

func Do(cmd Executer) error {
	return cmd.Execute()
}

func Undo(cmd Executer) error {
	return cmd.Undo()
}
