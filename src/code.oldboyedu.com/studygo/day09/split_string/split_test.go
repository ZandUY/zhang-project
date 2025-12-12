package splitstring

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {
// 	type testCase struct {
// 		str  string
// 		sep  string
// 		want []string
// 	}
// 	testGroup := []testCase{
// 		testCase{"babxbef", "b", []string{"", "a", "x", "ef"}},
// 		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
// 		testCase{"abcef", "bc", []string{"a", "ef"}},
// 		testCase{"沙河有沙又有河", "沙", []string{"", "河有", "又有河"}}}
// 	for _, tc := range testGroup {
// 		got := Split(tc.str, tc.sep)
// 		if !reflect.DeepEqual(got, tc.want) {
// 			t.Errorf("want:%#v,got:%#v", tc.want, got)
// 		}
// 	}

// }
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case1": testCase{"babxbef", "b", []string{"", "a", "x", "ef"}},
		"case2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case3": testCase{"abcef", "bc", []string{"a", "ef"}},
		"case4": testCase{"沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want:%#v,got:%#v", tc.want, got)
			}
		})
	}

}
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
func benchmarkFibo(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fibo(n)
	}
}
func BenchmarkFibo2(b *testing.B) {
	benchmarkFibo(b, 1)
}
func BenchmarkFibo3(b *testing.B) {
	benchmarkFibo(b, 3)
}

func BenchmarkFibo10(b *testing.B) {
	benchmarkFibo(b, 10)
}
func BenchmarkFibo50(b *testing.B) {
	benchmarkFibo(b, 50)
}
func BenchmarkFibo100(b *testing.B) {
	benchmarkFibo(b, 100)
}
