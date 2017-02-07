/*
Copyright 2017 Nicholas Leskiw <nleskiw@gmail.com>

This file is part of goplaycards

goplaycards is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

goplaycards is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with goplaycards.  If not, see <http://www.gnu.org/licenses/>.

*/

package main

import (
	"fmt"

	"github.com/nleskiw/goplaycards/deck"
)

func main() {

	fmt.Println("goplaycards Copyright (C) 2017  Nicholas Leskiw")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY; for details please visit")
	fmt.Println("<https://www.gnu.org/licenses/gpl-3.0.txt>")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Printf("under certain conditions. Please see the GPLv3 license at the URL listed above.\n\n")

	var d deck.Deck
	d.Initialize()
	for i := 5; i > 0; i-- {
		fmt.Printf("First Card: %+v\n", d.Cards[0])
		fmt.Printf("length of deck: %d\n", len(d.Cards))
		fmt.Printf("Shuffing...\n")
		e := d.Shuffle()
		if e != nil {
			panic(e)
		}
	}

	hand, e := d.Draw(5)
	if !(e == nil) {
		panic(e)
	}

	fmt.Printf("\n\n\nDrawn Hand:")
	for _, c := range hand {
		fmt.Printf(" %s ", c.ToStr())
	}
	fmt.Printf("\n\n")

	fmt.Printf("Length of deck: %d\n", len(d.Cards))

	pair := false
	for i, c := range hand {
		for j := i + 1; j <= len(hand)-1; j++ {
			if c.Equal(&hand[j]) {
				pair = true
			}
		}
	}

	if pair {
		fmt.Printf("You have a pair!\n")
	} else {
		fmt.Printf("You don't have a pair.\n")
	}

	return
}
