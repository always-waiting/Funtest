package main

import (
	"fmt"
	"regexp"
)

func MergeArray1(dest []string, src []string) (result []string) {
	result = make([]string, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return
}

func MergeArray2(dest []string, src []string) (result []string) {
	return append(dest, src...)
}

func MergeArray3(dest []string, src []string) []string {
	for _, val := range src {
		dest = append(dest, val)
	}
	return dest
}

//func main() {
//	a1 := []string{"a", "b", "e"}
//	a2 := []string{"c", "d", "a", "a", "a"}
//	m1 := MergeArray1(a1, a2)
//	m2 := MergeArray2(a1, a2)
//	m3 := MergeArray3(a1, a2)
//	fmt.Println(cap(m1))
//	fmt.Println(cap(m2))
//	fmt.Println(cap(m3))
//}

/*
func main() {
	a := make([]int, 0, 40)
	m := append(a, 1)
	fmt.Println(a)
	fmt.Println(m)
}
*/

type Person struct {
	Items []string
}

func (this *Person) AddItem(n string) {
	this.Items = append(this.Items, n)
}

func NewPerson() Person {
	return Person{
		Items: []string{},
	}
}

func main1() {
	a := NewPerson()
	a.AddItem("aa")
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := b[:0]
	fmt.Println(c)
	c = append(c, 11)
	c = append(c, 22)
	fmt.Println(c)
	fmt.Println(b)
	reg := regexp.MustCompile(`device_type\s*=\s*(?:'|")*([a-zA-Z_0-9]*)(?:'|")*`)
	expr := `aa{device_type="mq2"}[4m]>100`
	matcher := reg.FindStringSubmatch(expr)
	fmt.Println(matcher)
}
