package main
    "fmt"
)
func mutiply(a int ,b int) int { //func 함수명(인자 ,인자 인자타입) 
    defer fmt.Println("working") //defer은 함수가 값을 리턴힌 후에  실행된다
    return a*b
      
}

func main(){
    fmt.Println(mutiply(9999999,1111111111))

}