package main

import (
	"fmt"
	"math/rand"
)

////////////////////
// Initialization // - optional, but will run before main() if used
////////////////////
func init(){
	fmt.Println("This is the init() function.")
}

  //////////
 // Main // - neccessary and runs first to start program (unless init() is used)
//////////
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
	fmt.Println("\nTesting if statement with 'idiom' function")
	ifIdiom(x)
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
		fmt.Printf("Generated value %d is Greater than or Equal to %d", b, a)			//	This is called an 'idiom'.
	} else {
		fmt.Printf("Generated value %d is Less than %d", b, a)
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
