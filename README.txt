				
				Graph Theory Project 2018
				
	Author: Timothy Cassidy		Student No: G00333333
	
				How to run the program
			
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
	sufficient.
	
	
	The project is split up in the following way:
	
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
	'*' character
	
	The fourth case is entered when the rune is equal to the 
	'?' character
	
	The fifth case is entered when the rune is equal to the 
	'+' character
	
	The sixth case is entered when the rune is equal to the 
	'^' character
	
	The last case is entered when the rune is equal to the 
	'$' character
	
	InfixToPofix
	RegexMatch
	addState
	main
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	