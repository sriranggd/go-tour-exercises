package main

import "io"

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	upRow1Start := int('A')
	upRow2Start := upRow1Start + 13
	lowRow1Start := int('a')
	lowRow2Start := lowRow1Start + 13

	buffer := make([]byte, len(b))
	count, err := r13.r.Read(buffer)
	if err != nil {
		return count, err
	}

	for i := range buffer {
		c := int(buffer[i])
		if c >= upRow1Start && c < upRow1Start+26 {
			if c >= upRow2Start {
				b[i] = byte(upRow1Start + c - upRow2Start)
			} else {
				b[i] = byte(c + 13)
			}
		} else if c >= lowRow1Start && c < lowRow1Start+26 {
			if c >= lowRow2Start {
				b[i] = byte(lowRow1Start + c - lowRow2Start)
			} else {
				b[i] = byte(c + 13)
			}
		} else {
			b[i] = byte(c)
		}
	}

	return count, nil
}
