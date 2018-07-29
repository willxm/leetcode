package problem

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	rows := make([][]int, 0)
	//-1, 0, 1, 2, -1, -4
	//-4,-1,-1,0,1,2
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return nil
	}
	i := 0
	j := l - 1
	for {
		if j-i < 2 {
			break
		}
		fmt.Println(i, j)
		m := nums[i] + nums[j]
		if m > 0 {
			k := i + 1
			for {
				fmt.Println(k)
				if k == j {
					j--
					for {
						if i == j {
							break
						}
						if nums[j] == nums[j+1] {
							j--
						} else {
							break
						}
					}

					break
				}
				if m+nums[k] == 0 {
					fmt.Println(nums[i], nums[k], nums[j])
					rows = append(rows, []int{nums[i], nums[k], nums[j]})
					i++
					if i == j {
						break
					}
					break
				}
				k++
				for {
					if k == j {
						break
					}
					if nums[k] == nums[k-1] {
						k++
					} else {
						break
					}
				}
			}
		} else {
			k := j - 1
			for {
				fmt.Println(k)
				if k == i {
					i++
					for {
						if i == j {
							break
						}
						if nums[i] == nums[i-1] {
							i++
						} else {
							break
						}
					}
					break
				}
				if m+nums[k] == 0 {
					fmt.Println(nums[i], nums[k], nums[j])
					rows = append(rows, []int{nums[i], nums[k], nums[j]})

					j--
					if i == j {
						break
					}
					break
				}
				k--
				for {
					if k == i {
						break
					}
					if nums[k] == nums[k+1] {
						k--
					} else {
						break
					}
				}
			}
		}
	}
	return rows
}
