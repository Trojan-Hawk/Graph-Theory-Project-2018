package main

import (
	"fmt"
)


//************************************************************************************************************************************************************
//		Shunting Yard Algorithm
func Intopost(infix string) string {
    // creating a map of special characters and assigning them a value
    specials := map[rune]int{'*': 10, '.': 9, '|': 8}

    // a rune is a character as it's displayed on the screen (utf8)
    postfix := []rune{} // initialise an array of runes
    stack := []rune{}
    
    // will loop through infix string, the first thing it will return is the character position
    // when you call range on infix, because it is a string, range converts each character to a rune
    for _, r := range infix {
        switch {
            // the current character is an open bracket
            // this is popped onto the basic rune stack
            case r == '(':
                stack = append(stack, r)

            // the current character is a close bracket
            // we need to loop through the stack until an open bracket is found 
            // while we loop throught the stack any special characters need to be appended onto the postfix array
            case r == ')':
                for stack[len(stack)-1] != '(' {
                    // append the special characters onto postfix array
                    postfix = append(postfix, stack[len(stack)-1])
                    // make s equal to everything in s, except for the last character
                    stack = stack[:len(stack)-1]
                }// for

                // make s equal to everything in s, except for the last character
                // this discards the open bracket
                stack = stack[:len(stack)-1]

            // the current character is in the specials map if anything other than 0 is returned
            case specials[r] > 0:
                // while the stack still has elements on it 
                // and the precedence of the current character that you are reading
                // is less than the precedence of whatever is at the top of the stack
                for len(stack) > 0 && specials[r] <= specials[stack[len(stack)-1]] {
                    // append the special characters onto postfix array
                    postfix = append(postfix, stack[len(stack)-1])
                    // make s equal to everything in s, except for the last character
                    stack = stack[:len(stack)-1]
                }// for 
                // append the current character onto the stack
                stack = append(stack, r)

            // take the character and append to the end of the array
            default:
                postfix = append(postfix, r)
        }// switch 
    }// for

    // add any remaining characters to the end of the postfix array
    for len(stack) > 0 {
        // append the special characters onto postfix array
        postfix = append(postfix, stack[len(stack)-1])
        // make s equal to everything in s, except for the last character
        stack = stack[:len(stack)-1]
    }// for

    // return the postfix array cast to a string
    return string(postfix)
}// Intopost
//************************************************************************************************************************************************************

func main() {
	
	fmt.Println()
} // main
