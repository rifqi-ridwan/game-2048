package box

func (b Box) MoveLeft() bool {
	var isValid bool
	for i := 0; i < len(b); i++ {
		b[i], isValid = moveRow(b[i])
	}
	return isValid
}

func (b Box) MoveRight() bool {
	b.Reverse()
	isValid := b.MoveLeft()
	b.Reverse()
	return isValid
}

func (b Box) MoveUp() bool {
	b.Rotate()
	b.Rotate()
	b.Rotate()
	isValid := b.MoveLeft()
	b.Rotate()
	return isValid
}

func (b Box) MoveDown() bool {
	b.Rotate()
	isValid := b.MoveLeft()
	b.Rotate()
	b.Rotate()
	b.Rotate()
	return isValid
}

func (b Box) Reverse() {
	for i := 0; i < len(b); i++ {
		b[i] = reverseRow(b[i])
	}
}

// Rotate will rotate matrix array to clockwise
func (b Box) Rotate() {
	results := make([][]int, Grid)
	size := len(b)
	for i := 0; i < size; i++ {
		results = append(results, make([]int, size))
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			results[i][j] = results[size-j-1][i]
		}
	}
	b = results
}

func reverseRow(arr []int) []int {
	result := []int{}
	for i := len(arr) - 1; i <= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

// move element from [2 0 2 0] to [2 2 0 0]
func moveRow(arr []int) ([]int, bool) {
	index := 0
	isValid := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr[index] = arr[i]
			index++
			isValid = true
		}
	}

	for i := index; i < len(arr); i++ {
		arr[i] = 0
	}
	return mergeElement(arr, isValid)
}

// merge [4 4 2 2] [2 2 0 0]
func mergeElement(arr []int, isValid bool) ([]int, bool) {
	index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[index] {
			arr[index] += arr[i]
			arr[i] = 0
			isValid = true
		} else {
			index++
			arr[index] = arr[i]
		}
	}
	return arr, isValid
}
