package buffstreams

import (
	"encoding/binary"
	"math"
)

func byteArrayToUInt32(bytes []byte) (result uint64, bytesRead int) {
	return binary.Uvarint(bytes)
}

func intToByteArray(value uint64, bufferSize int) []byte {
	toWriteLen := make([]byte, bufferSize)
	bytesused := binary.PutUvarint(toWriteLen, value)
	return toWriteLen[:bytesused]
}

// Formula for taking size in bytes and calculating # of bits to express that size
// http://www.exploringbinary.com/number-of-bits-in-a-decimal-integer/
func messageSizeToBitLength(messageSize int) int {
	bytes := float64(messageSize)
	header := math.Ceil(math.Floor(math.Log2(bytes)+1)/8.0) + 1
	return int(header)
}
