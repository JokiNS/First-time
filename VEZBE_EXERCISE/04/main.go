/*--------------------------------------------------
Create a program that prints to the terminal
asking for a user to enter a small number
and a larger number. Print the remainder
of the larger number divided by the smaller number.
--------------------------------------------------*/

package main

import "fmt"

func main(){
  var Vecibroj int
  var Manjibroj int
  fmt.Print("Unesite veci broj: ")
  fmt.Scan(&Vecibroj)
  fmt.Print("Unesite manji broj: ")
  fmt.Scan(&Manjibroj)
  fmt.Println("Izracunavanje razlike ova dva broja ide ovako: ", Vecibroj ," - " , Manjibroj, " = ", Vecibroj - Manjibroj)
}
