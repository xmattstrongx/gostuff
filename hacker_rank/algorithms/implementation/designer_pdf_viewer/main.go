package main

import "fmt"

const (
	size = 26
)

type PDFArea struct {
	sizeChart []int
	word      string
	maxHeight int
	area      int
}

func main() {
	pdf := PDFArea{}
	pdf.getSizeChartAndWord()
	pdf.getmaxHeight()
	pdf.calculateArea()

	fmt.Println(pdf.area)
}

func (p *PDFArea) getSizeChartAndWord() {
	p.getSizeChart()
	p.getWord()
}

func (p *PDFArea) getSizeChart() {
	p.sizeChart = make([]int, 26)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &p.sizeChart[i])
	}
}

func (p *PDFArea) getWord() {
	fmt.Scanf("%s", &p.word)
}

func (p *PDFArea) getmaxHeight() {
	max := -1
	for _, character := range p.word {
		if max < p.sizeChart[character-97] {
			max = p.sizeChart[character-97]
		}
	}
	p.maxHeight = max
}

func (p *PDFArea) calculateArea() {
	p.area = len(p.word) * 1 * p.maxHeight
}
