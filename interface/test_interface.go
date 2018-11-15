package main

import (
	"fmt"
)

/*
接口是一系列方法声明的集合
一定是类型(struct)实现接口，所以在后面定义方法的时候，前面一定要加上类型（struct），表示该类型实现了该接口
实现接口一定要实现接口里的所有方法

eg.假设某公司有两个员工，一个普通员工和一个高级员工， 但是基本薪资是相同的，高级员工多拿奖金。计算公司为员工的总开支。
*/

//定义一个薪资计算器接口，每个员工都需要实现该接口，应为都要计算其工资
type SalaryCalculator interface {
	getSalary() int
}

//基本员工
type BasicEmployee struct {
	id       int
	basicpay int //基本工资
}

//高级员工
type AdvancedEmployee struct {
	id       int
	basicpay int //基本工资
	bonus    int //奖金
}

//BasicEmployee实现SalaryCalculator接口
func (baseicEmployee *BasicEmployee) getSalary() int {
	return baseicEmployee.basicpay
}

//AdvancedEmployee实现SalaryCalculator接口
func (advancedEmployee *AdvancedEmployee) getSalary() int {
	return advancedEmployee.basicpay + advancedEmployee.bonus
}

//计算总开支,s实际上是实现了SalaryCalculator接口的类型
func getTotalPay(s []SalaryCalculator) int {
	var totalpay int
	for _, employee := range s {
		totalpay = totalpay + employee.getSalary()
	}
	return totalpay
}

func main() {

	baseicEmployee := &BasicEmployee{
		id:       1,
		basicpay: 3000,
	}
	AdvancedEmployee := &AdvancedEmployee{
		id:       2,
		basicpay: 3000,
		bonus:    2000,
	}
	fmt.Println(baseicEmployee.getSalary())
	fmt.Println(AdvancedEmployee.getSalary())

	//实现了接口，就可以用接口来实现多态
	var employee []SalaryCalculator
	employee = append(employee, baseicEmployee)
	employee = append(employee, AdvancedEmployee)

	fmt.Println(employee[0].getSalary())
	fmt.Println(employee[1].getSalary())
}
