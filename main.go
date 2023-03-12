var x string = "Hello, World"

func main() {
 fmt.Println(x)
}

func f() {
 fmt.Println(x)
}

// The f function now has access to the x variable. Now suppose that we wrote this
// instead:

func main() {
 var x string = "Hello, World"
 fmt.Println(x)
}

func f() {
 fmt.Println(x)
}

// If you run this program, you should see an error:
// .\main.go:11: undefined: x

// The compiler is telling you that the x variable inside of the f function doesn’t exist. It
// only exists inside of the main function. The range of places where you are allowed to
// use x is called the scope of the variable. According to the language specification, “Go
// is lexically scoped using blocks.” Basically, this means that the variable exists within
// the nearest curly braces ({ }), or block, including any nested curly braces (blocks),
// but not outside of them. Scope can be a little confusing at first; as we see more Go
// examples, it should become more clear.
