package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
    Name     string
    Children []*Node
    IsDir    bool
}


func constructTree(path,currentFile string) (*Node,error) {
	root := &Node{
		Name: currentFile,
		IsDir: true,
	}

	contents, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error reading directory: ", err)
		return nil ,nil	
	}

	for _, file := range contents {
		fullPath := filepath.Join(path, file.Name()) 

		if file.IsDir() {
			
			child, err := constructTree(fullPath, file.Name())
			if err != nil {
				return nil, err 
			}
			root.Children = append(root.Children, child)
		} else {
			
			root.Children = append(root.Children, &Node{
				Name:  file.Name(),
				IsDir: false,
			})
		}
	}

	return root, nil
}