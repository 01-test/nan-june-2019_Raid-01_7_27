package Raid01

import (
	"fmt"
)

func printline( n,x,y int) {
	if(x == 1 || x==y){
		fmt.Print("A")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("B")
		}
		if n >=2 {
			fmt.Print("C")
		}
		fmt.Print("\n")
	} else {
		fmt.Print("B")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print(" ")
		}
		if n >=2 {
			fmt.Print("B")
		}
		fmt.Print("\n")
	}
}

func Raid1d(x,y int)  {
	for i:=1;i<=y;i++{
		printline(x,i,y)
	}
}
