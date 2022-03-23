package sorting

func Bubble(n []int) []int {

	var isDone = false
	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}
	return n
}
