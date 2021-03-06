package main

import (
	"fmt"
)

// state struct
type state struct {
	symbol rune
	edge1  *state
	edge2  *state
} // state struct

// nfa struct
type nfa struct {
	initial *state
	accept  *state
} // nfa struct

//************************************************************************************************************************************************************
//		NFA Algorithm
func PofixToNfa(postfix string) *nfa {
	// an array of pointers to nfa's, the curley braces denote we want an empty one
	nfastack := []*nfa{}

	// ends with and starts with placeholders
	startsWith := []*nfa{}
	endsWith  := []*nfa{}

	// loop throught the regular expression one rune at a time
	for _, r := range postfix {
		switch r {
		case '.':				// concatenate
			// get the last thing off the stack and store in frag2
			frag2 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already in frag2
			nfastack = nfastack[:len(nfastack)-1]
			// get the last thing off the stack and store in frag1
			frag1 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already in frag1
			nfastack = nfastack[:len(nfastack)-1]

			// join the fragments together by setting the accept state of frag1
			// to the initial state of frag2
			frag1.accept.edge1 = frag2.initial

			// then we append the new nfa we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})
		case '|':				// one or the other
			// get the last thing off the stack and store in frag2
			frag2 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag2
			nfastack = nfastack[:len(nfastack)-1]
			// get the last thing off the stack and store in frag1
			frag1 := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			// new accept state
			accept := state{}
			// the new initial state where it's two edges point to the two fragments initial states
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			// we then set the accept states of the fragments to the new accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			// then we append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '*':				// zero or more of
			// get the last thing off the stack and store in frag1
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
			// and points to the new accept state at edge2
			initial := state{edge1: frag.initial, edge2: &accept}
			// join the fragment edge1 to the new accept state
			frag.accept.edge1 = &accept
			// join the fragment edge2 to the new initial state
			frag.accept.edge2 = &initial

			// then we append the new nfa accept state and initial state we created above to the nfastack
            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        case '?':				// zero or one of
			// get the last thing off the stack and store in frag1
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
			// and points to the new accept state at edge2
			initial := state{edge1: frag.initial, edge2: &accept}
			// join the fragment edge1 to the new accept state
			frag.accept.edge1 = &accept

			// then we append the new nfa accept state and initial state we created above to the nfastack
            nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
        case '+':				// one or more of
			// get the last thing off the stack and store in frag1
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
			// and points to the new accept state at edge2
            initial := state{edge1: frag.initial}
            // a middle state
            middle := state{edge1: frag.initial, edge2: &accept}
			// join the fragment edge1 to the middle state
			frag.accept.edge1 = &middle

			// then we append the new nfa accept state and initial state we created above to the nfastack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		case '^':				// starts with
			// get the last thing off the stack and store in frag1
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
            initial := state{edge1: frag.initial}
			// the fragment is accepted on edge1
			frag.accept.edge1 = &accept

			// then we append the new nfa accept state and initial state we created above to the startsWith stack
			startsWith = append(startsWith, &nfa{initial: &initial, accept: &accept})

			// prepending the new nfa to the start of the stack
			nfastack = append(nfastack, startsWith[0])
			// resetting the startsWith array
			startsWith = startsWith[:len(startsWith)-1]
		case '$':				// ends with
			// get the last thing off the stack and store in frag
			frag := nfastack[len(nfastack)-1]
			// get rid of the last thing on the stack, because it's already on frag1
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			// the new initial state that points to the initial of the fragment at edge1
            initial := state{edge1: frag.initial}
			// the fragment is accepted on edge1
			frag.accept.edge1 = &accept

			// then we append the new nfa accept state and initial state we created above to the endsWith stack
			endsWith = append(endsWith, &nfa{initial: &initial, accept: &accept})
		default:
			// new empty accept state
			accept := state{}
			// new initial state with the symbol value of r
			// and it's only edge points at the new accept state
			initial := state{symbol: r, edge1: &accept}

			// appending the new nfa to the stack
			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		} // switch
	} // for

	// if the endsWith and nfastack arrays are populated
	if len(endsWith) > 0 && len(nfastack) > 0 {
		// get the last thing off the stack and store in frag
		frag := nfastack[len(nfastack)-1]
		// get rid of the last thing on the stack, because it's already on frag1
		nfastack = nfastack[:len(nfastack)-1]

		accept := state{}
		// the new initial state that points to the initial of the fragment at edge1
        initial := state{edge1: frag.initial}
		// the fragment, on edge 1, then points to the first position on the endsWith array
		frag.initial.edge1 = endsWith[0].initial
		// the endsWith array then points to the accept state
		endsWith[0].accept = &accept

		// appending the new nfa to the stack
		nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})	
	}// if
	// if the endsWith array is populated
	if len(endsWith) > 0 {
		// return the endsWith array nfa
		return endsWith[0]
	}// if

	// error handling, if the nfastack has more than one nfa
	if len(nfastack) != 1 {
		fmt.Println("Error the NFA stack has been overloaded by ", len(nfastack)-1, ".")
	} // if

	// the only item will be the actual nfa that you want to return
	return nfastack[0]
} // PofixToNfa
//************************************************************************************************************************************************************
//		Shunting Yard Algorithm
func InfixToPofix(infix string) string {
    // creating a map of special characters and assigning them a value
    specials := map[rune]int{'$': 13, '+': 12, '?': 11, '*': 10, '.': 9, '^': 8, '|': 7}

    // a rune is a character as it's displayed on the screen (utf8)
    postfix := []rune{} // initialise an empty array of runes
    stack := []rune{} // initialise an empty array of runes
    
    // will loop through infix string, the first thing it will return is the character position
    // when you call range on infix, because it is a string, range converts each character to a rune
    for _, r := range infix {
        switch {
			case r == '(':
				// the current character is an open bracket
            	// this is popped onto the basic rune stack
                stack = append(stack, r)
			case r == ')':
				// the current character is a close bracket
				// we need to loop through the stack until an open bracket is found 
				// while we loop throught the stack any special characters need to be appended onto the postfix array
                for stack[len(stack)-1] != '(' {
                    // append the special characters onto postfix array
                    postfix = append(postfix, stack[len(stack)-1])
                    // make s equal to everything in s, except for the last character
                    stack = stack[:len(stack)-1]
				}// for
				
                // make the stack equal to everything in the stack, except for the last character
                // this discards the open bracket
                stack = stack[:len(stack)-1]
			case specials[r] > 0:
				// the current character is in the specials map if anything other than 0 is returned
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
			default:
				// take the character and append to the end of the array
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
}// InfixToPofix
//************************************************************************************************************************************************************
//      Regular Expression Matching Function
func RegexMatch(po string, s string) bool {
	ismatch := false

	// creating a postfix nfa, based on the postfix string
	ponfa := PofixToNfa(po)

	// an array of pointers that track the current positions in the nfa
	current := []*state{}
	// an array of pointers that track the next position based on the current position
	next := []*state{}

	// setting the current state for the first time
	current = addState(current[:], ponfa.initial, ponfa.accept)

    // loop throught each character
	for _, r := range s {
        // for each initial position in the nfa
		for _, c := range current {
            // if the symbol of one of the states matches r
			if c.symbol == r {
                // the next state is updated by calling the addState method
				next = addState(next[:], c.edge1, ponfa.accept)
			} // if
		} // inner for

		// when a character is read, set current to the values of next
		// and reset the next array to a blank state pointer array
		current, next = next, []*state{}
	} // outer for

	// check to see if the postfix nfa's accept state is in current
	for _, c := range current {
        // if it is and is an accept state
		if c == ponfa.accept {
            // set the boolean value to true and break the loop
            ismatch = true
            break
		} // if
	} // for

	return ismatch
}// RegexMatch
//************************************************************************************************************************************************************
//     Add State Helper Function
func addState(l []*state, s *state, a *state) []*state {
    // add the initial state to the array
	l = append(l, s)

    // if the initial is not equal to the accept state
    // and the initial state symbol is not set
	if s != a && s.symbol == 0 {
        // call the method recursively
		l = addState(l, s.edge1, a)
        // if the edge2 of the initial state is not a null value
        if s.edge2 != nil {
            // call the method recursively
			l = addState(l, s.edge2, a)
		} // if
	} // if

	return l
} // addState
//************************************************************************************************************************************************************

func main() {
	// 						TESTING
	// PofixToNfa, RegexMatch and addState method testing
	// '^' starts with and '$' ends with operator tests
	fmt.Println("Should return 'true', returns: ", RegexMatch("c^ab..", "cab"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("b$", "b"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("en.d.$", "end"))
	fmt.Println("Should return 'false', returns: ", RegexMatch("c^ab..", "crab"))
	fmt.Println("Should return 'false', returns: ", RegexMatch("b$", "ub"))
	fmt.Println("Should return 'false', returns: ", RegexMatch("en.d.$", "enl"))

	// '+' and '?' operator tests
	fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c?|", "c"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c?|", ""))
	fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c+|", "c"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c+|", "cc"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c?|", "abcd"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c?|", "cc"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c?|", "a"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c+|", ""))

    // '*', '|' and '.' operator tests 
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c*|", "000"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c*|", "ab0"))
    fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c*|", "ab"))
    fmt.Println("Should return 'false', returns: ", RegexMatch("ab.c*|", "abc"))
    fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c*|", ""))
	fmt.Println("Should return 'true', returns: ", RegexMatch("ab.c*|", "c"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("c*", "cc"))
	fmt.Println("Should return 'true', returns: ", RegexMatch("c*", "ccc"))

    // InfixToPofix method Testing
	// Answer: ab.c*.
    fmt.Println("Infix:   ", "a.b.c*")
    fmt.Println("Postfix(ab.c*.): ", InfixToPofix("a.b.c*"))
    // Answer: abd|.*
    fmt.Println("Infix:   ", "(a.(b|d))*")
    fmt.Println("Postfix(abd|.*): ", InfixToPofix("(a.(b|d))*"))
    // Answer: abd|.c*
    fmt.Println("Infix:   ", "(a.(b|d)).c*")
    fmt.Println("Postfix(abd|.c*): ", InfixToPofix("(a.(b|d)).c*"))
    // Answer: abb.+.c.
    fmt.Println("Infix:   ", "a.(b.b)+.c")
    fmt.Println("Postfix(abb.+.c.): ", InfixToPofix("a.(b.b)+.c"))
    // Answer: ab.c?.
    fmt.Println("Infix:   ", "a.b.c?")
    fmt.Println("Postfix(ab.c?.): ", InfixToPofix("a.b.c?"))
    // Answer: abd|.+
    fmt.Println("Infix:   ", "(a.(b|d))+")
	fmt.Println("Postfix(abd|.+): ", InfixToPofix("(a.(b|d))+"))
	
} // main