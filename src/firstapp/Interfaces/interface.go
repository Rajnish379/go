package main 

import (
	"fmt"
	"bytes"
	"io"
)

func main(){
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0;i<10;i++ {
		fmt.Println(inc.Increment())
	}
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello peeps, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	//wc.(io.Reader)
	// It fails because our wc is expecting a write object not a read one
	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello YouTube listeners, this is a test"))
		wc.Close()
	}
	s,ok := myObj.(io.Reader)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("Conversion failed")
	}

	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}

	var newWc WriterCloser = myWriterCloser{}
	fmt.Println(newWc)

}

// In go we doesn't use the implements keyword to implement interfaces unlike other programming languages

type Writer interface {
	Write([]byte) (int,error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct {}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int,error) {
	n, err := fmt.Println(string(data))
	return n,err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err!= nil {
		return 0,err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_,err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0,err
		}
	}
	return n,nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (mwc myWriterCloser) Write(data []byte) (int,error) {
	return 0, nil
}

func (mwc myWriterCloser) Close() error {
	return nil
}

// If any of the methods require a pointer receiver in an interface, you're going to have to implement that interface with a pointer only
// But if they don't require pointer receiver at all, you can work with both values and pointers