package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)-1; i++ {
		for n := i + 1; n < len(podium); n++ {
			if podium[i][0] > podium[n][0] {
				podium[i], podium[n] = podium[n], podium[i]
			}
		}
	}
	return podium
}
