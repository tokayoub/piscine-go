package piscine

func StrRev(s string) string {
	S := []rune(s)
	len := 0
	for range s {
		len++
	}
	for i, n := 0, len-1; i < n; i, n = i+1, n-1 {
		x := S[i]
		S[i] = S[n]
		S[n] = x
	}
	return (string(S))
}
