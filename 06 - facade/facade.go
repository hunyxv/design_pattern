package facede

import "fmt"

type Status int

const (
	New Status = iota
	Running
	Sleeping
	Restart
	Zombie
)

// System 系统接口
type System interface {
	Boot()
	Kill()
}

type FileSystem struct {
	Name   string
	Status Status
}

func NewFileSystem() *FileSystem {
	return &FileSystem{
		Name:   "FileSystem",
		Status: New,
	}
}

func (fs *FileSystem) Boot() {
	fmt.Printf("booting the %s...\n", fs.Name)
	fs.Status = Running
}

func (fs *FileSystem) Kill() {
	fmt.Printf("Killing %s...\n", fs.Name)
	fs.Status = Zombie
}

func (fs *FileSystem) CreateFile(user, fileName string, permissions int) {
	if fs.Status == Running {
		fmt.Printf("trying to create the file '%s' for user '%s' with permissions %o\n", user, fileName, permissions)
		return
	}
	fmt.Println("filesysetem is killed.")
}

type ProcessSystem struct {
	Name   string
	Status Status
}

func NewProcessSystem() *ProcessSystem {
	return &ProcessSystem{
		Name: "ProcessSystem",
		Status: New,
	}
}

func (ps *ProcessSystem) Boot() {
	fmt.Printf("booting the %s...\n", ps.Name)
	ps.Status = Running
}

func (ps *ProcessSystem) Kill() {
	fmt.Printf("Killing %s...\n", ps.Name)
	ps.Status = Zombie
}

func (ps *ProcessSystem) CreateProcess(user, name string) {
	if ps.Status == Running{
		fmt.Printf("trying to create the process '%s' for user '%s'\n", name, user)
		return
	}
	fmt.Printf("ProcessSystem is Killed.\n")
}

type OperatingSystem struct {
	Fs *FileSystem
	Ps *ProcessSystem
}

func NewOperationSystem() *OperatingSystem {
	return &OperatingSystem{
		Fs: NewFileSystem(),
		Ps: NewProcessSystem(),
	}
}

func (os *OperatingSystem) Start(){
	os.Fs.Boot()
	os.Ps.Boot()
}

func (os *OperatingSystem) CreateFile(user, name string, permissions int) {
	os.Fs.CreateFile(user, name, permissions)
}

func (os *OperatingSystem) CreateProcess(user, name string) {
	os.Ps.CreateProcess(user, name)
}

