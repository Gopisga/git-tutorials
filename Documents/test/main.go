package main

import (
	"fmt"
	//"os"
	//"text/template"
)

//struct
/*type secret struct {
	Name   string
	Secret string
}*/

//interface
/*type language struct {
	owner string
}

func (language) coding() {
	fmt.Println("coding the web app")
}

type programming struct {
	owner string
}

func (programming) coding() {
	fmt.Println("design the app")
}

type engineer interface {
	coding()
}

func programme(reactjs engineer) {
	reactjs.coding()
	}*/

func main() {

	fmt.Println("こんにちは　世界!")

	//template
	/*stars := []language{{"GOPI"}, {"GOPI"}}
	for _, star := range stars {
		star.coding()
	}
	fmt.Println()

	engineers := []programming{{"GOPI"}, {"GOPI"}}
	for _, engineer := range engineers {
		engineer.coding()
	}
	fmt.Println()

	reactjs := []engineer{language{"GOPI"}, programming{"GOPI"}}
	for _, reactjss := range reactjs {
		programme(reactjss)
	}*/

	/*secret := secret{"GOPI", "react.js&vue.js"}
	templatepath := "/Users/gopi.b/Documents/test/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatepath)
	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)
	if err != nil {
		fmt.Println(err)
	}*/

	//template ends up here

	//block
	/*{
		var integer = 6

		fmt.Println(integer)
	}*/
	/*1.)integer declaration :
	var integer, string = 1, "string"

	fmt.Println(integer, string)*/

	//2.)short declaration :
	/*boolean := false

	fmt.Println(boolean)*/

	/*3.)pointer declaration :
	//var X int = 1
	//var p *int = &X
	fmt.Println(*p)*/

	//4.)type declaration :
	/*type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)
	fmt.Println(f,c)*/

	//float
	/*pi := 3.141
	pi = float64(3.141)
	fmt.Printf("%T", pi)*/

	// andpad := 1e100
	//fmt.Println(andpad)

	/*float
	a := 1
	b := 2.12
	b = float64(a)
	fmt.Println(b)
	a = int(b)
	fmt.Println(a)*/

	/*MATH LIBRARY
	fmt.Println(math.Ceil(3.999999))
	fmt.Println(math.Floor(3.9934))
	fmt.Println(math.Min(1, 0))
	fmt.Println(math.Max(1, 0))
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Sqrt(64))
	fmt.Println(math.Pow(2,3))*/

	/*STRING and STRING LIBARY AND BULIDER
	s := "HELLO WORLD"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%c", s[0])*/

	/*SLICE
	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s , %s", abc, b)
	fmt.Println()
	abc2 := string(b)
	fmt.Println(abc2)
	//SLICE
	fmt.Println("こんにちは　世界!"[0])
	fmt.Println(len("こんにちは　世界!"))
	fmt.Println("こんにちは　世界!"[0:25])*/

	/*arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))*/

	// LIBRARY
	/*fmt.Println(strings.ToUpper("hello world!"))
	fmt.Println(strings.ToLower("hello world!"))
	fmt.Println(strings.HasPrefix("hello world!", "hello"))
	fmt.Println(strings.HasSuffix("hello world!", "world!"))
	fmt.Println(strings.Replace("hello こんにちは world!", "hello", "こんにちは", 1))
	fmt.Println(strings.Replace("hello こんにちは world!", "world", "世界", 1))*/

	//bulider
	/*var sb strings.Builder
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("hello")
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("This is a string builder:", sb.String())
	sb.Reset()
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("This is a string builder:", sb.String())*/

	/*const five int = 5
	fmt.Println(five)*/

	//array
	/*b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	for _, v := range b {
		fmt.Println(v)

	}*/

	/*array := [...]int{0: 1}
	fmt.Println(array)
	array2 := [...]int{1: 1, 4: 5}
	fmt.Println(array2)

	strarray := [...]string{"string1", "string2"}
	fmt.Println(strarray)*/

	/*struct
	type andpad struct {
		name     string
		language []string
		level    int
	}
	GOPI := andpad{name: "GOPI"}
	GOPI = andpad{
		name: "GOPI", language: []string{"react.js vue.js"}, level: 99,
	}
	fmt.Println(GOPI)
	fmt.Println()
	fmt.Println(GOPI.name)
	fmt.Println(GOPI.language)
	fmt.Println(GOPI.level)
	GOPI.level++
	GOPI.language = append(GOPI.language, "nuxt.js")
	fmt.Println(GOPI)
	fmt.Println()
	fmt.Println(GOPI.name)
	fmt.Println(GOPI.language)
	fmt.Println(GOPI.level)*/

	/*maps
	var m map[string]string
	fmt.Println(m == nil)
	m = map[string]string{}
	fmt.Println(m == nil)

	fmt.Println(len(m))
	m = make(map[string]string, 5)
	fmt.Println(len(m))*/

	/*array
	fixed := [...]int{0, 1, 2}
	fmt.Println(fixed)
	a := []int{0, 1}
	fmt.Println(a)
	fmt.Println(len(a))
	a = []int{0, 1, 2}
	fmt.Println(a)
	a = append(a, 3, 4, 5, 6)
	fmt.Println(a)
	fmt.Println(cap(a), len(a))

	b := make([]int, 5)
	fmt.Println(b)
	fmt.Println()

	fmt.Println(a)
	a = a[0:7]
	fmt.Println(a)
	a = a[0:8]
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	a = a[0:7]
	fmt.Println(a)*/

}
