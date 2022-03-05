package BookLib

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//Struct operations
type Deletable interface {
	Delete(a int) []Book
}
type AuthorInfo struct {
	Name  string
	Birth string
}
type Book struct {
	Id          int
	BookName    string
	StockCode   string
	ISBNno      int
	PageNumber  int
	Price       int
	StockNumber int
	AuthorInfo
	IsDeleted bool
}

// Book struct constructor
func NewBook(bookName string, bookId int, authSlice string) Book {
	p := new(Book)
	p.Id = bookId
	p.BookName = bookName
	p.ISBNno = rand.Intn(100)
	p.PageNumber = rand.Intn(100)
	p.Price = rand.Intn(100)
	p.StockNumber = rand.Intn(100)
	p.StockCode = "book" + strconv.Itoa(p.ISBNno)
	p.Name = authSlice
	p.Birth = "1980"
	if p.Id%2 == 0 {
		p.IsDeleted = true
	} else {
		p.IsDeleted = false
	}
	return *p
}

//Function to calculate stock number
func (book Book) Stock(stockNo int) int {
	return book.StockNumber - stockNo
}

//Delete book which has specific ıd from book slice
func (book Book) Delete(index int, bookSlice []Book) {
	bookSlice = append(bookSlice[:index], bookSlice[index+1:]...)
	fmt.Println("New length", len(bookSlice))
}

//Listing operations
func List(bookStruct []Book) {
	for _, v := range bookStruct {
		fmt.Println(v)
	}
}

//Searching book with book name, sku number or author name
func Search(bookStruct []Book, searching string) {
	fmt.Printf("Searching Results: \n")
	tempSearching := strings.ToLower(searching)

	for _, v := range bookStruct {
		tempV := strings.ToLower(v.BookName)
		tempAuthorInfo := strings.ToLower(v.Name)
		tempSku := strings.ToLower(v.StockCode)
		if strings.Contains(tempV, tempSearching) {
			fmt.Printf("Founded in book name %s id: %d  \n", v.BookName, v.Id)
		}
		if strings.Contains(tempAuthorInfo, tempSearching) {
			fmt.Printf("Founded in authorInfo %s id: %d \n", v.Name, v.Id)
		}
		if strings.Contains(tempSku, tempSearching) {
			fmt.Printf("Founded in sku %s id: %d \n", v.StockCode, v.Id)
		}
	}
}

//Deletion operations
func Deletion(bookStruct []Book, ıd int) {
	var idCheck = false
	var counter = 0
	for _, v := range bookStruct {
		if v.Id == ıd {
			idCheck = true
			v.Delete(counter, bookStruct)
			return
		}
		counter = counter + 1
	}
	if idCheck == false {
		fmt.Printf("ERROR: book id not found \n")
	}

}

//Bought operations
func Buy(bookStruct []Book, ıd int, numberOfBooksToBuy int) {
	var idCheck = false

	for _, v := range bookStruct {
		if v.Id == ıd {
			idCheck = true
			if v.StockNumber >= numberOfBooksToBuy {
				newStock := v.Stock(numberOfBooksToBuy)
				fmt.Printf("ID: %d - Stock Number: Old:%d New:%d\n ", ıd, v.StockNumber, newStock)
				v.StockNumber = newStock
				return
			} else {
				fmt.Printf(" Inadequate Stock: You have to decrease number of books to buy \n")
			}
		}
	}
	if idCheck == false {
		fmt.Printf("ERROR: book id not found \n")
	}
}
