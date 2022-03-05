package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-ElifBerilSayli/BookLib"
)

var bookStructSlice = make([]BookLib.Book, 5)
var bookSlice = make([]string, 5)
var authSlice = make([]string, 5)

// Error handling
var ErrInArgument = errors.New("ERROR: Arguments are invalid")
var ErrInId = errors.New("ERROR: Book ıd or another argument have problem. (Need integer value)")
var ErrInvalidIdNumber = errors.New("ERROR: Book ıd not found")

// Book related info initializations operations
func init() {
	counter := 1

	bookSlice[0] = "In Search of Lost Time"
	authSlice[0] = "Marcel Proust"

	bookSlice[1] = "Ulysses"
	authSlice[1] = "James Joyce"

	bookSlice[2] = "Don Quixote"
	authSlice[2] = "Miguel de Cervantes"

	bookSlice[3] = "The Great Gatsby"
	authSlice[3] = "Scott Fitzgerald"

	bookSlice[4] = "The Great Gatsby Second"
	authSlice[4] = "Scott Fitzgerald"

	for i := range bookStructSlice {
		n := BookLib.NewBook(bookSlice[i], counter, authSlice[i])
		counter = counter + 1
		bookStructSlice[i] = n
	}
}
func main() {

	args := os.Args

	var bookNameSlice []string
	var bookName string
	// Arguments and operations to list search buy and delete books
	if args[1] == "list" {
		BookLib.List(bookStructSlice)
		return
	}
	if args[1] == "search" {
		err := checkCommandSize(args)
		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {
			for i := 2; i < len(args); i++ {
				bookNameSlice = append(bookNameSlice, args[i])
			}
			bookName = strings.Join(bookNameSlice, " ")
			BookLib.Search(bookStructSlice, bookName)
		}
	}
	if args[1] == "buy" {

		err := checkCommandSize(args)

		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {

			ıd, err := strconv.Atoi(args[2])
			if err != nil {
				fmt.Printf("error running program: %s \n", err.Error())
			} else {
				err = checkIdValidError(ıd)
				if err != nil {
					fmt.Printf("error running program: %s \n", err.Error())
				} else {
					numberOfBooksToBuy, err := strconv.Atoi(args[3])
					if err != nil {
						fmt.Printf("error running program: %s \n", err.Error())
					} else {
						BookLib.Buy(bookStructSlice, ıd, numberOfBooksToBuy)
					}
				}
			}
		}
	}
	if args[1] == "delete" {

		deletionId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("error running program: %s \n", err.Error())
		} else {
			err = checkIdValidError(deletionId)
			if err != nil {
				fmt.Printf("error running program: %s \n", err.Error())
			} else {
				BookLib.Deletion(bookStructSlice, deletionId)
			}
		}

	}
}

//Error handling details
func checkCommandSize(args []string) error {
	if len(args) <= 2 {
		return ErrInArgument
	}
	return nil
}

func checkIdValidError(id int) error {
	for i := range bookStructSlice {
		if bookStructSlice[i].Id == id {
			return nil
		}
	}
	return ErrInvalidIdNumber
}
