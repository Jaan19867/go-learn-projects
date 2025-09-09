package main



import (
	"fmt"
	"math" 
)

type Abser interface {
	Abs() float64 
}


func main(){

	var a Abser 
	f :=MyFloat(-math.Sqrt(2))
	v:=Vertex{3 , 4}
	a=f  // a MyFloat impleents Abser
	a=&v // a*Vertex implements Abser 
	
}