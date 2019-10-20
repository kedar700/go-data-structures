package questions

/*
public int[] asteroidCollision(int[] asteroids) {
        LinkedList<Integer> s = new LinkedList<>();
        for (int i : a) {
            if (i > 0)
                s.add(i);
            else {
                while (!s.isEmpty() && s.getLast() > 0 && s.getLast() < -i)
                    s.pollLast();
                if (!s.isEmpty() && s.getLast() == -i)
                    s.pollLast();
                else if (s.isEmpty() || s.getLast() < 0)
                    s.add(i);
            }
        }
        return s.stream().mapToInt(i->i).toArray();
    }
  answers := make([]int, 0)

    for _, v := range asteroids{
        keep := true
        for len(answers) > 0 && v < 0 && answers[len(answers)-1] > 0{
            if answers[len(answers)-1] < -v{
                answers = answers[:len(answers)-1]
                continue
            }else if answers[len(answers)-1] == -v{
                answers = answers[:len(answers)-1]
            }
            keep = false
            break
        }

        if keep{
            answers = append(answers, v)
        }
    }
    return answers
*/

func AsteroidCollision(asteroids []int) []int {
	fList := make([]int, 0)
	for _, asteroidVal := range asteroids {
		keep := true
		for len(fList) > 0 && asteroidVal < 0 && fList[len(fList)-1] > 0 { // if the asteroid value is negetive and the last element is negetive
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
