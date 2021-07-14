package main

import "fmt"

func main() {
	//names := [] string {"北京", "天津", "广州", "上海", "深圳"}
	//for i, v := range names {
	//	fmt.Println("i : ", i, "v : ", v)
	//}
	//fmt.Println("length is : ", len(names), "capacity is : ", cap(names))
	//names = append(names, "西藏")		// 超出之后，一下子分配两倍
	//fmt.Println("length is : ", len(names), "capacity is : ", cap(names))
	//
	//nums := []int{}
	for i := 0; i < 50; i++ {
		nums = append(nums, i)
		fmt.Println("len is : ", len(nums), "cap is : " , cap (nums))
	}

	//上面重复命名了
	names := [7] string {"北京", "上海", "广州", "深圳", "洛阳", "南京", "秦皇岛"}
	names1 := [3] string {}
	names1[0] = names[0]
	names1[1] = names[1]
	names1[2] = names[2]

	names2 := names[0:3]
	fmt.Println("names2 : ", names2)

	names3 := names[:5]
	fmt.Println("names3 : ", names3)

	names4 := names[3:]
	fmt.Println("names4 : ", names4)

	names5 := names[:]
	fmt.Println("names5 : ", names5)

	sub1 := "helloworld"[5:7]
	fmt.Println("sub1:", sub1)


	//创建空切片的时候，可以明确指定切片的容量，这样就可以提高效率了
	//创建一个容量为20，当前长度是0的string切片
	//make的时候，初始值都为对应类型的0值，如0，false，空串
	str2 := make([] string, 10, 20)//第三个参数不是必须的，如果没填写，则默认与长度相同，第三个是容量
	fmt.Println("str2 : ", &str2[0])
	fmt.Println("length is : ", len(str2), " capacity is : ", cap(str2))
	str2[0] = "hello"
	str2[1] = "world"
	fmt.Println("str2 is ", str2)

	//如果想让切片完全独立于原来的数组，可以使用copy
	namesCopy := make([]string, len(names))
	copy(namesCopy, names[:])
	namesCopy[0] = "香港"
	fmt.Println("namesCopy:", namesCopy)
	fmt.Println("naemes本身:", names)
}
