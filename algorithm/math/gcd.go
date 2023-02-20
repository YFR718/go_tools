package math

func Gcd(p int, q int) int {
	if q == 0 {
		return p
	} else {
		return Gcd(q, p%q)
	}
}
