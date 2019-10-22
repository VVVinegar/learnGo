package main

import (
	"fmt"
)

type Father interface {
	Hello()
}


type Child struct {
	Name string
}

func (s Child)Hello()  {

}

func main(){
	var buf  Father
	buf = Child{}
	f(&buf)
}
func f(out *Father){
	if out != nil{
		fmt.Println("surprise!")
	}
}
