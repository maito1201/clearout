package clearout

import (
	"fmt"
	"io"
	"os"
)

type Output struct {
	Out    io.Writer
	Prefix string
	Suffix string
	buf    []byte
}

const CLEAR = "\033[2J\033[0;0H"

func (o *Output) Render() {
	if o.Out == nil {
		o.Out = os.Stdout
	}
	fmt.Fprint(o.Out, CLEAR)
	fmt.Fprint(o.Out, o.Prefix)
	fmt.Fprint(o.Out, string(o.buf))
	fmt.Fprint(o.Out, o.Suffix)
	o.Flush()
}

func (o *Output) Print(i interface{}) *Output {
	o.buf = append(o.buf, []byte(fmt.Sprintf("%v", i))...)
	return o
}

func (o *Output) Println(i interface{}) *Output {
	o.buf = append(o.buf, []byte(fmt.Sprintf("%v\n", i))...)
	return o
}

func (o *Output) Printf(format string, args ...interface{}) *Output {
	o.buf = append(o.buf, []byte(fmt.Sprintf(format, args...))...)
	return o
}

func (o *Output) WithPrefix(i interface{}) *Output {
	o.Prefix = fmt.Sprintf("%v", i)
	return o
}

func (o *Output) WithSuffix(i interface{}) *Output {
	o.Suffix = fmt.Sprintf("%v", i)
	return o
}

func (o *Output) Flush() {
	o.buf = nil
}
