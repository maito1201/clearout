package clearout

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPrint(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want string
	}{
		{
			name: "string",
			in:   "test",
			want: "\033[2J\033[0;0Htest",
		},
		{
			name: "int",
			in:   1,
			want: "\033[2J\033[0;0H1",
		},
		{
			name: "nil",
			in:   "nil",
			want: "\033[2J\033[0;0Hnil",
		},
	}

	for _, tt := range tests {
		testOut := bytes.NewBuffer([]byte{})
		o := Output{Out: testOut}
		o.Print(tt.in).Render()
		if diff := cmp.Diff(string(testOut.Bytes()), tt.want); diff != "" {
			t.Errorf("unexpected out (got:-, want:+) %v", diff)
		}
	}
}

func TestPrintln(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want string
	}{
		{
			name: "string",
			in:   "test",
			want: "\033[2J\033[0;0Htest\n",
		},
		{
			name: "int",
			in:   1,
			want: "\033[2J\033[0;0H1\n",
		},
		{
			name: "nil",
			in:   nil,
			want: "\033[2J\033[0;0H<nil>\n",
		},
	}

	for _, tt := range tests {
		testOut := bytes.NewBuffer([]byte{})
		o := Output{Out: testOut}
		o.Println(tt.in).Render()
		if diff := cmp.Diff(string(testOut.Bytes()), tt.want); diff != "" {
			t.Errorf("unexpected out (got:-, want:+) %v", diff)
		}
	}
}

func TestPrintf(t *testing.T) {
	tests := []struct {
		name   string
		in     interface{}
		format string
		want   string
	}{
		{
			name:   "string",
			in:     "test",
			format: "%s",
			want:   "\033[2J\033[0;0Htest",
		},
		{
			name:   "int",
			in:     1,
			format: "%02d",
			want:   "\033[2J\033[0;0H01",
		},
		{
			name:   "nil",
			in:     nil,
			format: "%v",
			want:   "\033[2J\033[0;0H<nil>",
		},
	}

	for _, tt := range tests {
		testOut := bytes.NewBuffer([]byte{})
		o := Output{Out: testOut}
		o.Printf(tt.format, tt.in).Render()
		if diff := cmp.Diff(string(testOut.Bytes()), tt.want); diff != "" {
			t.Errorf("unexpected out (got:-, want:+) %v", diff)
		}
	}
}

func TestWithPrefix(t *testing.T) {
	tests := []struct {
		name   string
		in     interface{}
		prefix interface{}
		want   string
	}{
		{
			name:   "string",
			in:     "test",
			prefix: "prefix",
			want:   "\033[2J\033[0;0Hprefixtest",
		},
		{
			name:   "int",
			in:     2,
			prefix: 1,
			want:   "\033[2J\033[0;0H12",
		},
		{
			name:   "nil",
			in:     nil,
			prefix: nil,
			want:   "\033[2J\033[0;0H<nil><nil>",
		},
	}

	for _, tt := range tests {
		testOut := bytes.NewBuffer([]byte{})
		o := Output{Out: testOut}
		o.Print(tt.in).WithPrefix(tt.prefix).Render()
		if diff := cmp.Diff(string(testOut.Bytes()), tt.want); diff != "" {
			t.Errorf("unexpected out (got:-, want:+) %v", diff)
		}
	}
}

func TestWithSuffix(t *testing.T) {
	tests := []struct {
		name   string
		in     interface{}
		suffix interface{}
		want   string
	}{
		{
			name:   "string",
			in:     "test",
			suffix: "suffix",
			want:   "\033[2J\033[0;0Htestsuffix",
		},
		{
			name:   "int",
			in:     1,
			suffix: 2,
			want:   "\033[2J\033[0;0H12",
		},
		{
			name:   "nil",
			in:     nil,
			suffix: nil,
			want:   "\033[2J\033[0;0H<nil><nil>",
		},
	}

	for _, tt := range tests {
		testOut := bytes.NewBuffer([]byte{})
		o := Output{Out: testOut}
		o.Print(tt.in).WithSuffix(tt.suffix).Render()
		if diff := cmp.Diff(string(testOut.Bytes()), tt.want); diff != "" {
			t.Errorf("unexpected out (got:-, want:+) %v", diff)
		}
	}
}

func TestRender(t *testing.T) {
	testOut := bytes.NewBuffer([]byte{})
	want := "\x1b[2J\x1b[0;0H"
	o := Output{Out: testOut}
	o.Render()
	if diff := cmp.Diff(string(testOut.Bytes()), want); diff != "" {
		t.Errorf("unexpected out (got:-, want:+) %v", diff)
	}
}

func TestFlush(t *testing.T) {
	testOut := bytes.NewBuffer([]byte{})
	want := "\x1b[2J\x1b[0;0H"
	o := Output{Out: testOut}
	o.Print("test")
	o.Flush()
	o.Render()
	if diff := cmp.Diff(string(testOut.Bytes()), want); diff != "" {
		t.Errorf("unexpected out (got:-, want:+) %v", diff)
	}
}
