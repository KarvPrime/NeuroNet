package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestNeuroNet_main(t *testing.T) {
	main()
}
