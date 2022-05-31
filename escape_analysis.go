package main

import "fmt"

type User struct {
	firstName string
	lastName  string
}

// this will be initialized in stack
func creatorFactory() User {
	return User{
		firstName: "Foo1",
		lastName:  "bar1",
	}
}

// this will be initialized in heap
func anotherCreatorFactory() *User {
	return &User{
		firstName: "foo2",
		lastName:  "bar2",
	}
}

func main() {
	user1 := creatorFactory()

	user2 := anotherCreatorFactory()

	fmt.Printf("name %s %s\n", user1.firstName, user1.lastName)
	fmt.Printf("name %s %s\n", user2.firstName, user2.lastName)
}
