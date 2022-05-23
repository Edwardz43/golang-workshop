package hw2

import "errors"

func Add(i, j int) (int, error) {
	if i < 0 || j < 0 {
		return 0, errors.New("negative number")
	}
	return i + j, nil
}

func Sum(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("negative number")
	}
	sum := 0
	for i > 0 {
		sum += i
		i--
	}
	return sum, nil
}
