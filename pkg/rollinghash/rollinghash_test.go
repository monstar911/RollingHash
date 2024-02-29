package rollinghash

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestComputeDelta(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name              string
		original, updated *os.File
		expectedDelta     []byte
	}{
		{
			name:          "Test case 1: Identical files",
			original:      createTestFile([]byte("Hello, world!")),
			updated:       createTestFile([]byte("Hello, world!")),
			expectedDelta: []byte{},
		},
		{
			name:          "Test case 2: Added content",
			original:      createTestFile([]byte("Hello, world!")),
			updated:       createTestFile([]byte("Hello, world! How are you?")),
			expectedDelta: []byte("Hello, world! How are you?"),
		},
		{
			name:          "Test case 3: Deleted content",
			original:      createTestFile([]byte("Hello, world! How are you?")),
			updated:       createTestFile([]byte("Hello, world!")),
			expectedDelta: []byte("Hello, world!"),
		},
	}

	r := NewRollingHash(1024)
	// Loop through test cases and run test
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			delta, err := r.ComputeDelta(tc.original, tc.updated)
			if err != nil {
				t.Errorf("Test case %q failed with error: %v", tc.name, err)
			}
			if bytes.Compare(delta, tc.expectedDelta) != 0 {
				t.Errorf("Test case %q failed. Got %v, expected %v", tc.name, string(delta), string(tc.expectedDelta))
			}
		})
	}
}

// Helper function to create test files
func createTestFile(data []byte) *os.File {
	file, _ := ioutil.TempFile("", "")
	file.Write(data)
	file.Seek(0, 0)
	return file
}
