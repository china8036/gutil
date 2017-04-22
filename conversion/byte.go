package conversion

import "fmt"

type ByteSize int64

func (bs ByteSize) String() string {
	if bs < 1024 {
		return fmt.Sprintf("%d B", bs)
	}
	if bs < 1024*1024 {
		return fmt.Sprintf("%2f K", float32(bs/1024))
	}
	if bs < 1024*1024*1024 {
		return fmt.Sprintf("%2f M", float32(bs/(1024*1024)))
	}

	if bs < 1024*1024*1024*1024 {
		return fmt.Sprintf("%2f G", float32(bs/(1024*1024*1024)))
	}
	return fmt.Sprintf("%2f G", float32(bs/(1024*1024*1024*1024)))
}
