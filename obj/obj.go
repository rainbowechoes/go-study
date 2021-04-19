package obj

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Intro() {
	fmt.Printf("My name is %s, %d years old\n", p.Name, p.Age)
}

func (p *Person) Run() {
	fmt.Println("I'm running")
}

type Student struct {
	School string
	P      Person
}

func (stu *Student) Intro() {
	fmt.Printf("My name is %s, %d years old\n", stu.P.Name, stu.P.Age)
}

func (stu *Student) Run() {
	fmt.Println("Student is running")
}

func (stu Student) Compare(other Person) {
	nameEqual := other.Name == stu.P.Name
	ageEqual := other.Age == stu.P.Age

	if nameEqual && ageEqual {
		fmt.Printf("name and age are equal\n")
	} else {
		fmt.Printf("name and age are not equal\n")
	}
}

type Plugin interface {
	Install()
	Uninstall()
}

type PluginJ struct {
	Name string
}

type Software interface {
	Install()
	Uninstall()
}

func (plu *PluginJ) Install() {
	fmt.Println("Java Plugin Installing")
}

// 接口的实现类的选择子如果是值类型，那么接口的实例变量就必须是值类型；如果选择子是引用类型，那么实例变量就必须是引用类型
func (plu *PluginJ) Uninstall() {
	fmt.Println("Java Plugin Uninstall")
}