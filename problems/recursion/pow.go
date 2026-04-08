package recursion

const mod int64 = 1_000_000_007

func IterativemyPow(x, n int64) int64 {
	duplicate := n
	var ans int64 = 1
	for duplicate > 0 {
		if duplicate%2 == 0 {
			x *= x

			x = (x * x) % mod
			duplicate = duplicate / 2
		} else {
			ans = (ans * x) % mod
			duplicate = duplicate - 1

		}

	}

	return ans

}

func RecursivemyPow(x, n int64) int64 {
	x = x % mod
	//base condition
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		half := RecursivemyPow(x, n/2)
		return (half * half) % mod

	} else {
		return (x * RecursivemyPow(x, n-1)) % mod
	}

}
