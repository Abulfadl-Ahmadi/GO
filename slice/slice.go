package main
import "fmt"

func main(){

	
	my_slice :=[]int{9,0,0,89,87}
	my_slice[0] = 1000
	fmt.Println(my_slice)
	
	my_slice = append(my_slice, 0, 2)
	
	fmt.Println(my_slice)
 
}

