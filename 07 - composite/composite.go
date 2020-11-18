package composite

import "fmt"

// NodeType 节点类型
type NodeType int

const (
	// DirNode 文件夹，容器构件
	DirNode NodeType = iota
	// FileNode 文件，叶子构件
	FileNode
)

// Component 抽象构件
type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

func NewComponent(kind NodeType, name string) Component {
	var c Component
	switch kind {
	case DirNode:
		c = NewDir()
	case FileNode:
		c = NewFile()
	default:
		panic("error")
	}

	c.SetName(name)
	return c
}

var _ Component = (*component)(nil)

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

// DirComposite 目录节点
type DirComposite struct {
	component
	childern []Component
}

func NewDir() Component {
	return &DirComposite{
		childern: make([]Component, 0),
	}
}

func (dir *DirComposite) AddChild(child Component) {
	child.SetParent(dir)
	dir.childern = append(dir.childern, child)
}

func (dir *DirComposite) Print(prefix string) {
	fmt.Printf("|%s %s\n", prefix, dir.Name())
	prefix += "-"
	for _, dir := range dir.childern {
		dir.Print(prefix)
	}
}

// FileLeaf 文件节点
type FileLeaf struct {
	component
}

func NewFile() Component {
	return &FileLeaf{}
}

func (leaf *FileLeaf) Print(prefix string) {
	fmt.Printf("|%s └%s\n", prefix, leaf.Name())
}
