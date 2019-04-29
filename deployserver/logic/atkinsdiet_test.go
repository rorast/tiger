package logic

import "testing"

func TestSpinNotFails(t *testing.T) {
	m := New()
	_, _, err := m.Spin(88, 20)
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkSpin(b *testing.B) {
	m := New()
	for n := 0 ; n < b.N ; n++ {
		m.Spin(1, 20)
	}
}

//func Test_mapSymbolsToStops(t *testing.T) {
//	type args struct {
//		stops [3][5]int8
//	}
//	tests := []struct {
//		name string
//		args args
//		want [3][5]symbol
//	}{
//		"all zeroes",
//		args{stops: [3][5]int8{
//			{0,0,0,0,0},
//			{0,0,0,0,0},
//			{0,0,0,0,0},
//		}},
//	},
//}