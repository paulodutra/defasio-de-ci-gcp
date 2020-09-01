package main
import "fmt"


func sum(a int, b int) int{
	result := a + b
	return result;
}

func main(){
	a:=5;
	b:=5;
	fmt.Printf("%d \n", sum(a,b)) 
}