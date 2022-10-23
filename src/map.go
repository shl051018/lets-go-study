package main
import "fmt"
func main(){
    shl := map[string]string{"name":"shl","age":"18"} //변수명 :=map[key타입]vaule타입{key:value}
    for key,value :=range shl{
        fmt.Println(key,value)
    }
}