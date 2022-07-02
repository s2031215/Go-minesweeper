package main

import (
	"math/rand"
	"time"
)

func bombinit() [260]int {
	tempArr := [260]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 40; i++ {
		j := rand.Intn(260)
		for tempArr[j] == -1 {
			j = rand.Intn(260)
		}
		tempArr[j] = -1
	}

	for i := 0; i < 260; i++ {
		if tempArr[i] == 0 {
			tempArr[i] = len(checkaround(i, tempArr, -1))
		}
	}

	return tempArr
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

func getaround(i int, bombarr [260]int) []int {
	findarr := []int{}
	//check upper
	if i > 19 && bombarr[i-20] != -1 {
		findarr = append(findarr, i-20)
	}
	//check upper left
	if i > 19 && i%20 != 0 && bombarr[i-21] != -1 {
		findarr = append(findarr, i-21)
	}
	//check upper right
	if i > 19 && i%20 != 19 && bombarr[i-19] != -1 {
		findarr = append(findarr, i-19)
	}
	//check down
	if i < 240 && bombarr[i+20] != -1 {
		findarr = append(findarr, i+20)
	}
	//check down left
	if i < 240 && i%20 != 0 && bombarr[i+19] != -1 {
		findarr = append(findarr, i+19)
	}
	//check down right
	if i < 240 && i%20 != 19 && bombarr[i+21] != -1 {
		findarr = append(findarr, i+21)
	}
	//check left
	if i > 0 && i%20 != 0 && bombarr[i-1] != -1 {
		findarr = append(findarr, i-1)
	}
	//check right
	if i < 259 && i%20 != 19 && bombarr[i+1] != -1 {
		findarr = append(findarr, i+1)
	}

	return findarr
}

func checkaround(i int, bombarr [260]int, key int) []int {
	findarr := []int{}
	//check upper
	if i > 19 && bombarr[i-20] == key {
		findarr = append(findarr, i-20)
	}
	//check upper left
	if i > 19 && i%20 != 0 && bombarr[i-21] == key {
		findarr = append(findarr, i-21)
	}
	//check upper right
	if i > 19 && i%20 != 19 && bombarr[i-19] == key {
		findarr = append(findarr, i-19)
	}
	//check down
	if i < 240 && bombarr[i+20] == key {
		findarr = append(findarr, i+20)
	}
	//check down left
	if i < 240 && i%20 != 0 && bombarr[i+19] == key {
		findarr = append(findarr, i+19)
	}
	//check down right
	if i < 240 && i%20 != 19 && bombarr[i+21] == key {
		findarr = append(findarr, i+21)
	}
	//check left
	if i > 0 && i%20 != 0 && bombarr[i-1] == key {
		findarr = append(findarr, i-1)
	}
	//check right
	if i < 259 && i%20 != 19 && bombarr[i+1] == key {
		findarr = append(findarr, i+1)
	}

	return findarr
}

func flagcheck(active int, bombarr [260]int, blockstatus [260]int) []int {
	tmpcleanarr := []int{active}
	flag := checkaround(active, blockstatus, 2)
	bomb := checkaround(active, bombarr, -1)
	if len(flag) == len(bomb) {

		for _, v := range bomb {
			if !contains(flag, v) {
				die = true
			}
		}

		if !die {
			for _, v := range checkaround(active, blockstatus, 0) {
				if blockstatus[v] == 0 {
					tmpcleanarr = append(tmpcleanarr, findempty(v, bombarr, []int{})...)
				}
			}
			tmpcleanarr = append(tmpcleanarr, getaround(active, bombarr)...)
		}

	}
	return tmpcleanarr
}

func findempty(active int, bombarr [260]int, cleanarr []int) []int {
	//fmt.Printf("start %2v\n", active)
	i := active
	tempArr := bombarr
	//block will show Array
	tmpcleanarr := append(cleanarr, i)
	//check upper

	if i > 19 && tempArr[i-20] == 0 && !contains(tmpcleanarr, i-20) {
		//tmpcleanarr = append(tmpcleanarr, i-20)
		tmpcleanarr = findempty(i-20, bombarr, tmpcleanarr)
	}
	//check down
	if i < 240 && tempArr[i+20] == 0 && !contains(tmpcleanarr, i+20) {
		//tmpcleanarr = append(tmpcleanarr, i+20)
		tmpcleanarr = findempty(i+20, bombarr, tmpcleanarr)
	}
	//check left
	if i > 0 && tempArr[i-1] == 0 && i%20 != 0 && !contains(tmpcleanarr, i-1) {
		//tmpcleanarr = append(tmpcleanarr, i-1)
		tmpcleanarr = findempty(i-1, bombarr, tmpcleanarr)
	}
	//check right
	if i < 259 && tempArr[i+1] == 0 && i%20 != 19 && !contains(tmpcleanarr, i+1) {
		//tmpcleanarr = append(cleanarr, i+1)
		tmpcleanarr = findempty(i+1, bombarr, tmpcleanarr)
	}
	//check upper left
	if i > 19 && i%20 != 0 && bombarr[i-21] == 0 && !contains(tmpcleanarr, i-21) {
		tmpcleanarr = findempty(i-21, bombarr, tmpcleanarr)
	}
	//check upper right
	if i > 19 && i%20 != 19 && bombarr[i-19] == 0 && !contains(tmpcleanarr, i-19) {
		tmpcleanarr = findempty(i-19, bombarr, tmpcleanarr)
	}
	//check down left
	if i < 240 && i%20 != 0 && bombarr[i+19] == 0 && !contains(tmpcleanarr, i+19) {
		tmpcleanarr = findempty(i+19, bombarr, tmpcleanarr)
	}
	//check down right
	if i < 240 && i%20 != 19 && bombarr[i+21] == 0 && !contains(tmpcleanarr, i+21) {
		tmpcleanarr = findempty(i+21, bombarr, tmpcleanarr)
	}
	if bombarr[i] == 0 {
		//check upper
		if i > 19 && bombarr[i-20] > 0 && !contains(tmpcleanarr, i-20) {
			tmpcleanarr = append(tmpcleanarr, i-20)
		}
		//check upper left
		if i > 19 && i%20 != 0 && bombarr[i-21] > 0 && !contains(tmpcleanarr, i-21) {
			tmpcleanarr = append(tmpcleanarr, i-21)
		}
		//check upper right
		if i > 19 && i%20 != 19 && bombarr[i-19] > 0 && !contains(tmpcleanarr, i-19) {
			tmpcleanarr = append(tmpcleanarr, i-19)
		}
		//check down
		if i < 240 && bombarr[i+20] > 0 && !contains(tmpcleanarr, i+20) {
			tmpcleanarr = append(tmpcleanarr, i+20)
		}
		//check down left
		if i < 240 && i%20 != 0 && bombarr[i+19] > 0 && !contains(tmpcleanarr, i+19) {
			tmpcleanarr = append(tmpcleanarr, i+19)
		}
		//check down right
		if i < 240 && i%20 != 19 && bombarr[i+21] > 0 && !contains(tmpcleanarr, i+21) {
			tmpcleanarr = append(tmpcleanarr, i+21)
		}
		//check left
		if i > 0 && i%20 != 0 && bombarr[i-1] > 0 && !contains(tmpcleanarr, i-1) {
			tmpcleanarr = append(tmpcleanarr, i-1)
		}
		//check right
		if i < 259 && i%20 != 19 && bombarr[i+1] > 0 && !contains(tmpcleanarr, i+1) {
			tmpcleanarr = append(tmpcleanarr, i+1)
		}
	}

	//fmt.Printf("end %2v\n", i)
	return tmpcleanarr
}
