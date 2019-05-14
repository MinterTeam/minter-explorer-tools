package helpers

import (
	"math"
)

func ComputeChunksCount(dataLength int, chunkSize int) int {
	return int(math.Ceil(float64(dataLength) / float64(chunkSize)))
}

func ComputeDataRange(dataLength int, chunkSize int, step int) (int, int) {
	start := chunkSize * step
	end := chunkSize + start
	if end > dataLength {
		end = dataLength
	}

	return start, end
}
