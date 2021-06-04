// go run main.go

package main

import (
	//"errors"
	"fmt"
	//"strings"
)


func add(number1 int, number2 int) int {

   return number1 + number2;
}


func multiply(number1 int, number2 int) int {

   return number1 * number2;
}


func divide(number1 int, number2 int) float32 {

    if number2 ==  0 {
       panic("you can not divide to 0") //or replace panic with fmt.Println();
    }

   return float32(number1 / number2);
}


func substract(number1 int, number2 int) int {

   return number1 - number2;
}

func main() {
	fmt.Println("add 2+2=", add(2,2) )
}
