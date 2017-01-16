package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Boilerplate

func Assert(condition bool, items ...interface{}) {
	if !condition {
		panic("assertion failed: " + fmt.Sprint(items...))
	}
}

func Log(items ...interface{}) {
	fmt.Println(items...)
}

var Input = bufio.NewReader(os.Stdin)

func ReadByte() byte {
	b, e := Input.ReadByte()
	if e != nil {
		panic(e)
	}
	return b
}

func MaybeReadByte() (byte, bool) {
	b, e := Input.ReadByte()
	if e != nil {
		if e == io.EOF {
			return 0, false
		}
		panic(e)
	}
	return b, true
}

func UnreadByte() {
	e := Input.UnreadByte()
	if e != nil {
		panic(e)
	}
}

func NewLine() {
	for {
		b := ReadByte()
		switch b {
		case ' ', '\t', '\r':
			// keep looking
		case '\n':
			return
		default:
			panic(fmt.Sprintf("expecting newline, but found character <%c>", b))
		}
	}
}

func ScanInt(low, high int) int {
	return int(ScanInt64(int64(low), int64(high)))
}

func ScanUint(low, high uint) uint {
	return uint(ScanUint64(uint64(low), uint64(high)))
}

func ScanInt64(low, high int64) int64 {
	Assert(low <= high)
	for {
		b := ReadByte()
		switch b {
		case ' ', '\t', '\r':
			// keep looking
		case '\n':
			panic(fmt.Sprintf(
				"unexpected newline; expecting range %d..%d", low, high))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if high < 0 {
				panic(fmt.Sprintf(
					"found <%c> but expecting range %d..%d", b, low, high))
			}
			lw := low
			if lw < 0 {
				lw = 0
			}
			UnreadByte()
			x, e := _scanu64(uint64(lw), uint64(high))
			if e != "" {
				panic(fmt.Sprintf("%s %d..%d", e, low, high))
			}
			return int64(x)
		case '-':
			if low > 0 {
				panic(fmt.Sprintf(
					"found minus sign but expecting range %d..%d", low, high))
			}
			h := high
			if h > 0 {
				h = 0
			}
			x, e := _scanu64(uint64(-h), uint64(-low))
			if e != "" {
				panic(fmt.Sprintf("-%s %d..%d", e, low, high))
			}
			return -int64(x)
		default:
			panic(fmt.Sprintf(
				"unexpected character <%c>; expecting range %d..%d", b, low, high))
		}
	}
}

func ScanUint64(low, high uint64) uint64 {
	Assert(low <= high)
	for {
		b := ReadByte()
		switch b {
		case ' ', '\t', '\r':
			// keep looking
		case '\n':
			panic(fmt.Sprintf(
				"unexpected newline; expecting range %d..%d", low, high))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			UnreadByte()
			x, e := _scanu64(low, high)
			if e != "" {
				panic(fmt.Sprintf("%s %d..%d", e, low, high))
			}
			return x
		default:
			panic(fmt.Sprintf(
				"unexpected character <%c>; expecting range %d..%d", b, low, high))
		}
	}
}

func _scanu64(low, high uint64) (result uint64, err string) {
	x := uint64(0)
buildnumber:
	for {
		b, ok := MaybeReadByte()
		if !ok {
			break buildnumber
		}
		switch b {
		case ' ', '\t', '\r':
			break buildnumber
		case '\n':
			UnreadByte()
			break buildnumber
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			d := uint64(b - '0')
			if (high-d)/10 < x {
				return x, fmt.Sprintf("%d%c... not in range", x, b)
			}
			x = x*10 + d
		default:
			return x, fmt.Sprintf("%d%c found; expecting range", x, b)
		}
	}
	if x < low || x > high {
		return x, fmt.Sprintf("%d not in range", x)
	}
	return x, ""
}

func ScanBytes(short, long int) []byte {
	Assert(1 <= short && short <= long)
	var b byte
	buf := make([]byte, long)
skipws:
	for {
		b = ReadByte()
		switch b {
		case ' ', '\t', '\r':
			// keep looking
		case '\n':
			panic(fmt.Sprintf("unexpected newline; expecting string"))
		default:
			break skipws
		}
	}
	buf[0] = b
	length := 1
buildstring:
	for {
		var ok bool
		b, ok = MaybeReadByte()
		if !ok {
			break buildstring
		}
		switch b {
		case ' ', '\t', '\r':
			break buildstring
		case '\n':
			UnreadByte()
			break buildstring
		default:
			if length >= long {
				panic(fmt.Sprintf("string length not in range %d..%d", short, long))
			}
			buf[length] = b
			length++
		}
	}
	if length < short {
		panic(fmt.Sprintf("string length not in range %d..%d", short, long))
	}
	return buf[:length]
}

func ScanString(short, long int) string {
	return string(ScanBytes(short, long))
}

// Math

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}

// Sort

type Ints64 []int64

func (a Ints64) Len() int           { return len(a) }
func (a Ints64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ints64) Less(i, j int) bool { return a[i] < a[j] }

/* Sort tempalte
func (a ) Len() int           { return len(a) }
func (a ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ) Less(i, j int) bool { return  }
*/
