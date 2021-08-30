package reader

import "bufio"

// there are already two functions that give control over length of data read
// io.LimitReader() to limit max bytes of a reader. It will put reader's EOF near
// io.ReadFull() to read to a max amount of bytes.

// there are also ones for using a delimiter
// bufio.reader.ReadByte()
// bufio.reader.ReadString()

// however, a combination of both use cases needs to be implemented via
// bufio.reader.ReadSlice() and its amazing.

func ReadSliceLimit(b *bufio.Reader, maxBytes int, delim byte) ([]byte, error) {
	b = bufio.NewReaderSize(b, maxBytes)
	return b.ReadSlice(delim)
}
