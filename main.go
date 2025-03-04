package main

import "fmt"

// auth "todo/Auth"
// "todo/pck/dbconnection"
// "todo/pck/handlers"
// "todo/pck/middleware"
// "todo/pck/repository"
// "todo/pck/server"
// "todo/pck/service"

// "github.com/rs/zerolog/log"

func main() {
	// log.Info().Msg("connecting to database")
	// db, err := dbconnection.Dbconnect()
	// log.Info().Msg("CONNECTED TO DATABASE")

	// if err != nil {
	// 	return
	// }
	// repo, err := repository.Newrepository(db)
	// if err != nil {
	// 	return
	// }
	// auth, err := auth.NewAuth(auth.SingNature)
	// if err != nil {
	// 	return
	// }

	// mid, err := middleware.NewMiddleware(auth)
	// if err != nil {
	// 	return
	// }
	// service, err := service.NewService(repo, auth)
	// if err != nil {
	// 	return
	// }
	// handlers, err := handlers.NewHandler(service)
	// if err != nil {
	// 	return
	// }
	// server.StartingServer(handlers, mid)

	arr := []int{1, 2, 9, 8, 2, 1, 9, 9}
	target := 3
	first, last := LinaerSearch(arr, target)
	fmt.Println(first, last)
	v := findUnique1(arr)
	fmt.Println(v)
	//first := firstbinarysearch(arr, target)
	//last := lastbinarysearch(arr, target)
	//fmt.Println(first, last)
}

var first, last int

func LinaerSearch(arr []int, t int) (int, int) {
	// //count := 0
	// min := arr[0]
	// max := arr[0]
	// for _, v := range arr {
	// 	if v >= max {
	// 		max = v
	// 	} else if v <= min {
	// 		min = v
	// 	}
	// }
	// return min, max
	first, last = -1, -1
	for i, v := range arr {
		if v == t {
			if first == -1 {
				first = i
			}
			last = i
		}
	}
	return first, last
}

func findUnique(arr []int) int {
	result := 0
	for _, num := range arr {
		result ^= num // XOR cancels out duplicates
	}
	return result
}

func firstbinarysearch(arr []int, target int) int {
	reslut := -1
	left, rigth := 0, len(arr)-1
	var mid int

	for left <= rigth {
		mid = left + (rigth-left)/2
		if arr[mid] == target {
			reslut = mid
			rigth = mid - 1

		} else if arr[mid] < target {
			left = mid + 1
		} else {
			rigth = mid - 1
		}
	}

	return reslut
}

func lastbinarysearch(arr []int, target int) int {
	reslut := -1
	left, rigth := 0, len(arr)-1
	var mid int

	for left <= rigth {
		mid = left + (rigth-left)/2
		if arr[mid] == target {
			reslut = mid
			left = mid + 1

		} else if arr[mid] < target {
			left = mid + 1
		} else {
			rigth = mid - 1
		}
	}

	return reslut
}

func findUnique1(arr []int) int {
	freq := make(map[int]int)

	// Count occurrences
	for _, num := range arr {
		freq[num]++
	}

	// Find the number that appears only once
	for key, val := range freq {
		if val == 1 {
			return key
		}
	}
	return -1 // No unique element found
}
