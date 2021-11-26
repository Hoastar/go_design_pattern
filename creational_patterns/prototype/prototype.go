package main

// 尝试通过基于操作系统文件系统的示例来理解原型模式
// package prototype
import (
	"fmt"
)

// inode 原型接口
type inode interface {
	print(string)
	clone() inode
}

// file 具体原型
type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

// folder 原型
type folder struct {
	name     string
	children []inode
}

func (fd *folder) print(indentation string) {
	fmt.Println(indentation + fd.name)
	for _, i := range fd.children {
		i.print(indentation + indentation)
	}
}

func (fd *folder) clone() inode {
	cloneFolder := &folder{name: fd.name + "_clone"}
	var tmepChildren []inode
	for _, i := range fd.children {
		copy := i.clone()
		tmepChildren = append(tmepChildren, copy)
	}
	cloneFolder.children = tmepChildren
	return cloneFolder
}

// 客户端代码
func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}

	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("  ")

	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("  ")
}
