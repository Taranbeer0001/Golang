package main

import (
	"fmt"
	"math"
	"strings"
)

type person struct {
	first_name     string
	last_name      string
	favourite_city []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle   vehicle
	fourWheel bool
}

type sedan struct {
	vehicle vehicle
	luxury  bool
}

type square struct {
	l int
	w int
}

type circle struct {
	r float64
}

func (s square) Area() int {
	return s.l * s.w

}

func (c circle) Area() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	Areaofshape()
}

func (s square) Areaofshape() int {
	return s.l * s.w
}

func (c circle) Areaofshape() float64 {
	return math.Pi * c.r * c.r
}

func repetition(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func main() {

	/*1. Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
		● first name
		● last name
		● favorite Country
	Create two VALUES of TYPE person.
	Print out the values, ranging over the elements in the slice which stores the favorite Countries.*/

	fmt.Println("STRUCTURE")

	rec1 := person{
		first_name:     "John",
		last_name:      "Snow",
		favourite_city: []string{"USA", "Australia", "Italy"},
	}

	rec2 := person{
		first_name:     "Tyrion",
		last_name:      "Lannister",
		favourite_city: []string{"Austria", "Turkey", "Dubai"},
	}

	fmt.Println(rec1)
	fmt.Println(rec2)

	//2. Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	//Access each value in the map. Print out the values, ranging over the slice.

	r1 := map[string][]string{
		"first_name":     {"Tyrion"},
		"last_name":      {"Lannister"},
		"favourite_city": {"Austria", "Turkey", "Dubai"},
	}

	fmt.Println(r1)

	r2 := map[string][]string{
		"last_name": {"Snow", "Lannister"},
	}

	for i, k := range r2 {
		fmt.Println(i, k)
	}

	/*3. Create a new type: vehicle.
	  ○ The underlying type is a struct.
	  ○ The fields:
	            ■ doors
	            ■ color
	*/
	value1 := truck{
		vehicle: vehicle{
			doors: 6,
			color: "black",
		},
		fourWheel: false,
	}
	fmt.Println(value1)

	value2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}
	fmt.Println(value2)

	fmt.Println(value1.vehicle)
	fmt.Println(value2.luxury)

	/*4.
		● create a type SQUARE
		● create a type CIRCLE
		● attach a method to each that calculates AREA and returns it
		          ○ circle area= π r 2
		          ○ square area = L * W
		● create a type SHAPE that defines an interface as anything that has the AREA method
	    ● create a func INFO which takes type shape and then prints the area
	    ● create a value of type square
	    ● create a value of type circle
	    ● use func info to print the area of square
	    ● use func info to print the area of circle
	*/
	square := square{
		l: 10,
		w: 10,
	}
	fmt.Println(square.Area())

	circle := circle{
		r: 10,
	}

	fmt.Println(circle.Area())

	fmt.Println(square.Areaofshape())

	fmt.Println(circle.Areaofshape())

	//5. write a program to count the occurrence of each word in a sentence and return a map
	//   containing word and number of occurrences of it in the given sentence

	input := "Hey, I'm Taran . Taran nickname is Goku . "
	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)
	}

}
