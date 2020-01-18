package questions

/**
 * The read4 API is already defined for you.
 *
 *     read4 := func(buf []byte) int
 *
 * // Below is an example of how the read4 API can be called.
 * file := File("abcdefghijk") // File is "abcdefghijk", initially file pointer (fp) points to 'a'
 * buf := make([]byte, 4) // Create buffer with enough space to store characters
 * read4(buf) // read4 returns 4. Now buf = ['a','b','C','d'], fp points to 'e'
 * read4(buf) // read4 returns 4. Now buf = ['e','f','g','h'], fp points to 'i'
 * read4(buf) // read4 returns 3. Now buf = ['i','j','k',...], fp points to End of file
 */
var solution = func(read4 func([]byte) int) func([]byte, int) int {
	pos := 0
	size := 0
	buffer := make([]byte, 4)
	return func(buf []byte, n int) int {
		count := 0
		for count < n {
			if size == pos {
				pos = 0
				size = read4(buffer)
				if size == 0 {
					return count
				}
			}
			buf[count] = buffer[pos]
			count++
			pos++

		}
		return count
	}
}
