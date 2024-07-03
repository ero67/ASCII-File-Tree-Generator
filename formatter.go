package main

import ("fmt")

func printTree(node *Node, indent string, isLast, isroot bool) {
	prefix := "|--"
	if isLast {
			prefix = "└──"
	}
	if isroot {
		prefix = ""
	}
	fmt.Printf("%s%s\n", indent, prefix+node.Name)

	// Handle children
	lastChild := len(node.Children) - 1
	for i, child := range node.Children {
		childIndent := indent
		if !isLast { 
			childIndent += "|   "
		} else {
			childIndent += "    "
		}

		printTree(child, childIndent, i == lastChild,false)
	}
}
