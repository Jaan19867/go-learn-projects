package tourofgo


// go me class nhi he 
// but can define methods on type 
import (
	"fmt"
	"math"
)

type Vertex struct{
	X , Y  float64 
}

type MyFloat float64 

func ( f MyFloat) Nfloat() float64 {
	if(f<0){
		return float64(-f)
	}
	return float64(f)
}

// so we can also use mehtod on a numeric type not only on composite datatype 



func (v Vertex) Abs() float64 {
	// here v is receiver 
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// here the function belong to a type 


func main(){
	v :=Vertex{3 , 4} ; 
	fmt.Println(v.Abs()) ; 
	// fucking crazy ham yeha pe v ke through pass kar rahe he . 


}