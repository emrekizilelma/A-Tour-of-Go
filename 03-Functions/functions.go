/*
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.
*/

package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func add_cool(x, y int) int {
  return x + y  
}

/* When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
In this example, we shortened x int, y int to x, y int
*/
func main() {
  fmt.Println("add function(x int, y int)")
  fmt.Println(add(42, 13))
  fmt.Println(add(123, 456))
  fmt.Println("add_cool function(x, y int)")
  fmt.Println(add_cool(42, 13))
  fmt.Println(add_cool(123, 456))
}
