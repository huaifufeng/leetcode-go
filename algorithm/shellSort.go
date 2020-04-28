package algorithm

import "fmt"

//希尔排序

func ShellSort(nums []int) {
	gap := len(nums) / 2
	for gap > 0 {
		//gap 有分隔区间之后，相当于有gap个分组，需要每个分组单独处理
		for i := 0; i < gap; i++ {
			//每个分组单独进行冒泡排序
			for j := i + gap; j < len(nums); j += gap {
				temp := nums[j]
				k := j - gap

				for ; k > 0; k -= gap {
					if nums[k] > temp {
						nums[k+gap] = nums[k]
					} else {
						break
					}
				}
				nums[k+gap] = temp
			}
		}

		fmt.Println(nums)

		gap = gap / 2
	}
}

//ShellSort2  内部排序使用冒泡排序
func ShellSort2(nums []int) {
	gap := len(nums) / 2
	for gap > 0 {
		//gap 有分隔区间之后，相当于有gap个分组，需要每个分组单独处理
		for i := 0; i < gap; i++ {
			//每个分组单独进行冒泡排序
			for j := i; j < len(nums)-gap; j += gap {
				for k := j; k < len(nums)-gap-j; k += gap {
					if nums[k] > nums[k+gap] {
						nums[k], nums[k+gap] = nums[k+gap], nums[k]
					}
				}

			}
		}

		gap = gap / 2
	}
}
