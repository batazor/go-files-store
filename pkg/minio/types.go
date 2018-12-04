package minio

type File struct {
	Bucket  string
	Name    string
	Payload []byte
	FileCH  chan File
}
