package collatzconjecture

import "fmt"


func CollatzConjecture(n int) (int, error) {
	steps := 0;
	if n <= 0 {
		return 0, fmt.Errorf("n less than 0")
	}
	if n == 1 {
		return 0, nil
	}
	for n > 1 {

		steps += 1
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = n * 3
			n += 1
		}

	}
	return steps, nil
}
