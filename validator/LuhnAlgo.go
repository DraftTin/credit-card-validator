package validator

func Luhn(number string) bool {
	n := len(number)
	sum := 0
	parity := (n - 2) % 2
	for i := n - 2; i >= 0; i-- {
		if i%2 != parity {
			sum += int(number[i]) - '0'
		} else {
			tmp := 2 * (int(number[i]) - '0')
			sum += tmp/10 + tmp%10
		}
	}
	return (10 - (sum % 10)) == int(number[n-1])-'0'
}
