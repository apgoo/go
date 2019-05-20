package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Speads", "Diamonds", "Hearts", "Clubs"}
	cardVal := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuit {
		for _, val := range cardVal {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards

}


func (d deck) print() {
	for i, car := range d {
		fmt.Println(i, car)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toSaveFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFile(fileName string) deck {
	bySli, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bySli), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
