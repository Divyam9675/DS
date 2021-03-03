package main

import (
	"fmt"
)

type Graph struct{

	vertices []*Vertex

}

type Vertex struct{

	key int
	adjacent []*Vertex
}


func (g *Graph) AddEdge(from, to int){
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())

	}else{

		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}	
}


func (g *Graph) getVertex(k int) *Vertex{

	for i, v := range g.vertices{
		if v.key == k{
			return g.vertices[i]
		}
	}
	return nil
}




func (g *Graph) AddVertex(k int){
	if contains(g.vertices, k){
		err := fmt.Errorf("Vertex %v  not added because it is an existing key", k)
		fmt.Println(err.Error())
	}else{
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}



func contains(s []*Vertex, k int)bool{

	for _, v := range s{
		if k == v.key{
			return true
		}
	}
	return false

}





func (g *Graph) Print(){

	for _ , v := range g.vertices{
		fmt.Printf("\n vertex %v:",v.key)

		for _, v := range v.adjacent{
			fmt.Printf(" %v ", v.key)
		}

	}

	fmt.Println()

}






func main(){

	test := &Graph{}

	for i:=0; i<5; i++{
		test.AddVertex(i)
	}

	//test.AddVertex(0)
	//test.AddVertex(1)
	//fmt.Println(test)

	test.AddEdge(1, 2)
	//test.AddEdge(1, 2)

	test.AddEdge(1, 4)
	test.AddEdge(6, 2)
	test.AddEdge(4, 2)
	test.AddEdge(3,2)
	
	test.Print()
}