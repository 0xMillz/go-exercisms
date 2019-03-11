// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// This package does not need documentation, it's just an exercism!
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a time.Time and returns the time.Time + a gigasecond
func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Second * 1000000000

	return t.Add(gigasecond)
}
