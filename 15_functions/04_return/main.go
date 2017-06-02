package main

import "fmt"

func main() {
  fmt.Println(greet("Jane", "Doe"))
}

func green(fname, lname string) string {
  return fmt.Sprint(fname, lname)
}
