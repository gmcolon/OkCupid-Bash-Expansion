# OkCupid Bash Expansion
### Description
This program implements a similar behavior to bash's brace expansion.  
For a valid input, it will expand the expression and return it as a string.  
For invalid input, it will print nothing and exit.  

Input characters are restricted to [a-zA-Z{},]  

An example invocation of the program will be:  
    - `$ echo "{A,B,C}" | ./OkCupidBashExpansion`  
And it should print:  
    - `A B C`
### Running
1. Clone project from github:
    - `https://github.com/gmcolon/OkCupid-Bash-Expansion.git`
2. Project can be built by running the following command in the project root:
    - `go build`
3. Project can be run by piping in input on the command line:
    - `$ echo "{A,B,C}" | ./OkCupidBashExpansion`

### Tests
1. Test can be run by running the following command in the project root:
    - `go test`