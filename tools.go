package main

type Tools struct {

}


func ( this *Tools ) PadStringLeft (s, p string, count int) string {

    ret := make( []byte, len( p ) * count + len( s ) )

    b := ret[:len( p ) * count ]
    bp := copy( b, p )
    for bp < len( b ) {
        copy( b [bp:], b[:bp] )
        bp *= 2
    }
    copy( ret [ len( b ):], s )
    return string( ret )
}

func ( this *Tools ) fillStringLeft ( s string, count int, fillingChar string) string {

	strLen := len ( s )
	fillCount := 0
	if count > strLen {
		fillCount = count - strLen
	}

	result := ""
	for i:=0; i < fillCount; i++ {
		result += fillingChar
	}

	result += s
	return result
}
