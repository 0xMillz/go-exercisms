package letter

func ConcurrentFrequency(ss []string) FreqMap {
	c := make(chan FreqMap)

	for _, s := range ss {
		go func(str string) {
			c <- Frequency(str)
		}(s)
	}

	m := FreqMap{}

	for range ss {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
