package lc_1025

//xet A va B doi dau voi nhau, nhan thay neu ai co so 2 truoc nguoi do se thang
//muon thang thi nguoi di dau tien phai cam duoc so chan, sau do se day so le cho nguoi con lai
// nguoi cam so le se luon phai dua lai so chan, vi uoc cua le = le, le - le = chan
// co so chan chac chan se thang

func divisorGame(n int) bool {
	return n%2 == 0
}

func handle(n int, turn int) bool {
	if turn%2 == 1 {
		turn = 1
	} else if turn%2 == 0 {
		turn = 2
	}
	if n == 1 && turn == 2 {
		return false
	}

	if n == 1 && turn == 1 {
		return true
	}

	a := getDivide(n)

	check := make([]bool, len(a))

	for index, item := range a {
		check[index] = handle(n-item, turn+1)
	}

	for _, item := range check {
		if !item {
			return false
		}
	}

	return true
}

func getDivide(n int) []int {
	result := make([]int, 0)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}

	return result
}
