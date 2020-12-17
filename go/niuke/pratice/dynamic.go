package main

/**
 *
 * @param s string字符串
 * @return int整型
 */
func longestValidParentheses(s string) int {
	if len(s) < 1 {
		return 0
	}
	// write code here
	max := 0
	last := -1
	stack := []int{}
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				last = i
			} else {
				stack = stack[0 : len(stack)-1]
				if len(stack) == 0 {
					max = bigger(max, i-last)
				} else {
					max = bigger(max, i-stack[len(stack)-1])
				}
			}
		}
	}
	return max
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
