package main

import (
	"fmt"
)

//方法一般和结构体绑定

type Person struct {
	id    int
	name  string
	age   int
	phone string
}

//定义一个Person的方法，获取id
func (person Person) getId() int {
	return person.id
}

//定义一个Person的方法，修改id
func (person Person) setId(id int) {
	person.id = id
}

func main() {
	person := &Person{
		id:    1,
		name:  "zhangsan",
		age:   18,
		phone: "15899271",
	}

	p := Person{
		id:    2,
		name:  "lisi",
		age:   20,
		phone: "15899272",
	}

	id := person.getId()
	fmt.Println("before set,person_id=", id)
	person.setId(9)
	id1 := person.getId()
	fmt.Println("after set,person_id=", id1)

	//自动取p得地址
	pId := p.getId()
	fmt.Println("before set,p_id=", pId)
	p.setId(10)
	pId1 := p.getId()
	fmt.Println("after set,p_id=", pId1)

}
