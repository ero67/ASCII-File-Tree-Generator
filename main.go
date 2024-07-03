package main

import ("os")


func main() {
	arg1 := os.Args[1]
	treeContent,_ := constructTree(arg1,arg1)
	printTree(treeContent, "",true,true)
}


