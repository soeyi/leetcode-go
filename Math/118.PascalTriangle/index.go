package Math

//Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 5
//Output:
//[
//[1],
//[1, 1],
//[1, 2, 1],
//[1,3, 3, 1],
//[1,4, 6, 4, 1]
//]

func generate(numRows int) [][]int {
	var res [][]int
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		m := []int{0}
		m = append(m, res[i-1]...)
		for j := 0; j < len(m)-1; j++ {
			m[j] = m[j] + m[j+1]
		}
		res = append(res, m)
	}
	return res
}
