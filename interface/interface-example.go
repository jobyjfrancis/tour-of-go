// https://www.alexedwards.net/blog/interfaces-explained

package main

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s \nAuthor: %s\n", b.Title, b.Author)
}

type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func main() {
	b := Book{"Alice in Wonderland", "Lewis Carrol"}
	fmt.Println(b) //Learn how this works
	WriteLog(b)

	c := Count(10)
	WriteLog(c)

}
