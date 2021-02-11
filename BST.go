package main

import "fmt"

type node struct{

	key int
	left *node
	right *node

}

func (l *node)insert(item int){

	if l.key <item{

		if l.right == nil{

			l.right = &node{key: item}

		} else {

			l.right.insert(item)
		}

	} else if l.key > item{ 

		if l.left == nil{

			l.left = &node{key: item}

		}else{

			l.left.insert(item)
		}


	}

}


func (l *node)search(item int)bool{

	if l == nil{

		return false

	}

	if l.key < item{ 

		return l.right.search(item)

	}else if l.key > item{

		return l.left.search(item)

	}

	return true

}

func main(){

	temp:= &node{key : 500}

	//fmt.Println(temp)

	temp.insert(200)
	temp.insert(300)
	temp.insert(600)
	temp.insert(100)
	temp.insert(550)
	temp.insert(640)
	temp.insert(680)
	temp.insert(660)
	temp.insert(700)

	fmt.Println(temp)

	fmt.Println(temp.search(400))

}