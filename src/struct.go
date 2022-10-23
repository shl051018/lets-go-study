package main 
import "fmt"
type person struct{ //어떤 구조체를 만들건지를 정의 해줘야함
    name string
    age int
    favoritefood[]string
}
func main(){
    favoritefood :=[]string{"Pork cutlet","Raw fish"} 
    shl := person{"seungho",18,favoritefood} //변수에 구조체 내용 저장
    shl := person{name:"seungho",age:18,favoritefood:favoritefood} //심플 버전
    fmt.Println(shl)
    fmt.Println(shl.favoritefood)

}