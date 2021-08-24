package main

import (
	"net/http"

	"github.com/rainbowechoes/go-study/rgin"
	"github.com/gin-gonic/gin"
	"github.com/rainbowechoes/go-study/obj"
)

func personDoSports(person *obj.Person) {
	person.Intro()
	person.Run()
}

func installPlugin(plugin obj.Plugin) {
	plugin.Install()
}

func installSoftware(software obj.Software) {
	software.Install()
}

var ablums = []rgin.Album{
	{ID: "1", Title: "Blue", Artist: "yangjian", Price: 56.99},
	{ID: "2", Title: "Green", Artist: "yangjian", Price: 34.22},
}

func getAlbums(c *gin.Context) {
	c.IndentedJson(http.StatusOK, ablums)
}

func main() {
	// person := &obj.Person{
	// 	Name: "yangjian",
	// 	Age:  22,
	// }
	// fmt.Printf("My type is %T\n", person)
	// person.Intro()

	// stu := obj.Student{
	// 	School: "yizhong",
	// 	P: obj.Person{
	// 		Name: "yangjian",
	// 		Age:  21,
	// 	},
	// }
	// stu.Intro()

	// compareM := stu.Compare
	// compareM(*person)
	// comPareME := obj.Student.Compare
	// fmt.Printf("method variable type is %T\n", compareM)
	// // 不管原方法的选择子是值类型还是引用类型，变成方法表达式后，选择子所对应的参数必须都是值类型，类似 Java 中的静态引用
	// comPareME(stu, *person)
	// fmt.Printf("method expression type is %T\n", comPareME)

	// personDoSports(person)
	// var pers obj.Person = stu.P
	// personDoSports(&pers)

	// var pluginJ obj.Plugin = &obj.PluginJ{
	// 	Name: "lombok",
	// }

	// installPlugin(pluginJ)
	// installSoftware(pluginJ)
	// plj, ok := pluginJ.(*obj.PluginJ)
	// if ok {
	// 	plj.Uninstall()
	// }

	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 10}

	// streamChan := make(chan int)
	// outChan := make(chan string)

	// go parallel.StreamOf(nums, streamChan)
	// go parallel.MapToStr(streamChan, outChan, func(num int) string {
	// 	return strconv.Itoa(num * num)
	// })

	// for i := range outChan {
	// 	fmt.Printf("out is %s\n", i)
	// }
}
