package memento

import (
	"testing"
	"time"
)

func TestMemento(t *testing.T) {
	storageCaretaker := NewStorageCaretaker()

	edit := NewEdit("X123", "abcdefg", storageCaretaker)
	t.Log(edit)
	time.Sleep(time.Second)

	edit.Update("---------")
	t.Log(edit)

	edit.Update("+++++++++++")
	t.Log(edit)
}
