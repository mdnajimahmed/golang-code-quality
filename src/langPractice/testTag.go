package langPractice

import (
	"fmt"
	"reflect"
)

type Person struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty" bson:"middle_name,unique"`
}

type myInt int

func TestTag() {
	p := Person{"Najim", "Ahmed", "najim.ju@gmail.com"}
	t := reflect.TypeOf(p)

	for _, fieldName := range []string{"FirstName", "LastName", "MiddleName"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: User.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'json': %q\n", field.Tag.Get("json"))
		fmt.Printf("\tValue of 'bson': %q\n", field.Tag.Get("bson"))
	}

	var i int = 1
	fmt.Printf("%T %v\n", i, i)

	var mi myInt = 2
	fmt.Printf("%T %v\n", mi, mi)

	//*** does not work,though their underlying type is same, need conversion!
	//var newInt int = mi
	var newInt int = int(mi)
	fmt.Printf("%T %v\n", newInt, newInt)

	//var newInt2 myInt = 5
	// *** Type assertions can only take place on interfaces. This code wont work
	//var newInt2val int = newInt2.(int)
	//fmt.Printf("%T %v\n", newInt2, newInt2)

	var greeting interface{} = 42

	switch g := greeting.(type) {
	case string:
		fmt.Println("g is a string with length", len(g))
	case int:
		fmt.Println("g is an integer, whose value is", g)
	default:
		fmt.Println("I don't know what g is")
	}
	// g is available ony inside, out side it's not available.
	// raw string literal
	s := `te st`
	fmt.Println(s)
	// string is a slice of byte
	bs := []byte(s)
	fmt.Printf("%T -> %v\n", s, bs)

	js := "構 構"
	jbs := []byte(js)
	fmt.Printf("%T -> %v\n", jbs, jbs)

	for i, v := range js {
		fmt.Printf("%v -> %#U\n", i, v)
	}
	jpbytes := []uint8{230, 167, 139, 32, 230, 167, 139}
	fmt.Printf(string(jpbytes))

	fmt.Println(&s)
	s += "immutability test"
	fmt.Println(&s)

	testBytes := []uint8{130, 97, 65}
	fmt.Printf(string(testBytes))
	decodeChar("Ɏ") // [201 142]

}

func decodeChar(s string) {
	fmt.Println("")
	fmt.Printf("%v %U ", []byte(s), s[0])
	fmt.Printf("%+q\n", s) // print unicode codepoint
}
