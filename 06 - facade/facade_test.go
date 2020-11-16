package facede

import "testing"

func TestFacede(t *testing.T) {
	os := NewOperationSystem()
	os.Start()
	os.CreateFile("foo", "file.txt", 0755)
	os.CreateProcess("foo", "ls /tmp")
}