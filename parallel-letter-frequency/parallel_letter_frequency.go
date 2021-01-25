package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts letters in multiple strings concurrently
func ConcurrentFrequency(strlist []string) FreqMap {
	queue := make(chan FreqMap)
	// Submit a job for each string in the slice
	for _, s := range strlist {
		go func(s string) {
			queue <- Frequency(s)
		}(s)
	}

	// range over strlist to make sure results from all jobs are collected
	result := FreqMap{}
	for range strlist {
		for k, v := range <-queue { // merge all results on the queue
			result[k] += v
		}
	}

	return result
}
