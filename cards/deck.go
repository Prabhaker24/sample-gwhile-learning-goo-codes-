package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of deck
//which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuite := []string{"spades", "heart", "diamond", "clubs"}
	cardValues := []string{"ace", "two", "three", "four"}

	for _, val1 := range cardSuite {
		for _, val2 := range cardValues {
			cards = append(cards, val2+" of "+val1)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, no int) (deck, deck) {
	return d[:no], d[no:]
}

func (d deck) writeTOFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(strings.Join([]string(d), ",")), 0666)

}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	temp := deck(strings.Split(string(bs), ","))
	fmt.Println(temp)
	return temp
}
