package RoundFunction

func pReplace(input [32]int)(output [32]int){
	table := [56]int {
		16, 7,  20, 21,
		29, 12, 28, 17,
		1,  15, 23, 26,
		5,  18, 31, 10,
		2,  8,  24, 14,
		32, 27, 3,  9,
		19, 13, 30, 6,
		22, 11, 4,  25}
	for i:= 0; i< 8; i++ {
		for j:= 0; j < 4; j++ {
			index := i*4 + j
			output[index] = input[table[index] - 1]
		}
	}
	return output
}