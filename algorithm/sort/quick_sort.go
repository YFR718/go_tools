package sort

func QuickSort(a []int, l int, r int) {
	if l < r {
		i, j, k := l-1, r+1, a[(l+r)/2]
		for i < j {
			i++
			for a[i] < k {
				i++
			}
			j--
			for a[j] > k {
				j--
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
			}
		}
		QuickSort(a, l, j)
		QuickSort(a, j+1, r)
	}
}

func TopKQuickSort(a []int, l int, r int, k int) {
	if l < r {
		i, j, m := l, r, a[(l+r)/2]
		for i < j {
			for a[i] < m {
				i++
			}
			for a[j] > m {
				j--
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
				i++
				j--
			}
		}
		if j >= k {
			TopKQuickSort(a, l, j, k)
		} else {
			TopKQuickSort(a, j+1, r, k)
		}
	}
}
