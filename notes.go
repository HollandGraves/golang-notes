package main

/*



GO NOTE:
:= sign

	// the two statements:
	//		var card string = "Ace of Spades"
	//		card := "Ace of Spades"
	// are totally equal to each other. The first case is a very specific form of declaration. Although, you do not need this level of specificity as the
	// go compiler will be told to intuit the data type being stored when using the := notation



GO NOTE:
= sign

	// after a variable has been declared with := then to reassign a new value to the same variable you use simply the = sign



GO NOTE:
intializing variables

	// you can initialize a variable outside of the func main() {} function, although you cannot assign a value to it
	// you can also initialize a variable, e.g. var card int
	//		without immediately assigning a variable to it; or later assigning a variable to it.



GO NOTE:
creating new functions outside of func main

	// whenever you create a function outside of main to be called inside of main, you use normal syntax as with most languages when writing functions
	// one thing, is when you're using a function that e.g. only returns a value e.g. a string you will have to put the data type of the item being returned
	// e.g. func newCard() string {return "Ace of Spades"}



GO NOTE:
Arrays in go

	// there are two basic data structures for handeling records in Go: Arrays and Slices
	// Arrays are of fixed length
	// Slices are arrays that can grow or shrink at will
	// Slices and Arrays must both be defined with ONE data type
	// 		repeating: a slice and/or an array may only have one data type in it
	// e.g. newSlice := []string{"one", "two", "three"}



GO NOTE:
Adding new elements to slices

	// to add a new element to a slice you would take the variable the slice is stored in and set it equal to the append() function,
	// 		while using at least two arguments:
	// 		one argument will be the variable that contains the slice (so as to re-store the original
	// 		slice's values into the new iteration of the cards variable; very similar to how to store values into an array using a spread
	// 		operator)
	// 		the other arugment will be the new values
	// e.g.
	// newSlice := []string{"one", "two", "three"}
	// newSlice = append(newSlice, "four", "five")
	//
	// something especially important to note here is that the newSlice variable that is equal to the append statement will be returning
	// 		a whole new version of newSlice. Go essentially goes back to the space where newSlice was originally written, deletes that entry
	// 		and writes and new iteration of newSlice, elsewhere in the RAM, database, etc



GO NOTE:
Iterating over a slice

	// to iterate over a slice you will use the following syntax
	// e.g.
	// newSlice := []string{"one", "two", "three"}
	// for i, variablePlaceholder := range newSlice {
	// 		fmt.Println(i, variablePlaceholder)
	// }
	// the important thing to notice here is that at each iteration, just as mentioned in the last section in the above note, the data for i and
	// 		variablePlaceholder are comepletely being deleted and recreated with new values
	// 		such that, at the first iteration of i:
	// 			i := 0
	// 			variablePlaceholder := "one"
	// 		then at the second iteration of i:
	//			i := 1
	// 			variablePlaceholder := "two"
	// at each step above i is being deleted from memory and recreated, same as variablePlaceholder
	//
	// also important, breaking down each step of the syntax: for i, placeHolderVariable := range newSlice {}
	// 		here the most important part is the "range" keyword, the range keyword is one of many keywords that signifies the for loop iterating over
	// 		every item in the slice.
	// also notice how a := symbol is being used, this supports that Go will delete each entry each iteration, search for the value type of the new
	// 		iteration, and resave the new iteration under the same placeHolderVariable name
	//
	// finally, to mention, in the syntax: for i, card := range cards {
	//											fmt.Println(card)
	// 										}
	// 		WONT compile, because the i variable is not being used.
	//		something about Go is that if a variable is being declared IT MUST BE USED, so you must actually use the syntax as follows:
	//			for i, card := range cards {
	// 				fmt.Println(i, card)
	// 			}



GO NOTE:
The Go version of a Class and Extending A Class

	// first the Go syntax for defining a "class" and "extending" the class to another "class"
	//
	// e.g.
	// type deck []string
	//
	// in this example we're saying that the "class" (i.e. type) named deck is "extending" the "class" []string (i.e. a slice that only contains strings)
	// in other words we're adding deck as a new data type and giving it the properties that a (e.g. []string) slice that only returns strings has



GO NOTE:
Creating a receiver function

	// a receiver function is a function of a type (e.g. type deck []string) that can take an argument and is tied access as a method of anything that
	// 		uses the deck type of slice
	// this is the syntax for a receiver function
	//
	// e.g.
	// func (d deck) print() {
	//
	// }
	//
	// in this case you will define the data type that print() is
	// 		it is a function, so we will use "func"
	// second you will use the receiver notation
	// 		(d deck)
	// 		this will allow the function to be used through dot notation on any variable that contains the use of the data type slice (which remember is
	// 		a slice data type that contains strings at heart with a new name due to the type deck []string syntax)
	// third we will write the name of the function we want to be able to use through dot notation



GO NOTE:
Using a receiver function

	// in the case of using a receiver function that is tied to a "class" (aka type that is uniquely named and extending the properties of another keyword
	// 		data type), you use the function through dot notation
	//
	// so imagining we have a type deck []string and a receiver function func (d deck) print() {} already defined in another file
	// 		we would use the print() function as so:
	//
	// e.g. start
	//
	// 			deck.go file
	// package main
	//
	// import "fmt"
	//
	// type deck []string
	//
	// func (d deck) print() {
	// 		for i, card := range d {
	// 			fmt.Println(i, card)
	// 		}
	// }
	//
	//			main.go file
	// package main
	//
	// import "fmt"
	//
	// func main() {
	// 		cards := deck{"Ace of Spades", "King of Hearts"}
	// 		cards = append(cards, "Queen of Clubs")
	// 		cards.print() 		// notice how the print() method is being accessed from the deck.go file, and is being used on a variable that has the deck data type stored in it, thus giving the variable access to the methods that have deck within their receiver function
	// }
	//
	// e.g. finish
	//
	// Because of both files using package main, and presumably being in the same folder, they both have access to each other's code
	// In this case the values of the cards variable are being passed as an argument into the print() function, in this case the value of cards is being
	// 		passed in as d, within the receiver function print(), and is being used as the range of the print() function.
	//
	// It is common practice to name the receiver argument (e.g. d) as the first letter, or first two letters, of the data type the receiver function is tied to [e.g. (d deck)]



GO NOTE:
using specific ranges of a slice

	// Very similarly to how you can access an item in an array (e.g. array[0] ) you can access items in a slice!
	// even more so, you can access a range of items in a slice with the syntax
	//
	// e.g.
	// slice[0:2]
	//
	// this would access the items STARTING FROM 0 and UP TO BUT NOT INCLUDING 2
	// so the items 0 and 1 would get selected from the slice[0:2] array
	//
	// you can also access all the items of a slice STARTING FROM A SPECIFIC POINT using the syntax
	// e.g.
	// slice[0:]
	// this would start from the 0 indexed item and select all items up to the end of the slice INCLUDING THE LAST ITEM
	// you could also do
	// e.g.
	// slice[2:]
	// this would start from the 2 indexed item and include all the indexed items starting from 2 onward, including the last item
	//
	// you can also access all the items from the beggining of the slice up until the specified point, NOT INCLUDING THE SPECIFIED POINT
	// e.g.
	// slice[:7]
	// this would start from the 0 indexed item and include all items UP TO BUT NOT INCLUDING the 7th indexed item



GO NOTE:
setting arguments in a function

	// when setting arguments in a function you must specify the name of the argument and it's type. Go is a statically typed langauge after all
	// e.g.
	// func dealCards (d deck, handSize int) {}
	//
	// in this case above the first argument is referenced in the function by the letter d and is of custom type defined deck
	// in the case of the second argument, it is referenced by handSize and is of data type int



GO NOTE:
remembering to set return types for functions

	// just like you declare a data type for arguments in functions, if your function returns at least one value, the value must be specified
	// 		at the end of the function declaration right before the curly braces
	//
	// e.g.
	// func deal(d deck, handSize int) (deck, deck) {}
	//
	// in this case above the function's name is deal
	// it has an argument d of type deck, and an argument handSize of type int
	// and returns two values, the first value will be of type deck and the second value will be of type deck



GO NOTE:
quick look at what a full receiver function looks like with arguments and two different returning values

	// func (d deck) shuffleDeck(arg1 string, arg2 []byte) ([]slice, int) {
	//
	// }



GO NOTE:
using the "ioutil" standard go package and the "strings" standard go package;
as well as the:
ioutil.   WriteFile(filename string, data []byte, perm os.FileMode) error {}
ioutil.   ReadFile(filename string) ([]byte, error) {}
strings.   Join(a []string, sep string) string {}
strings.   Split(s string, sep string) []string {}

	// 1st function. to go over is the WriteFile() standard go function
	// func WriteFile(filename string, data []byte, perm io.FileMode) error {}
	// a common perm value to use that says that anyone can read or write to the file is: 0666
	// e.g.
	// ioutil.WriteFile("new_file", byteDataType, 0666)
	// just like that and the file will read and be writable by anyone
	// the error value can be "nil" or "EOF"
	// if the error value is nil, everything ran well
	// if the error value is EOF, there was an actual error and the function did not successfully run
	//
	// 2nd function. then there is the ReadFile() standard go function
	// what it does is it reads a file that is all []byte data and returns an error
	// also to mention, a common perm value to use that says that anyone can read or write to the file is: 0666
	// e.g.
	// ioutil.ReadFile("new_file", byteDataType, 0666)
	// just like that and the file will read and be writable by anyone
	// the error value can be "nil" or "EOF"
	// if the error value is nil, everything ran well
	// if the error value is EOF, there was an actual error and the function did not successfully run
	//
	// 3rd function. is the Join() standard go function
	// this is for taking a slice of strings and concatonating all the values into one string
	//
	// 4th function. is the Split() standard go function
	// this is for taking a string that contains values seperated by some common symbol and puts each value into a index in a slice of strings
	// what it does is it writes a file within the directory the function's containing file is in and returns an error value



GO NOTE:
type conversion

	// type conversion is where you turn data from one type to antoehr, and is fairly simple
	// e.g.
	// strToByte = []byte("This will turn into a byte type")
	//
	// e.g.
	// intToFloat64 = float64(17)



GO NOTE:
using the type conversion and the ioutil and strings golang packages to turn a unique slice type into a byte type

	// type deck []string
	//
	// func (d deck) toByte() []byte {
	// 		return byte(strings.Join([]string(d), ","))
	// }
	//
	// the deck type turns into a []string, then string, then byte



GO NOTE:
using the type conversion and the ioutil and strings golang packages to turn a []byte file into a unique slice type

	// type deck []string
	//
	// func (b []byte) toDeck() {
	// 		return deck(strings.Split(string(b), ","))
	// }
	//
	// the []byte type turns into a string, then []string, then deck type



GO NOTE:
using the os golang library standard package, and the os.Exit() function

	// the os.Exit() function consists of one argument without a returning value
	// os.Exit(code int) would be the argument and type of argument
	// the way the code argument works is that, when the int 0 is passed into os.Exit() the function registers that nothing went wrong
	// 		and to not exit the program
	// but when the code argument is any other int besides 0, the os.Exit() function will activate and close the program



GO NOTE:
error handeling

	// when an error value is returned from a function it will either be: nil (which is a value type that signifies that there was no error)
	// 		or the value will be of some type that is not nil that signifies an error is present
	// typically the error value from a returning function is stored in the variable: err
	// after the value of the error is stored in err, then an if statement checks: if err != nil
	// this check allows you to handle the various cases an actual error is returned
	// typically the first thing you want to do when handeling an err that is != nil is to print the error with fmt.Println("Error:", err)
	// that way your error value is printed and you can see what the problem is with the error
	// then after that you want to write some code to handle how the error will affect the rest of the code
	// e.g.
	// if the err causes an infinite loop to happen it may be better to run the io.Exit() function and pass in an int argument into the function
	// 		that makes the io.Exit() function to close the program, e.g.



GO NOTE:
shuffling with the rand.Shuffle() function

	// the Shuffle() function is found within the "math/rand" package
	// the syntax of the function looks like
	//
	// e.g.
	// func Shuffle(n int, func (i, j int))
	//
	// the shuffle function takes a len(n) argurment typically, for this represents the length of the slice as an integer
	// as well as another line of code that defines what is done with the i and j variables
	//
	// e.g.
	// numbers := []int{1,2,3,4,5}
	// rand.Shuffle(len(numbers), func (i, j int) {
	// 		numbers[i], numbers[j] = numbers[j], numbers[i]
	// })
	//
	// this will shuffle the numbers according to a certain precalculated order based upon the "seed" or "source" of the order
	// this "seed" is the exact number (not necessarily an int) that the order is calculated from, as mentioned
	// you may change this order with another function as you shall now see



GO NOTE:
changing an RNG's seed number which is a "Source" data type, and the time.Now().UnixNano() functions as an argument

	// First of all it is worth noting that a RNG (random number generator) creates its random order by first have a seed value
	// 		this value, is a specific number that is used to create a specific "random" order. To change the random order in Golang
	// 		you must use a specific function to change that value
	// the function used to change the seed number or the "source" number is rand.NewNumber() found in the "math/rand" pakage
	//
	// e.g.
	// rand.NewSource(seed int64) Source {}
	//
	// you will notice that the function returns a Source data type. This data type is the order that and "math/rand" package function
	// 		will use for RNG's (and potentially other uses as of now)
	//
	// it is also worth noting that a good execution of using the rand.NewSource() function is with an argument that generates a new int64
	// 		type everytime the function is called. A good example of this is:
	//
	// e.g.
	// rand.NewSource(time.Now().UnixNano())
	//
	// this argument will call from the "time" standard package two functions that use the current time ".Now()" and then modifies that time
	// 		to a int64 data type with the ".UnixNano()" function



GO NOTE:
using a test.go file

	// a test file is used to test the functions of another file within that package, be it the main or whatever dependency package you've made
	// a test file is useful because it can test the logic in the other file without having to run the other file and cause errors to a
	// 		live database
	// to make a test file, create a file ending with _test.go
	// this file will first begin with the file it will be testing, for example. If you have a deck.go file that you want to write test
	// 		code for then you would create a deck_test.go file to contain the function that will test you deck.go file
	// you will need to import the standard package "testing" to have access to the ability run tests on functions from another file
	// to run all the tests in a package, use the command in the terminal:
	//
	// go test
	//
	// VScode source code editor is awesome because it will allow you to run individual tests inside of the editor, instead of
	// 		running the whole file and all the tests at once
	//
	// when making a test for a function, you want to ask yourself: "what do I really care about this function doing?" and then make a test for
	// 		everyone of those things


*/

func main() {}
