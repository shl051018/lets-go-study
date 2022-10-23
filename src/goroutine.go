package main
import "fmt"
import "time"
func main(){
    count("shl")
    count("flyy")
}
func count(person string){
    for i:=0;i<10;i++{
        go fmt.Println(person,"like people",i)
                    time.Sleep(time.Second)

    }

}