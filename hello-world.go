package main

import (
		"fmt"
		"math/rand"
		"time"
		"net/http"
)

func genEven() int{
	rand.Seed(time.Now().UnixNano())
	n:=1
	for i:=0;i<1;{
		n=rand.Int()
		if (n%2)==0 {
		i=1
		
		}
	}
	return n
}

func isEven(x int) bool {
	if(x%2)==0{ 
		return true
	
	}else{
		return false
	}
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%d", genEven())
	})

	http.ListenAndServe(":8090", nil)

}