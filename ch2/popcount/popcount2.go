package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	result := pc[byte(x>>(0*8))]
	for i := uint(1); i < 8; i++ {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}
