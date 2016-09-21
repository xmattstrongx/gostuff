package main_test

import (
	"bytes"
	"fmt"
	"testing"
)

// go test -bench .

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "a" + "b" + "c"
	}
}

func BenchmarkSimpleVariable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = i + i + i
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", "a", "b", "c")
	}
}

func BenchmarkSprintfVariable(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("a")
		buffer.WriteString("b")
		buffer.WriteString("c")
	}
}

func BenchmarkByte(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("a")
		buffer.WriteString("b")
		buffer.WriteString("c")
	}
}

func BenchmarkByteVariable(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString(string(i))
		buffer.WriteString(string(i))
		buffer.WriteString(string(i))
	}
}

/* Results
BenchmarkSimple-8         	2000000000	         0.61 ns/op
BenchmarkSimpleVariable-8 	2000000000	         0.31 ns/op
BenchmarkSprintf-8        	 5000000	       407 ns/op
BenchmarkSprintfVariable-8	30000000	        44.3 ns/op
BenchmarkByte-8           	50000000	        41.0 ns/op
BenchmarkByteVariable-8   	20000000	        74.8 ns/op
*/
