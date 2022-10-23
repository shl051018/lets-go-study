package main

import "fmt"

func money(don int) bool{
    switch don{
        case 100:
            return true
        case 1000:
            return true
            case 5000:
            return false
        case 10000:
            return false

    }
    return false //아무 조건에 만족하지 못할때
    }
    

    func main(){
        fmt.Println(money(100))
    }
