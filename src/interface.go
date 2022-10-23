package main

import "fmt"

type MusicPlayer interface {
	power()
	Sound()
}

type O struct {
}

func (on O) power() {
	fmt.Println("전원이 켜져있습니다..")
}

func (on O) Sound() {
	fmt.Println("소리가 켜져있습니다..")
}

type  X struct {
}

func (off  X) power() {
	fmt.Println("전원이 꺼져있습니다..")
}

func (off  X) Sound() {
	fmt.Println("소리가 껴져있습니다..")
}


func condition(M ...MusicPlayer) {
	for _, u := range M{
		u.power()
		u.Sound()
	}
}

func main() {
	yes := O{}
	not :=  X{}
	condition(yes,not)
}
