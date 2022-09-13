package sort

// MergeSort 归并排序
func MergeSort(a []int, l int, r int) {
	if l < r {
		mid := (l + r) >> 1
		MergeSort(a, l, mid)
		MergeSort(a, mid+1, r)
		tmp := make([]int, 0)
		i, j := l, mid+1
		for i <= mid && j <= r {
			if a[i] <= a[j] {
				tmp = append(tmp, a[i])
				i++
			} else {
				tmp = append(tmp, a[j])
				j++
			}
		}
		for i <= mid {
			tmp = append(tmp, a[i])
			i++
		}
		for j <= r {
			tmp = append(tmp, a[j])
			j++
		}
		for i := 0; i < len(tmp); i++ {
			a[l+i] = tmp[i]
		}
	}
}

// ReserveNums 求逆序对个数
func ReserveNums(a []int, l int, r int) int {
	if l < r {
		mid := (l + r) >> 1
		x := ReserveNums(a, l, mid)
		y := ReserveNums(a, mid+1, r)
		tmp := make([]int, 0)
		i, j, nums := l, mid+1, 0
		for i <= mid && j <= r {
			if a[i] <= a[j] {
				tmp = append(tmp, a[i])
				i++
			} else {
				tmp = append(tmp, a[j])
				nums += mid - i + 1
				j++
			}
		}
		for i <= mid {
			tmp = append(tmp, a[i])
			i++
		}
		for j <= r {
			tmp = append(tmp, a[j])
			j++
		}
		for i := 0; i < len(tmp); i++ {
			a[l+i] = tmp[i]
		}
		return x + y + nums
	}
	return 0
}
