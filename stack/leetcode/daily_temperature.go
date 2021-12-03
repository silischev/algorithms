package main

func dailyTemperatures(temperatures []int) []int {
	values := make([]int, len(temperatures))
	tmp := []int{}

	for day := 0; day <= len(temperatures)-1; day++ {
		if day == len(temperatures)-1 {
			values[day] = 0
			break
		}

		if temperatures[day] < temperatures[day+1] {
			values[day] = 1

			if len(tmp) == 0 {
				continue
			}

			for i := len(tmp) - 1; i >= 0; i-- {
				if temperatures[tmp[i]] < temperatures[day+1] {
					values[tmp[i]] = day + 1 - tmp[i]
					tmp = tmp[:i]
				}
			}

			continue
		}

		tmp = append(tmp, day)
	}

	return values
}
