package main

import "fmt"

func main() {

	//var idNames map[int] string		//空的 不能赋值
	//idNames[0] = "duke"
	//idNames[1] = "lily"
	//for key, value := range idNames {
	//	fmt.Println("key: ", key, " value : ", value)
	//}


	var idNames map[int] string
	idScore := make(map[int] float64)
	idNames = make(map[int] string, 10)

	idNames[0] = "duke"
	idNames[1] = "lily"
	for key, value := range idNames {
		fmt.Println("key is ", key, " value is ", value)
	}

	name9 := idNames[9]
	fmt.Println("name9: ", name9)
	fmt.Println("idSore[100] : ", idScore[100])

	value, ok := idNames[1]
	if ok {
		fmt.Println("id=1这个key是存在的，value为:", value)
	}
	value, ok = idNames[10]
	if ok {
		fmt.Println("id=10这个key是存在的，value为:", value)
	} else {
		fmt.Println("id=10这个key不存在，value为:", value)
	}

	fmt.Println("idNames删除前:", idNames)
	delete(idNames, 1)   //"lily"被kill
	delete(idNames, 100) //不会报错
	delete(idNames, 50)
	fmt.Println("idNames删除后:", idNames)
}
