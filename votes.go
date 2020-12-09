package main

import (
	"fmt"
	"log"
	"../datafile"	// 경로 설정 매우 중요하다...
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	// var counts map[string]int
	// var 맵이름 map[키 타입]값 타입

	counts := make(map[string]int)
	// 키 타입은 이름이 되고 값 타입은 득표수가 된다.
	// 맵에서는 키를 통해 값을 읽어온다.

	for _, line := range lines {
		counts[line]++
	}
	// 앞에서 데이터파일을 읽어온 lines 변수에서 처리하는데,
	// lines를 돌면서 맵의 이름을 키로 정하고, 값은 ++ 시키는 것이다.

	for name, count := range counts {
		fmt.Printf("votes for %s : %d\n", name, count)
	}
}
