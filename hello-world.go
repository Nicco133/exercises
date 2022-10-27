package main

import (
		"fmt"
		"math/rand"
		"time"
)

func genEven(){
	rand.Seed(time.Now().UnixNano())
	n:=1
	for i:=0;i<1;{
		n=rand.Int()
		fmt.Println(n)
		if (n%2)==0 {
		i=1
		}
	}
}

func main(){
	genEven()

}