package sorting

func Insertion(n []int) []int {
	var i = 1
	for i < len(n) {
		var j = 1
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]
			j--
		}
		i++
	}
	return n
}
