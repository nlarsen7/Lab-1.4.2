//Nicholas Larsen
//February 11, 2020
//Finds out your miles per gallon
package main

import "fmt"

func main() {
  var tank float64
  fmt.Println("How many gallons of gas is in your gas tank?")
  //asks how many gallons in a tank
  fmt.Scan(&tank)
  var miles float64
  fmt.Println("How many miles did you go until you ran out of gas?")
  //asks how many miles traveled
  fmt.Scan(&miles)
  //scans each variable so user input is accepted
  mpg:= (miles/tank)
  fmt.Println("your miles per gallon is",mpg)
  //makes a new varialbe for miles per gallon and writes it out for you
}