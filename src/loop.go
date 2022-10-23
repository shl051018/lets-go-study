package main

 
 import (
 	"fmt"
 )
 

 func add(numbers ...int) int { //삼도트(...)은 가자를 여러개 받아오기 위해 사용한다
 	total := 0
 	for _, number := range numbers { //for 인덱스,요소값=range 컬렉션이름
 		total += number
 	}
 	return total
 }

 func main() {
    for x:=0;x<7;x++{
        fmt.Println(x)
    }
 	result := add(1, 2, 3, 4, 5, 6)
 	fmt.Println(result)
 }