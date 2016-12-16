package MinEditDistance

const insertCost = 1
const deleteCost = 1
const editCost = 2

func minDistance(x int, y int, z int) int {
	if x > y {
		if y > z {
			return z
		}
		return y
	}
	if x > z {
		return z
	}
	return x
}

// CalMinEditDistance 计算两个[]rune的最小编辑距离
func CalMinEditDistance(target []rune, source []rune) (distance int) {
	n := len(target)
	m := len(source)

	// create distance matrix
	var matrix [][]int
	matrix = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		matrix[i] = make([]int, m+1)
	}

	for i := 0; i <= m; i++ {
		matrix[0][i] = i
	}
	for j := 0; j <= n; j++ {
		matrix[j][0] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			var insertDistance = 0
			var substituteDistance = 0
			var deleteDistance = 0
			insertDistance = matrix[i-1][j] + insertCost
			if source[j-1] == target[i-1] {
				substituteDistance = matrix[i-1][j-1]
			} else {
				substituteDistance = matrix[i-1][j-1] + editCost
			}
			deleteDistance = matrix[i][j-1] + deleteCost
			matrix[i][j] = minDistance(insertDistance, substituteDistance, deleteDistance)
		}
	}
	distance = matrix[n][m]
	return
}
