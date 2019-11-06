package main

import "fmt"
import "math/rand"
import "time"

func rollDadu()int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}
func main(){

	var posisiOrang = 0
	var command = "Anu"

	fmt.Printf("Enter %q command: ","roll")
	fmt.Scanln(&command)
	for command == "roll" {
		var dadu = rollDadu()
		posisiOrang+=dadu
		printBoard(posisiOrang)
		fmt.Println(dadu)
		fmt.Scanln(&command)
	}
}

func printBoard(posisi int){
	var k int = 100
	for i := 0; i <= 20; i++ {
		for j := 0; j < 10; j++ {
			if i%2==0 {
				fmt.Printf("+----")	
			} else {
				if posisi == k {
					fmt.Printf("|%-4s","O")	
				} else{
					fmt.Printf("|%-4d",k)	
				}
				k--
			}
		}
		if i%2 == 0 {
			fmt.Println("+")	
		} else {
			fmt.Println("|")	
		}
	}
}