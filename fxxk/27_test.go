package fxxk

import (
	"fmt"
	"testing"
)

/*
27. 移除元素

给你一个数组nums和一个值val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度。
不要使用额外的数组空间，你必须仅使用 O(1)额外空间并原地修改输入数组。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
  print(nums[i]);
}

示例 1：
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
示例 2：
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。

*/

func removeElement(nums []int, val int) int {
	//快慢两个指针
	//快指针，从index=1位置开始每次都往前走1，判断快指针指向的数是否等于val
	//慢指针从index=0 处开始
	//快指针如果指向的数等val，则继续+1。如果不等val，则这个数需要写到数组前面的位置，然后加1继续。
	//慢指针指向要被替换的数。一旦被替换了，+1.否则一直原地不动
	if len(nums) == 0 {
		return 0
	}
	fast := 1
	slow := 0
	if nums[0] != val {
		slow = 1
	}
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func TestAl27(t *testing.T) {
	fmt.Println(removeElement([]int{1, 2, 3, 3, 2, 2, 3}, 3))
}