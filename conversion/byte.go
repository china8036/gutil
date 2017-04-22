package conversion

import "fmt"

type ByteSize int64

func (bs ByteSize) String() string {
	if bs < 1024 {
		return fmt.Sprintf("%d B", bs)
	}
	if bs < 1024*1024 {
		return fmt.Sprintf("%2f K", bs/1024)
	}
	if bs < 1024*1024*1024 {
		return fmt.Sprintf("%2f M", bs/(1024*1024))
	}

	if bs < 1024*1024*1024*1024 {
		return fmt.Sprintf("%2f G", bs/(1024*1024*1024))
	}
	return fmt.Sprintf("%2f G", bs/(1024*1024*1024*1024))
}
