package main

import "fmt"

func main() {
  // 채널 생성
  gochannel := make(chan string) //채널명 :=make(chan 타입)
  
  go func() { 
    // 값 전달
    /*채널명 <-데이터(송신)
    데이터:=<채널명(수신)
    <-채널명(이런 경우도가능
    */
    gochannel <- "Welcome to gochannel"
    
    
    // 채널 닫음
    close(gochannel)
  } ()
  
  // for~range를 사용해서 채널이 닫힐때 까지 수신가능
  for message := range gochannel {
    fmt.Println(message)
  }
}