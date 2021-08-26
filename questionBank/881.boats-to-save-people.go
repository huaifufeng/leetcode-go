/**
  题目地址：https://leetcode-cn.com/problems/boats-to-save-people/

  第 i 个人的体重为 people[i]，每艘船可以承载的最大重量为 limit。
  每艘船最多可同时载两人，但条件是这些人的重量之和最多为limit。
  返回载到每一个人所需的最小船数。(保证每个人都能被船载)。


  <pre>
    输入：people = [1,2], limit = 3
    输出：1
    解释：1 艘船载 (1, 2)
  </pre>

  <pre>
    输入：people = [3,2,2,1], limit = 3
    输出：3
    解释：3 艘船分别载 (1, 2), (2) 和 (3)
  </pre>

  <pre>
    输入：people = [3,5,3,4], limit = 5
    输出：4
    解释：4 艘船分别载 (3), (3), (4), (5)
  </pre>

  PS：
    1 <= people.length <= 50000
    1 <= people[i] <= limit <= 30000

  解题：
    1、先排序，然后最重和最轻相加，这都大于限制的话，说明最重只能单独一个，否则一起 O(nlogn)
*/
package questionBank

import "sort"

func numRescueBoats(people []int, limit int) int {
	if len(people) == 1 {
		return 1
	}

	sort.Ints(people)

	num := 0
	for len(people) > 1 {
		if people[0]+people[len(people)-1] > limit {
			people = people[0 : len(people)-1]
			num++
		} else {
			people = people[1 : len(people)-1]
			num++
		}
	}

	if len(people) == 1 {
		num++
	}

	return num
}
