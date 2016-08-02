package main

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

/*
	main.go에 165line을 보면 io.Copy() 메서드를 통해
	r의 내용을 os.Stdout으로 복사하고 있다. 이때,
	r은 사용자가 지정한 타입(rot13Reader)이다. 
	Copy(dst Writer, src Reader)이기 때문에 사용자 지정
	타입 r은 Reader인터페이스의 Read메서드를 구현해야한다.
	 
*/

func (rt13 *rot13Reader) Read(p []byte) (n int, err error) {
	length, err := rt13.r.Read(p)
	for i:=0; i < length; i++ {
		p[i] = rot13(p[i])
	}
	return len(p), err
}

func rot13(p byte) byte{
	sb := rune(p)
	if (( 'a' < sb ) && ( 'm' > sb )) || 
			(( 'A' < sb ) && ( 'M' > sb )) {
		p += 13		
	} else if (( 'n' < sb ) && ( 'z' > sb )) || 
					(( 'N' < sb ) && ( 'Z' > sb )) {
		p -= 13
	}
	return p
}


