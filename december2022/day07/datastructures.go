package day07

type cdSplittedLine = string
type fileName = string
type cdFreeLine = string
type fileSize = uint64

type task struct {
	path       []fileName
	cdFreeLine cdFreeLine
}

type AbsFile struct {
	path []fileName
	size *fileSize
}

type FileNode struct {
	name     *fileName
	children []FileNode
	size     *fileSize
}

func (fileNode *FileNode) Init() *FileNode {
	fileNode.name = nil
	fileNode.children = []FileNode{}
	fileNode.size = nil
	return fileNode
}

func (fileNode *FileNode) Add(absFile AbsFile) error {
	current := fileNode

	for _, p := range absFile.path {
		var newCurrent *FileNode
		if current.name == nil {
			*current.name = p
			// } else if *current.name != p {
			// 	return fmt.Errorf("non matching filenames %v and %v for path %v", current.name, p, absFile.path)
		} else {
			for _, child := range current.children {
				if p == *child.name {
					newCurrent = &child
					break
				}
			}
		}
		if newCurrent == nil {
			newCurrent = new(FileNode).Init()
		}
		current.children = append(current.children, *newCurrent)
		current = newCurrent
	}

	return nil
}
