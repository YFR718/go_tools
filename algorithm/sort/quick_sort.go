package sort

func QuickSort(a []int, l int, r int) {
	if l < r {
		i, j, k := l, r, a[l+r>>1]
		for i < j {
			for a[i] < k {
				i++
			}
			for a[j] < k {
				j--
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		QuickSort(a, l, l+r>>1)
		QuickSort(a, l+r>>1+1, r)
	}
}
