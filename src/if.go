package main

import "fmt"

func money(don int)bool{
    if don>10000{
        return true
    }
    return false
    }
    

    func main(){
        fmt.Println(money(5000))
    }
