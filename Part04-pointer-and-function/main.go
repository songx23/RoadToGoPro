package main

import "fmt"

func main() {
	//// Pointer
	pointer := "I'm a string, not a pointer"
	// declare a pointer type
	var pointerType *string
	pointerType = &pointer
	// declare a pointer directly
	p := &pointer
	fmt.Printf("String: %s\nPointerType: %v\nPointer: %v\n", pointer, pointerType, p)
	// output:
	//   String: I'm a string, not a pointer
	//   PointerType: 0xc0000xxxxx (a memory address in form of a hexadecimal number)
	//   Pointer: 0xc0000xxxxx (a memory address in form of a hexadecimal number)

	//// Pointer dereference
	fmt.Printf("Dereference pointer: %v\n", *p)
	// output:
	//    Dereference pointer: I'm a string, not a pointer

	*p = "Pretending to be a pointer now."
	fmt.Printf("New value set via pointer: %v\n", *p)
	fmt.Printf("Variable value changed: %s\n", pointer)
	// output:
	//    New value set via pointer: Pretending to be a pointer now.
	//    Variable value changed: Pretending to be a pointer now.

	//// Nil pointer dereference error
	// Uncomment below to see the famous nil pointer exception
	// var nilPointer *int
	// fmt.Printf("nil pointer: %v\n", *nilPointer)
	// output:
	//    panic: runtime error: invalid memory address or nil pointer dereference

	//// Value/Pointer receivers in functions

	person := Person{
		FirstName: "Jon",
		LastName:  "Snow",
		Age:       23,
	}
	// function with a value receiver
	fullName := person.FullName()
	fmt.Print(fullName)
	// output: Jon Snow
	// value receiver does not change the original variable
	fmt.Println(person.Age)
	// output: 23

	// function with a pointer receiver
	person.ChangeName("Aegon", "Targaryen")
	// pointer receiver does change the original variable
	fmt.Printf("Jon Snow is %v\n", person)

	//// Interface implementation
	var poodle Animal
	poodle = &Dog{
		Name: "Bolt",
	}
	// implementing with value receivers
	poodle.Eat()   // output: Bolt eats meat
	poodle.Drink() // output: Bolt drinks water

	var ragdoll Animal
	ragdoll = &Cat{Name: "Lightning"}
	// implementing with pointer receivers
	ragdoll.Eat()   // output: Lightning eats fish
	ragdoll.Drink() // output: Lightning drinks milk

	//// When to use pointers
	// 1. When you need to have a nilable (nullable) type
	invite := Invitation{FirstName: "Donald", LastName: "Trump"}
	fmt.Printf("%v\n", invite)
	// output: {Donald Trump <nil>}

	// Uncomment the next line and we will get a nil pointer dereference error. We need to watch out for this.
	// fmt.Printf("%v\n", *invite.Going)

	// 3. When in need of singletons
	// Passing a singleton by its pointer
	pClient := NewPointerSlackClient("slack.com", "apiKey")
	pService := NewPointerService(pClient)
	pService.Description()
	// output: Pointer Client is a singleton.

	// Passing a singleton by its value
	vClient := NewValueSlackClient("slack.com", "apiKey")
	vService := NewValueService(vClient)
	vService.Description()
	// output: Value Client is not a singleton.
}

//// Function examples
// function with one input argument and one output argument
func isCountryCovidSafe(country string) bool {
	// casting dark magic spell to determine whether requested country is covid safe or not
	return true // only to make the compiler happy
}

// function with two input argument (same type) and one output argument
func isStateCovidSafe(country, state string) bool {
	// casting dark magic spell to determine whether the requested state is covid safe or not
	return true // only to make the compiler happy
}

// function with no input and output argument
func doNothing() {
	fmt.Println("This factory does nothing. Not very useful.")
}

//// Receivers
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// function with value receiver
func (p Person) FullName() string {
	p.Age = 100
	return fmt.Sprintf("%s %s\n", p.FirstName, p.LastName)
}

// function with pointer receiver
func (p *Person) ChangeName(newFirstName, newLastName string) {
	if newFirstName != "" {
		p.FirstName = newFirstName
	}

	if newLastName != "" {
		p.LastName = newLastName
	}
}

//// Interface implementation
type Animal interface {
	// function sets
	Eat()
	Drink()
}

// implement the Animal interface
type Dog struct{ Name string }

// implementing using value receiver
func (d Dog) Eat() {
	fmt.Printf("%s eats meat\n", d.Name)
}
func (d Dog) Drink() {
	fmt.Printf("%s drinks water\n", d.Name)
}

// implement the Animal interface
type Cat struct{ Name string }

// implementing using pointer receiver
func (c *Cat) Eat() {
	fmt.Printf("%s eats fish\n", c.Name)
}
func (c *Cat) Drink() {
	fmt.Printf("%s drinks milk\n", c.Name)
}

//// Using pointer to store nil value in bool type
type Invitation struct {
	FirstName string
	LastName  string
	Going     *bool
}

//// Singleton
type SlackClient struct {
	URL    string
	APIKey string
}

func NewPointerSlackClient(url, apiKey string) *SlackClient {
	return &SlackClient{
		URL:    url,
		APIKey: apiKey,
	}
}

func NewValueSlackClient(url, apiKey string) SlackClient {
	return SlackClient{
		URL:    url,
		APIKey: apiKey,
	}
}

type PointerService struct {
	SlackClient *SlackClient
}

func (*PointerService) Description() {
	fmt.Println("Pointer Client is a singleton.")
}

func NewPointerService(client *SlackClient) *PointerService {
	return &PointerService{
		SlackClient: client,
	}
}

type ValueService struct {
	SlackClient SlackClient
}

func NewValueService(client SlackClient) ValueService {
	return ValueService{
		SlackClient: client,
	}
}

func (ValueService) Description() {
	fmt.Println("Value Client is not a singleton.")
}
