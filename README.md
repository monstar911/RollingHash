# Rolling Hash Algorithm

Make a rolling hash based file diffing algorithm. When comparing original and an updated version of an input, it should return a description ("delta") which can be used to upgrade an original version of the file into the new file. The description contains the chunks which:
- Can be reused from the original file
- have been added or modified and thus would need to be synchronized

The real-world use case for this type of construct could be a distributed file storage system. This reduces the need for bandwidth and storage. If many people have the same file stored on Dropbox, for example, there's no need to upload it again.

A library that does a similar thing is [rdiff](https://linux.die.net/man/1/rdiff). You don't need to fulfill the patch part of the API, only signature and delta.

## Requirements
- Hashing function gets the data as a parameter. Separate possible filesystem operations.
- Chunk size can be fixed or dynamic, but must be split to at least two chunks on any sufficiently sized data.
- Should be able to recognize changes between chunks. Only the exact differing locations should be added to the delta.
- Well-written unit tests function well in describing the operation, no UI necessary.

## Checklist
1. Input/output operations are separated from the calculations
2. detects chunk changes and/or additions
3. detects chunk removals
4. detects additions between chunks with shifted original chunks


## Technical details

This package implements a rolling hash algorithm to compute the delta (differences) between two files. The algorithm works by reading both the original and updated versions of the file in chunks, computing the hash of each chunk using the SHA256 hash function, and then comparing the hashes of the chunks from the original and updated files.

The package includes the following main components:

* A `chunk` struct that contains the hash of a chunk of bytes, and the bytes themselves.
* A `RollingHash` struct that contains a hash.Hash interface and the size of the chunks to be read from the file.
* The `NewRollingHash` function that returns a new `RollingHash` struct.
* The `ComputeHashes` method that reads a file in chunks and computes the hash of each chunk using the SHA256 hash function.
* The `ComputeDelta` method that generates a description of the differences between the original and updated versions of a file, by comparing the hashes of the chunks from the original and updated files.

To use the package, you can create a new `RollingHash` struct and call the `ComputeDelta` method to get the delta between two files. You can also call the `ComputeHashes` method to get the list of chunks and their corresponding hashes for a file.

## run and test command
* go run cmd/main.go
* go test ./... -v count=1