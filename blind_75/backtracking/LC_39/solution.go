package main

func combinationSum(nums []int, target int) [][]int {
	var res [][]int
	var dfs func(int,int,[]int)

	dfs = func(i int, t int, r []int) {
		
		if t == target {
			dup := append([]int{}, r...)
			res = append(res, dup)
			return
		}

		if i >= len(nums) || t > target {
			return
		}

		// pick(i)
		r = append(r, nums[i])
		dfs(i, t+nums[i], r)

		// !pick(i)
		r = r[:len(r)-1]
		dfs(i+1, t, r)
	}

	dfs(0, 0, []int{})

	return res
}