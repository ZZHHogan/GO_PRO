package main


import "fmt"

func main() {
	name := "duke"
	fmt.Println("name : ", name)

	// 换行需要反引号，tab上面那个
	usage := `./a.out <optinon>
	-h help
	-a xx`
	fmt.Println("usage : ", usage)
	l1 := len(name)
	fmt.Println("length : ", l1)
	for i := 0; i < len(name); i++ {
		fmt.Printf("i : %d, v : %c\n", i, name[i])
	}
	// 拼接字符串
	i, j := "hello", "world"
	fmt.Println("i + j =  ", i + j)
	// const修饰常量，也是不能修改的
	const address = "beijing"
	const test = 100
	fmt.Println("address: ", address)
	fmt.Println("test: ", test)

	if i > j {
		;
	}
}