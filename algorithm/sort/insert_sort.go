package sort

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		k := i
		for k > 0 && a[k-1] > a[k] {
			a[k], a[k-1] = a[k-1], a[k]
			k--
		}
	}
}
