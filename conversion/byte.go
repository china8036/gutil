package conversion

import "fmt"

type ByteSize int64

func (bs ByteSize) String() string {
	if bs < 1024 {
		return fmt.Sprintf("%d B", bs)
	}
	if bs < 1024*1024 {
		return fmt.Sprintf("%.2f K", float64(bs)/1024)
	}
	if bs < 1024*1024*1024 {
		return fmt.Sprintf("%.2f M", float64(bs)/(1024*1024))
	}

	if bs < 1024*1024*1024*1024 {
		return fmt.Sprintf("%2.f G", float64(bs)/(1024*1024*1024))
	}
	return fmt.Sprintf("%2.f G", float64(bs)/(1024*1024*1024*1024))
}
