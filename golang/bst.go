package main

import (
	_"strconv"
	"github.com/kr/pretty"
)

type BinarySearchTree struct {
	Value int
	Left *BinarySearchTree
	Right *BinarySearchTree
}

func (bst *BinarySearchTree) Insert(value int) {

	//Create new node
	node := BinarySearchTree{}
	node.Value = value

	var recursive func(bstNode *BinarySearchTree)
	recursive = func(bstNode *BinarySearchTree) {

        if(bstNode.Value > node.Value && bstNode.Left == nil){
			bstNode.Left = &node
		}else if(bstNode.Value > node.Value){
			recursive(bstNode.Left)
		}else if(bstNode.Value < node.Value && bstNode.Right == nil){
			bstNode.Right = &node
		}else if(bstNode.Value < node.Value){
			recursive(bstNode.Right)
		}	
    }

    recursive(bst)	
}

func (bst *BinarySearchTree) Contains(value int) bool {

	doesContain := false;

	var recursive func(bstNode *BinarySearchTree)
	recursive = func(bstNode *BinarySearchTree) {

        if(bstNode.Value == value){ 
			doesContain = true
		}else if(bstNode.Left == nil && value < bstNode.Value){
			recursive(bstNode.Left)
		}else if(bstNode.Right == nil && bstNode.Value > value){
			recursive(bstNode.Right)
		}	
    }

    recursive(bst)
    return doesContain	
}

func (bst *BinarySearchTree) DepthFirstLog(f func(int)) {
	var recursive func(bstNode *BinarySearchTree)

	recursive = func(bstNode *BinarySearchTree) {
     
		f(bstNode.Value)

		if(bstNode.Left != nil){
			recursive(bstNode.Left)
		}

		if(bstNode.Right != nil){
			recursive(bstNode.Right)
		}	
    }

    recursive(bst)    
}

func main(){
	var bst BinarySearchTree
	bst.Value = 70
	bst.Insert(10)	
	bst.Insert(15)	

	bst.DepthFirstLog(func(val int) {
		pretty.Println(val)
	})	
}