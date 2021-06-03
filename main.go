package main

import (
	"errors"
	"fmt"
	"strings"
)

func add(int number1, int number2) int {
    
   return number1 + number2;
}

func multiply(int number1, int number2) int {
    
   return number1 * number2;
}

func divide(int number1, int number2) float {
    
    if number2 ==  0 {
       panic("you can not divide to 0") //or replace panic with fmt.Println();
    }

   return number1 / number2;
}

func substract(int number1, int number2) int {
    
   return number1 - number2;
}

func main() {
  
}
