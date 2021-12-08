package main

func joinSegNumSlices(a, b []segNum) []segNum {
	newSlice := []segNum{}
	all := make([]segNum, len(a)+len(b))
	copy(all, a)
	copy(all[len(a):], b)

outer:
	for _, val := range all {
		for _, v := range newSlice {
			if v == val {
				continue outer
			}
		}

		newSlice = append(newSlice, val)
	}

	return newSlice
}

func maskSegments(a, b []segNum) []segNum {
	newSlice := []segNum{}
	for _, v := range a {
		for _, vb := range b {
			if v == vb {
				newSlice = append(newSlice, v)
			}
		}
	}

	return newSlice
}
