package main

import "fmt"

func print_basic_slice() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])
	fmt.Println(sl5[1 : len(sl5)-1])
}

func print_slice_append_make_len_cap() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))

	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
}

func print_slice_copy() {
	// sl := []int{100, 200}
	// sl2 := sl

	// sl2[0] = 1000

	// fmt.Println(sl)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)

	fmt.Println(sl2)

	//n => copyに成功した数
	n := copy(sl2, sl)

	fmt.Println(n, sl2)
}

func print_slice_for() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func print_map_basic() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "hoge"
	m4[2] = "fuga"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	s,ok := m4[4]
	if !ok {
		fmt.Println("error")
	} 
	fmt.Println(s, ok)

	m4[3] = "piyo"
	fmt.Println(m4)
	
	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

}

func print_map_for() {
	m := map[string]int{
		"apple": 100,
		"banana": 200,
	}

	for _, v := range m {
		fmt.Println(v)
	}
}

func main() {
	// sl := []int{1,2,3,4,5,6}
	// fmt.Println(Sum(sl...))

	// print_map_basic()

	print_map_for()
}
