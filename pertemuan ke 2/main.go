package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("ini project 2")

	//condition if
	if nilaiStudent := 38; nilaiStudent > 70 {
		fmt.Println("lulus")
	}
	if nilaiStudent := 38; nilaiStudent < 70 {
		fmt.Println("tidak lulus")
	}
	//end condition if

	//condition switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	//switch with relational operators
	var scores = 6

	switch {
	case scores == 8:
		fmt.Println("Perfect")
	case (scores < 8) && (scores > 3):
		fmt.Println("Not Bad")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("you need to learn more")
		}
	}

	//switch (fallthrough keyword)
	var nilai = 6

	switch {
	case nilai == 8:
		fmt.Println("Perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("Not Bad")
		fallthrough
	case nilai < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("Study Harder")
			fmt.Println("you don't have a good score yet")
		}
	}

	//nested condition
	var nilais = 10

	if nilais > 7 {
		switch nilais {
		case 10:
			fmt.Println("Perfect")
		default:
			fmt.Println("Nice!")
		}
	} else {
		if nilais == 5 {
			fmt.Println("not bad")
		} else if nilais == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if nilais == 0 {
				fmt.Println("try harder")
			}
		}
	}
	//end condition

	//looping

	//lopings (first way of looping)
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	//loopings (second way of looping)
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}

	//loopings (third way of looping)
	var a = 0

	for {
		fmt.Println("Angka", a)
		a++
		if a == 3 {
			break
		}
	}

	//loopings (break and continue keywords)
	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//loopings (nested looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	//loopings (label)
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
	//end looping

	//Array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var string1 = [3]string{"Airell", "Nanda", "Mailo"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", string1)

	//Array (Modify Element Though Index)
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "manggo"

	fmt.Printf("%#v\n", fruits)

	//Array (Loop through elements)
	var fruit = [3]string{"Apple", "banana", "mango"}
	//cara Pertama
	for i, v := range fruit {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))

	//cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruit[i])
	}

	//Array (Multi-dimensional array)
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	//end Array

	//Slice
	var buah = []string{"apple", "banana", "mango"}

	_ = buah

	//slice (make function)
	var buah1 = make([]string, 3)
	_ = buah1
	fmt.Printf("%#v\n", buah1)

	//slice (append function)
	var buah2 = make([]string, 3)
	_ = buah2

	buah2[0] = "apple"
	buah2[1] = "banana"
	buah2[2] = "mango"

	fmt.Printf("%#v\n", buah2)

	//slice (apend function with ellipsis)
	var buah3 = []string{"apple", "banana", "mango"}
	var buah4 = []string{"durian", "pineapple", "starfruit"}

	buah3 = append(buah3, buah4...)

	fmt.Printf("%#v\n", buah3)

	//slice (copy function)
	var buah5 = []string{"apple", "banana", "mango"}
	var buah6 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(buah5, buah6)

	fmt.Println("buah5 =>", buah5)
	fmt.Println("buah6 =>", buah6)
	fmt.Println("copied element =>", nn)

	//slice (slicing)
	var buah7 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var buah8 = buah7[1:4]
	fmt.Printf("%#v\n", buah8)

	var buah9 = buah7[0:]
	fmt.Printf("%#v\n", buah9)

	var buah10 = buah7[:3]
	fmt.Printf("%#v\n", buah10)

	var buah11 = buah7[:]
	fmt.Printf("%#v\n", buah11)

	//slice (Combining slicing and append)
	var buah12 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	buah12 = append(buah12[:3], "rambutan")

	fmt.Printf("%#v\n", buah12)

	//slice (Backing Array)
	var buah13 = []string{"apple", "mango", "durian", "pineapple", "banana"}

	var buah14 = buah13[2:4]

	buah14[0] = "rambutan"

	fmt.Println("buah13 => ", buah13)
	fmt.Println("buah14 => ", buah14)

	//slice (cap function)
	var buah15 = []string{"apple", "mango", "durian", "pineapple"}

	fmt.Println("buah15 cap :", cap(buah15))
	fmt.Println("buah15 len :", len(buah15))

	fmt.Println(strings.Repeat("#", 20))

	var buah16 = buah15[0:3]

	fmt.Println("buah16 cap :", cap(buah16))
	fmt.Println("buah16 len :", len(buah16))

	fmt.Println(strings.Repeat("#", 20))

	//slice (creating a new backing array)
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars: ", cars)
	fmt.Println("new cars: ", newCars)
	//end slice

	//Map
	var person map[string]string //deklarasi

	person = map[string]string{} //inisialisasi

	person["name"] = "Airell"

	person["age"] = "23"

	person["address"] = "Jalan Sudirman"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	var person1 = map[string]string{
		"name":    "Airell",
		"age":     "23",
		"address": "Jalan Sudirman",
	}

	fmt.Println("name: ", person1["name"])
	fmt.Println("age: ", person1["age"])
	fmt.Println("address: ", person1["address"])

	//map (Looping with Map)
	var person2 = map[string]string{
		"name":    "Airel",
		"age":     "12",
		"address": "jalan Sudirman",
	}

	for key, value := range person2 {
		fmt.Println(key, ":", value)
	}

	//map (Deleting value)
	var person3 = map[string]string{
		"name":    "Airel",
		"age":     "12",
		"address": "jalan Sudirman",
	}

	fmt.Println("before deleting: ", person3)

	delete(person3, "age")

	fmt.Println("after deleting: ", person3)

	//map (detecting a value)
	var person4 = map[string]string{
		"name":    "Airel",
		"age":     "12",
		"address": "jalan Sudirman",
	}

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person4, "age")

	value, exist = person4["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	//Map (Combininb slice with map)
	var people = []map[string]string{
		{"name": "Airell", "age": "23"},
		{"name": "Nanda", "age": "23"},
		{"name": "Mailo", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
	// End Map

	//Aliase
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type: %T\n", hour)
	//end Aliase

	//String In Depth
	city := "Jakarta"

	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	//looping over string (byte-by-byte)
	var city1 []byte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(city1))

	//looping over string (rune-by-rune)
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	//looping over string (rune-by-rune)
	str3 := "makan"
	str4 := "mânca"

	fmt.Printf("str3 byte length => %d\n", utf8.RuneCountInString(str3))
	fmt.Printf("str4 byte length => %d\n", utf8.RuneCountInString(str4))

	//looping over string (rune-by-rune)
	str5 := "mânca"

	for i, s := range str5 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}
	//end strings in dept
}
