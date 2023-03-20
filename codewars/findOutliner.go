package main

func even(number int) bool {
	return number%2 == 0
}

func findOutlier(integers []int) int {
	var e, o, ie, io, i int
	for _, v := range integers {
		if even(v) {
			e++
			ie = v
		} else {
			o++
			io = v
		}
	}

	if e > o {
		i = io
	} else {
		i = ie
	}
	return i
}

// Best Practice
// func FindOutlier(integers []int) int {
//   var odd, even *int
//   for i, integer := range integers {
//     if even != nil && odd != nil {
//       if integer % 2 == 0 {
//         return *odd
//       }
//       return *even
//     }
//     if integer % 2 == 0 {
//       even = &integers[i]
//     } else {
//       odd = &integers[i]
//     }
//   }
//   return integers[len(integers)-1]
// }
