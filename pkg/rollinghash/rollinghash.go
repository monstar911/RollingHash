package rollinghash

import (
	"crypto/sha256"
	"hash"
	"io"
	"os"
)

// This algorithm first reads the originalFilename and updatedFilename versions of the file
// in chunks,computing the hash of each chunk using the SHA256 hash function.
// It then compares the hashes of the chunks from the originalFilename and updatedFilename files

type chunk struct {
	hash  [sha256.Size]byte
	bytes []byte
}

type RollingHash struct {
	hasher    hash.Hash
	chunkSize int
}

// NewRollingHash returns a new RollingHash struct
func NewRollingHash(chunkSize int) *RollingHash {
	return &RollingHash{hasher: sha256.New(), chunkSize: chunkSize}
}

type RDiff interface {
	ComputeHashes(file *os.File) ([]chunk, error)
	ComputeDelta(original, updated *os.File) ([]byte, error)
}

// ComputeHashes reads a file in chunks and computes the hash of each chunk using the SHA1 hash function
func (r *RollingHash) ComputeHashes(file *os.File) ([]chunk, error) {
	chunks := make([]chunk, 0)
	buf := make([]byte, r.chunkSize)

	for {
		n, err := file.Read(buf)
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		sum := sha256.Sum256(buf[:n])
		chunks = append(chunks, chunk{hash: sum, bytes: buf[:n]})
	}

	return chunks, nil
}

// ComputeDelta generates a description of the differences between the original and updated versions of a file
func (r *RollingHash) ComputeDelta(original, updated *os.File) ([]byte, error) {
	originalChunks, err := r.ComputeHashes(original)
	if err != nil {
		return nil, err
	}
	updatedChunks, err := r.ComputeHashes(updated)
	if err != nil {
		return nil, err
	}

	// keep track of the current position in each list of chunks
	originalPos := 0
	updatedPos := 0
	var delta []byte

	for updatedPos < len(updatedChunks) {
		// if the hashes of the current chunk from the original and updated files match,
		// it means the chunk can be reused and we can move to the next chunk in both lists
		if originalPos < len(originalChunks) && originalChunks[originalPos].hash == updatedChunks[updatedPos].hash {
			originalPos++
			updatedPos++
			continue
		}

		// otherwise, the chunk has been added or modified, so we need to add it to the delta
		delta = append(delta, updatedChunks[updatedPos].bytes...)
		updatedPos++
	}

	return delta, nil
}
