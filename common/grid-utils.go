package common

func IsPosOutside[T any](pos [2]int, grid *[][]T) bool {
	return pos[0] < 0 || pos[0] >= len(*grid) || pos[1] < 0 || pos[1] >= len((*grid)[0])
}
