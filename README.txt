--------------------------------------------------------------------
				Graph Theory Project 2018
--------------------------------------------------------------------		
	Author: Timothy Cassidy		Student No: G00333333
--------------------------------------------------------------------
				How to run the program
--------------------------------------------------------------------

	To run the program you need to clone or download the 
	repository to your local machine. The link below will 
	show you how to do this:
	https://help.github.com/articles/cloning-a-repository/
	
	The program can be run using the windows terminal.
	You can access this by pressing the windows key,
	typing 'cmd' and then hitting the enter key.
	
	This opens the terminal, next navigate into the directory
	where you cloned/downloaded the repository. To run the 
	next command you need to have already downloaded a
	version of the go language, if not it can be found on
	the folowing link:	https://golang.org/dl/
	
	In the terminal type 'go run main.go', this will run the 
	program on the terminal window. The program already has
	multiple test cases for each operator. More can be added
	if needed but I thought the test cases provided are 
	sufficient. The project is split up in the following way:

--------------------------------------------------------------------
				PofixToNfa Function
--------------------------------------------------------------------
	
	The 'PofixToNfa' function parses a postfix string and 
	creates an Non-Deterministic Finite Automata based on
	it. It does so by first using two structs which are 
	declared above it. 
	
	The fist struct is a state struct, this keeps track of 
	the rune symbol, an edge1 state pointer and an edge2 
	state pointer. The second struct is an nfa struct, this
	keeps track of the initial and accept state pointers.
	
	The 'PofixToNfa' function starts off by first declaring
	three empty nfa pointer arrays, the first named nfastack
	will be heavily used through this function, the others 
	will be used later. Next is a for loop that parses each 
	rune character of the postfix string until there are 
	none left to be parsed. 
	
	Inside this for loop is a switch statement based on each 
	individual rune character. The switch statement has 
	seven different cases followed by a default statement.
	
	The first case is entered when the rune is equal to the 
	'.' character. This case will pull two fragments from the
	nfastack, named frag1 and frag2, then the last two items
	on the stack are removed because they are stored in the 
	frag variables. Next the accept state on frag1 at edge1
	is set to frag2's initial state and then an nfa with 
	frag1's initial state and frag2's accept state is appended
	to the nfastack.
	
	The second case is entered when the rune is equal to the 
	'|' character. This case will pull two fragments from the
	nfastack, named frag1 and frag2, then the last two items
	on the stack are removed because they are stored in the 
	frag variables. Two new states are made an initial
	and an accept state, this new initial state points to 
	frag1's initial state at edge1 and points to frag2's 
	initial state ate edge2. Next we make edge1 on both 
	fragments point to the new accept state and then an nfa 
	with the new initial state and the new accept state is 
	appended to the nfastack.
	
	The third case is entered when the rune is equal to the 
	'*' character. This case will pull a fragment from the
	nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. Two
	new states are made an initial and an accept state, this 
	new initial state points to frag1's initial state at edge1 
	and points to the new accept state at edge2. Next we make 
	edge1 on frag point to the new accept state and edge2 point 
	to the new initial state, then an nfa with the new initial 
	state and the new accept state is appended to the nfastack.
	
	The fourth case is entered when the rune is equal to the 
	'?' character. This case will pull a fragment from the
	nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. Two
	new states are made an initial and an accept state, this 
	new initial state points to the fragment's initial state at 
	edge1 and points to the new accept state at edge2. Next we 
	make edge1 on frag point to the new accept state then an nfa 
	with the new initial state and the new accept state is 
	appended to the nfastack.
	
	The fifth case is entered when the rune is equal to the 
	'+' character. This case will pull a fragment from the
	nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. Three
	new states are made an initial, an accept and a middle state, 
	this new initial state points to the fragment's initial state 
	at edge1 and the middle state points to the fragment's initial
	state at edge1 and points to the new accept state at edge2. 
	Next we make edge1 on frag point to the new middle state then 
	an nfa with the new initial state and the new accept state is 
	appended to the nfastack.
	
	The sixth case is entered when the rune is equal to the 
	'^' character. This case will pull a fragment from the
	nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. Two
	new states are made an initial and an accept state, this 
	new initial state points to the fragment's initial state at 
	edge1 then we make edge1 on frag point to the new accept 
	state. Next I appended the new initial state and the new 
	accept state to the startsWith nfa pointer array, declared 
	above, then the startsWith array, at position 0, is appended 
	to the nfastack array. Finally, the startsWith array is 
	reset.
	
	The last case is entered when the rune is equal to the 
	'$' character. This case will pull a fragment from the
	nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. Two
	new states are made an initial and an accept state, this 
	new initial state points to the fragment's initial state at 
	edge1 then we make edge1 on frag point to the new accept 
	state. Next I appended the new initial state and the new 
	accept state to the endsWith nfa pointer array, declared 
	above, to be used later.
	
	The default case in the switch statement is used to make
	fragments based on non-special characters. This case 
	creates two new states, an initial and accept state. The 
	initial state's symbol is populated with the current rune
	and it's edge1 points to the new accept state. Then an 
	nfa with the new initial state and the new accept state is 
	appended to the nfastack.
	
	After the for loop are three if statements, the first if 
	checks if the endsWith array is populated and if the 
	nfastack is populated. If it is, a fragment is pulled from
	the nfastack, named frag, then the last item on the stack 
	is removed because it is stored in the frag variable. 
	Two new states are made an initial and an accept state, this 
	new initial state points to the fragment's initial state at 
	edge1 then we make edge1 on frag point to the endsWith array's 
	initial state at position 0. Next the endsWith array's accept 
	state, at position 0, is pointed to then new accept state.
	Finally an nfa with the new initial state and the new accept 
	state is appended to the nfastack.
	
	The second if checks to see if the endsWith array is
	populated, if so, it is returned because the nfastack is
	empty.
	
	The last if is just for error handling, if the nfastack
	is not populated an error message is displayed to the 
	user.
	
	At the end of this function the nfastack is returned at
	position 0.
	
--------------------------------------------------------------------	
				InfixToPofix Function
--------------------------------------------------------------------

	This method I first declared a map of special characters and 
	assigned them a value to denote precendence. Then initialised
	two empty arrays of runes, called postfix and stack. Next is
	a for loop which will parse through the infix string, inside 
	this loop is a switch statement. The first case of the switch 
	checks to see if the current rune is equal to the '(' 
	character, if so, it is appended to the stack rune array.
	
	The second case of the switch checks to see if the current 
	rune is equal to the ')' character, if so, the stack array
	is looped over until a '(' character is found, until it is 
	found the last item on the stack is added to the postfix 
	array and the stack is shortened by one.
	
	The next case of the switch checks to see if the current 
	rune is in the specials map, if so, the stack is looped
	over while the stack is populated and while the current
	rune special characters precendence is lower then the 
	special character on top of the stack. Inside this loop
	the last item on the stack is added to the postfix 
	array and the stack is shortened by one. Finally, the 
	current rune is appended to the stack.
	
	The default case for this switch appends the current rune
	to the postfix array.
	
	After this is a second for loop that parses the stack until
	the length of the stack is 0. Inside this loop the last item 
	on the stack is added to the postfix array and the stack is 
	shortened by one.

	Finally, the postfix rune array is case to a string and it is
	then returned.
	
--------------------------------------------------------------------
				RegexMatch Function
--------------------------------------------------------------------	

	The 

--------------------------------------------------------------------
				addState Function
--------------------------------------------------------------------



--------------------------------------------------------------------
				main Method
--------------------------------------------------------------------



--------------------------------------------------------------------
						**END OF README**
--------------------------------------------------------------------