package boc

import (
	"math"
	"testing"
)

func TestAppend(t *testing.T) {
	s1 := NewBitString(8 * 10)
	s1.WriteUint(1, 80)
	s1.ReadBit()
	s2 := NewBitString(8 * 10)
	s2.WriteUint(1, 80)
	s2.ReadBit()
	s1.Append(s2)
}

func TestMinBits(t *testing.T) {
	for i := 0; i < 1000500; i++ {
		if minBitsRequired(uint64(i)) != int(math.Ceil(math.Log2(float64(i+1)))) {
			t.Fatal(i)
		}
	}
}

func BenchmarkMinbits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = minBitsRequired(uint64(i))
	}
}

func BenchmarkOldMinbits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = int(math.Ceil(math.Log2(float64(i + 1))))
	}
}

func TestBitString_WriteBit(t *testing.T) {
	bs1 := NewBitString(8)
	for i := 0; i <= 7; i++ {
		if err := bs1.WriteBit(true); err != nil {
			t.Errorf("WriteBit() failed: %v", err)
		}
	}

	tests := []struct {
		name      string
		bitstring BitString
		value     bool
		wantErr   string
	}{
		{
			name:      "overflow",
			bitstring: bs1,
			value:     true,
			wantErr:   "BitString overflow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.bitstring.WriteBit(tt.value)
			if len(tt.wantErr) > 0 {
				if err == nil {
					t.Errorf("WriteBit() must return an error")
				}
				if err.Error() != tt.wantErr {
					t.Errorf("WriteBit() error = %v, want = %v", err.Error(), tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Errorf("WriteBit() error = %v", err)
			}
		})
	}
}
