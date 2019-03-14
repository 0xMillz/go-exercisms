package letter

func ConcurrentFrequency(ss []string) FreqMap {
	m := FreqMap{} // master map

	c := make(chan FreqMap)

	for _, s := range ss {
		go func() {
			c <- Frequency(s)
		}()
	}

	for range ss {
		msg := <- c
		for k, v := range msg {
			m[k] = m[k] + v
		}
	}

	return m
}
