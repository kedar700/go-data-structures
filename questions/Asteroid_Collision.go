package questions

func AsteroidCollision(asteroids []int) []int {
	fList := make([]int, 0)
	for _, asteroidVal := range asteroids {
		keep := true
		for len(fList) > 0 && asteroidVal < 0 && fList[len(fList)-1] > 0 { // if the asteroid value is negetive and the last stored element is positive
			if fList[len(fList)-1] < -asteroidVal {
				fList = fList[:len(fList)-1]
				continue
			} else if fList[len(fList)-1] == -asteroidVal {
				fList = fList[:len(fList)-1]
			}
			keep = false
			break
		}
		if keep {
			fList = append(fList, asteroidVal)
		}
	}
	return fList
}
