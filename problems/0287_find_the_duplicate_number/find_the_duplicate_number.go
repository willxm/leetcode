package problems

func findDuplicate(nums []int) int {
	d := make([]byte, len(nums))
	for _, v := range nums {
		if d[v] != 1 {
			d[v] = 1
		} else {
			return v
		}
	}
	return 0
}

//快慢指针思想, fast 和 slow 是指针, nums[slow] 表示取指针对应的元素
//注意 nums 数组中的数字都是在 1 到 n 之间的(在数组中进行游走不会越界),
//因为有重复数字的出现, 所以这个游走必然是成环的, 环的入口就是重复的元素,
//即按照寻找链表环入口的思路来做
func findDuplicate1(a []int) int {
	slow, fast := a[0], a[a[0]]
	for slow != fast {
		slow, fast = a[slow], a[a[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = a[slow], a[fast]
	}

	return slow
}
