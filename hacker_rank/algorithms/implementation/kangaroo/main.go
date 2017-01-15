package main

import "fmt"

const dataSize = 4

type kangaroo struct {
	startingLocation int
	hopVelocity      int
}

func main() {
	roos := readInKangarooData()
	determineIfHopsMeet(roos)
}

func readInKangarooData() []kangaroo {
	var data = make([]int, dataSize)
	for i := 0; i < dataSize; i++ {
		fmt.Scanf("%d", &data[i])
	}
	return fillinKanagarooData(data)
}

func fillinKanagarooData(data []int) []kangaroo {
	var roos = make([]kangaroo, len(data)/2)

	var rooCount int
	for i := 0; i < dataSize; {
		roos[rooCount].startingLocation = data[i]
		roos[rooCount].hopVelocity = data[i+1]
		rooCount++
		i = i + 2
	}

	return roos
}

func determineIfHopsMeet(roos []kangaroo) {
	if roos[1].hopVelocity < roos[0].hopVelocity && (roos[1].startingLocation-roos[0].startingLocation)%(roos[1].hopVelocity-roos[0].hopVelocity) == 0 {
		success()
	} else {
		failure()
	}
}

func success() {
	fmt.Println("YES")
}

func failure() {
	fmt.Println("NO")
}
