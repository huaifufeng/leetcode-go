/**
  题目地址：https://leetcode-cn.com/problems/all-paths-from-source-to-target/

  给你一个有n个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
  二维数组的第 i 个数组中的单元都表示有向图中 i 号节点所能到达的下一些节点，空就是没有下一个结点了。
  译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。

  <pre>
    https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg
    输入：graph = [[1,2],[3],[3],[]]
    输出：[[0,1,3],[0,2,3]]
    解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
  </pre>

  <pre>
    https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg
    输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
    输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
  </pre>

  <pre>
    输入：graph = [[1],[]]
    输出：[[0,1]]
  </pre>

  <pre>
    输入：graph = [[1,2,3],[2],[3],[]]
    输出：[[0,1,2,3],[0,2,3],[0,3]]
  </pre>

  <pre>
    输入：graph = [[1,3],[2],[3],[]]
    输出：[[0,1,2,3],[0,3]]
  </pre>

  PS:
    n == graph.length
    2 <= n <= 15
    0 <= graph[i][j] < n
    graph[i][j] != i（即，不存在自环）
    graph[i] 中的所有元素 互不相同
    保证输入为 有向无环图（DAG）

  解题： 时间复杂度O(n) n为数组长度
	循环整个数组，比较当前元素和pos-1位置的元素的值，如果不相等，就将pos的值替换为当前的元素的值
*/
package questionBank

func allPathsSourceTarget(graph [][]int) [][]int {
	list := getNext(0, graph)
	res := make([][]int, 0)
	for _, ints := range list {
		if ints[len(ints)-1] != len(graph)-1 {
			continue
		}

		res = append(res, ints)
	}

	return res
}

func getNext(index int, graph [][]int) [][]int {
	if len(graph[index]) == 0 || index == len(graph)-1 {
		return [][]int{[]int{index}}
	}

	res := make([][]int, 0)
	for _, i2 := range graph[index] {
		nextList := getNext(i2, graph)
		for _, ints := range nextList {
			item := []int{index}
			item = append(item, ints...)
			res = append(res, item)
		}
	}

	return res
}
