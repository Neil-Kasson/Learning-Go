package main

import (
	"fmt"
	"math/rand"
	"time"
)

////////////////////
// Initialization // - optional, but will run before main() if used
////////////////////
func init(){
	fmt.Println("This is the init() function.")
}

  //////////
 // Main // - neccessary and runs first to start program (unless init() is used)
//////////			-testing IF & SWITCH statements
func main(){
	fmt.Println("\nThis is the start of main().")
	w := 42
	x:=42
	y:=910
	z:=783
	fmt.Println("\nTesting if statements:")
	ifStatements(w, x)
	ifStatements(x, y)
	ifStatements(y, z)
	fmt.Println("\nTesting if statement with 'idiom' function:")
	ifIdiom(x)
	fmt.Println("\nSwitch statements with comparrison operators:")
	switchA(x)
	switchA(y)
	switchA(z)
	fmt.Println("\nDirect case switches:")
	switchB(x)
	switchB(y)
	fmt.Println("\nSwitches with fallthrough:")
	switchF(x)
	switchF(y)

	   /////////////////////
	  // Move to main2() //
	 //    ~line 140		//
	/////////////////////
	main2()
}


  ///////////////////
 // If Statements //
///////////////////
func ifStatements(a, b int){
	if(a>b){
		fmt.Printf("%d is greater than %d.\n", a, b)
	} else if (a<b){		// <-<------------------------- 'else if'... not 'elif'
		fmt.Printf("%d is less than %d.\n", a, b)
	} else {
		fmt.Printf("%d and %d are equal.\n", a, b)
	}

	if ((a>b) || (b==a)) && (a<b){
		fmt.Println("This will never print, but I just wanted to show that '||' and '&&' are the same as Java")
	}
}

  ////////////
 // Idioms //
////////////
func ifIdiom(a int){
	if b:=2*rand.Intn(a); b>=a{			//  <----<--------	You can set a value in the if statement to lessen scope
		fmt.Printf("Generated value %d is Greater than or Equal to %d.\n", b, a)			//	This is called an 'idiom'.
	} else {
		fmt.Printf("Generated value %d is Less than %d.\n", b, a)
	}
}
   ////////////////////////////////////////////////////////////////
  // This is called the "ok idiom", used as sort of a try/catch //
 ////////////////////////////////////////////////////////////////
// func commaIdiom(tz String) int{
// 	if seconds, ok:=timeZone[tz]; ok{
// 		return seconds
// 	} else{
// 		log.Println("unknown time zone", tz)
// 		return 0
// 	}
// }


  ///////////////////////////
 // Switch - conditionals //
///////////////////////////
func switchA(x int){
	switch{
	case x<783:
		fmt.Printf("%d is less than 783.\n", x)
	case x>783:
		fmt.Printf("%d is greater than 783.\n", x)
	case x==783:
		fmt.Printf("%d equals 783.\n", x)
	default:
		fmt.Println("Switch test defaulted.")
	}
}

  ///////////////////////////
 // Switch - direct cases //
///////////////////////////
func switchB(x int){
	switch x{
	case 42:
		fmt.Println("x is 42")
	default:
		fmt.Print("Default: x is not 42... ")
		fmt.Printf("x is %d.\n", x)
	}
}

  /////////////////////////////
 // Switch with Fallthrough //
/////////////////////////////
func switchF(x int){
	switch x{
	case 42:
		fmt.Println("x is 42")
		fallthrough				//  <---<------- fallthrough means next case also executes.
	case 37:
		fmt.Println("x is still 42, not 37, but this is print anyway because of 'fallthrough'.")
	case 837421:
		fmt.Println("This won't print from fallthrough like the last case did.")
	default:
		fmt.Print("Default: x is not 42... ")
		fmt.Printf("x is %d.\n", x)
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////////////////////////

func main2(){
	fmt.Println("\n'Select' Statements and basic concurrency:")
	selects()
}

func selects(){
	// conncurrency
	// select a channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to type 'time.Duration', then use it in 'time.Sleep()'.
	// Int63n returns an int64
	// type 'duration' has underlying type is int64
	// time.Sleep() takes type 'duration'
	// time.Millisecond is a constant in the 'time' package
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	// go routine - concurrency
	go func(){
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	
	go func(){
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// A 'select' statement chooses which, from a set of possible send/recieve operations, will proceed.
	// It looks similar to a 'switch' statement, but with all cases referring to communication operations
	//		For more info: https://go.dev/ref/spec#Select_statements

	// My notes on what's happening:
	// Each channel 'sleeps' for a random mount of time.
	// Which ever wakes up first, when that happens, the valus is sent to it's 'v', and that case is selected
	select{
	case v1 := <-ch1:
		fmt.Println("Value from CH.1:", v1)
	case v2 := <-ch2:
		fmt.Println("Value from CH.2:", v2)
	}
}