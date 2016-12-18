package main

import (
	"fmt"
	"math"
)

type borrowedBook struct {
	expectedDay   int
	expectedMonth int
	expectedYear  int
	actualDay     int
	actualMonth   int
	actualYear    int
	fee           int
}

func main() {
	borrowed := newBorrowedBook()
	borrowed.calculateLateFee()
	borrowed.printFee()
}

func newBorrowedBook() borrowedBook {
	b := borrowedBook{}
	b.readInData()
	return b
}

func (b *borrowedBook) readInData() {
	fmt.Scanf("%d", &b.actualDay)
	fmt.Scanf("%d", &b.actualMonth)
	fmt.Scanf("%d", &b.actualYear)
	fmt.Scanf("%d", &b.expectedDay)
	fmt.Scanf("%d", &b.expectedMonth)
	fmt.Scanf("%d", &b.expectedYear)
}

func (b *borrowedBook) calculateLateFee() {
	if b.actualYear < b.expectedYear {
		b.fee = 0
	} else if b.actualYear > b.expectedYear {
		b.fee = 10000
	} else {
		if b.actualMonth > b.expectedMonth {
			b.fee = 500 * int(math.Abs((float64(b.actualMonth) - float64(b.expectedMonth))))
		} else if b.actualDay > b.expectedDay {
			b.fee = 15 * int(math.Abs((float64(b.actualDay) - float64(b.expectedDay))))
		}
	}
}

func (b *borrowedBook) printFee() {
	fmt.Println(b.fee)
}
