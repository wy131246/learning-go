package main

import (
	"fmt"
)

func main() {
	s0 := []int{1, 2, 3, 4}
	s1 := s0[:2]        // [1,2] len(2)
	s2 := append(s1, 5) // [1,2,5]
	s3 := append(s1, 6, 7)
	// 前面三步一直都在s0的基础上修改

	s4 := append(s1, 8, 9, 0) //  超过了原来的长度，重新分配了一个新的slice底层结构，和s0 无关系了

	fmt.Println("s0:", s0)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	// simple example => 在s1 => s0[:2] 的基础上 append一个元素， 实际上相当于 修改s0[2]这个元素值
	a := []int{1, 2, 3, 4}
	t := append(a[:2], 100) // 修改了a[2] 为100
	fmt.Println(t, a)       //[1 2 100] [1 2 100 4]

}

/*
output:

s0: [1 2 6 7]
s1: [1 2]
s2: [1 2 6]
s3: [1 2 6 7]
s4: [1 2 8 9 0]

slice 是一个底层数组的切片，slice的引用操作都是操作在这个底层数组上。
所以 s1 => [1,2] s0 => [1,2,3,4]  s1 实质上是 s0[:2]
s2 => [1,2,5]  s0 => [1,2,5,4]  s2 实质上是s0[:3]
s3 => [1,2,6,7] s0 => [1,2,6,7] s3 实质上是s0[:4]
此时 s0 => [1,2,6,7]
所以 s1 s2 s3 就得到了对应的输出

在s1 => s0[:2] 的基础上 append一个元素， 实际上相当于 修改s0[2]这个元素值

而s4和上面s1-3不同，因为原来s0 的len 是 4， s1 append 3个元素，这时的底层数组是会5个元素，已经超过了原来s0的len，此时会重新分配过一个新的底层数组给s4，底层数组的前两个元素是s1[1,2]  后三个元素就是[8,9,0], 也就是现在有两个底层数组，s0-s3对应第一个，s4对应第二个。s4的值就是[1,2,8,9,0]了
*/
