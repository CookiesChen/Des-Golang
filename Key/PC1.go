package Key

/* PC-1置换 */
func pc1(k [64]int) (output [56]int){

	table := [56]int {
		57, 49, 41, 33, 25, 17, 9,
		1,  58, 50, 42, 34, 26, 18,
		10, 2,  59, 51, 43, 35, 27,
		19, 11, 3,  60, 52, 44, 36,
		63, 55, 47, 39, 31, 23, 15,
		7,  62, 54, 46, 38, 30, 22,
		14, 6,  61, 53, 45, 37, 29,
		21, 13, 5,  28, 20, 12, 4 }

	for i:= 0; i< 8; i++ {
		for j:= 0; j < 7; j++ {
			index := i*7 + j
			output[index] = k[table[index] - 1]
		}
	}
	return output
}