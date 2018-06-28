package codefights

import "testing"

const TEST_STRING = "abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"

func benchmarkFirstNotRepeatingCharacter2(size int, firstNotRepeatingCharacter2 func(string) string, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		firstNotRepeatingCharacter2(TEST_STRING)
	}
}

func benchmarkFirstNotRepeatingCharacter(size int, firstNotRepeatingCharacter func(string) string, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		firstNotRepeatingCharacter(TEST_STRING)
	}
}

func BenchmarkFirstNotRepeatingCharacterFuncs225(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter2(25, firstNotRepeatingCharacter2, b)

}
func BenchmarkFirstNotRepeatingCharacterFuncs250(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter2(50, firstNotRepeatingCharacter2, b)

}
func BenchmarkFirstNotRepeatingCharacterFuncs2100(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter2(100, firstNotRepeatingCharacter2, b)

}

func BenchmarkFirstNotRepeatingCharacterFuncs21000(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter2(1000, firstNotRepeatingCharacter2, b)

}
func BenchmarkFirstNotRepeatingCharacterFuncs25(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter(25, firstNotRepeatingCharacter, b)

}
func BenchmarkFirstNotRepeatingCharacterFuncs50(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter(50, firstNotRepeatingCharacter, b)

}
func BenchmarkFirstNotRepeatingCharacterFuncs100(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter(100, firstNotRepeatingCharacter, b)

}

func BenchmarkFirstNotRepeatingCharacterFuncs1000(b *testing.B) {
	benchmarkFirstNotRepeatingCharacter(1000, firstNotRepeatingCharacter, b)

}
