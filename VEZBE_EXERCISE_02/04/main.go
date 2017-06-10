/*-------------------------------------------------------
What's the value of this expression:
(true && false) || (false && true) || !(false && false)?
--------------------------------------------------------*/

package main

import "fmt"

func main() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}

//true && false = false1
//false && true = false2
//false && false = false3
//
//false1 || false2 = false4
//
//false4 || !false3 = true

//Println will be true :)
