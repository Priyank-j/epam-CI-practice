package generateParenthesis

func generateParenthesis(n int) []string {

	var result []string
	findAllCombinations("(", 1, 0, n, &result)
	return result
}

func findAllCombinations(current string, open, close, n int, result *[]string) {
	if len(current) == n*2 {
		*result = append(*result, current)
		return
	}

	if open < n {
		findAllCombinations(current+"(", open+1, close, n, result)
	}
	if close < open {
		findAllCombinations(current+")", open, close+1, n, result)
	}
}
