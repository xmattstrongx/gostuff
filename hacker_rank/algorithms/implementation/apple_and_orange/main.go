package main

import (
	"fmt"
)

type houseHitCounter struct {
	appleHitCount  int
	orangeHitCount int
	*house
	*fruit
}

type house struct {
	houseStart int
	houseEnd   int
	hitCount   int
}

type fruit struct {
	apple
	orange
}

type apple struct {
	appleCount        int
	appleTreePosition int
	distances         []int
}

type orange struct {
	orangeCount        int
	orangeTreePosition int
	distances          []int
}

func main() {
	house := new(house)
	house.readInHouseData()
	fruit := new(fruit)
	fruit.readInFruitData()

	houseHitCounter := &houseHitCounter{
		house: house,
		fruit: fruit,
	}

	houseHitCounter.countHouseHits()
	fmt.Printf("%d\n", houseHitCounter.appleHitCount)
	fmt.Printf("%d\n", houseHitCounter.orangeHitCount)
}

func (h *house) readInHouseData() {
	fmt.Scanf("%d", &h.houseStart)
	fmt.Scanf("%d", &h.houseEnd)
}

func (f *fruit) readInFruitData() {
	fmt.Scanf("%d", &f.apple.appleTreePosition)
	fmt.Scanf("%d", &f.orange.orangeTreePosition)

	fmt.Scanf("%d", &f.apple.appleCount)
	fmt.Scanf("%d", &f.orange.orangeCount)

	f.apple.distances = make([]int, f.appleCount)
	for i := 0; i < f.apple.appleCount; i++ {
		fmt.Scanf("%d", &f.apple.distances[i])
	}

	f.orange.distances = make([]int, f.orangeCount)
	for i := 0; i < f.orange.orangeCount; i++ {
		fmt.Scanf("%d", &f.orange.distances[i])
	}
}

func (h *houseHitCounter) countHouseHits() {
	var appleDropLocation int
	for _, apple := range h.fruit.apple.distances {
		appleDropLocation = h.fruit.appleTreePosition + apple
		if appleDropLocation >= h.house.houseStart && appleDropLocation <= h.house.houseEnd {
			h.appleHitCount++
		}
	}

	var orangeDropLocation int
	for _, orange := range h.fruit.orange.distances {
		orangeDropLocation = h.fruit.orangeTreePosition + orange
		if orangeDropLocation >= h.house.houseStart && orangeDropLocation <= h.house.houseEnd {
			h.orangeHitCount++
		}
	}
}
