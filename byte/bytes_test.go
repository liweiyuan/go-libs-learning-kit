package byte

import (
	"bytes"
	"testing"
)

type ReplaceMode int

const (
	ReplaceAll ReplaceMode = iota
	ReplaceFirst
	ReplaceNone
)

func replaceWithMode(s, old, new []byte, mode ReplaceMode) []byte {
	switch mode {
	case ReplaceAll:
		return bytes.ReplaceAll(s, old, new)
	case ReplaceFirst:
		return bytes.Replace(s, old, new, 1)
	case ReplaceNone:
		return bytes.Replace(s, old, new, 0)
	default:
		return s
	}
}

func TestBufferWriteRead(t *testing.T) {
	// Test cases for byte operations
	var buf bytes.Buffer
	//测试写入
	input := []byte("test data")
	n, err := buf.Write(input)
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}
	if n != len(input) {
		t.Errorf("Write returned wrong length: got %d, want %d", n, len(input))
	}

	//测试读取
	output := make([]byte, len(input))
	n, err = buf.Read(output)
	if err != nil {
		t.Errorf("Read failed: %v", err)
	}
	if n != len(input) {
		t.Errorf("Read returned wrong length: got %d, want %d", n, len(input))
	}
	if !bytes.Equal(input, output) {
		t.Errorf("Read returned wrong data: got %q, want %q", output, input)
	}

	//测试重置
	buf.Reset()
	if buf.Len() != 0 {
		t.Errorf("Reset failed: buffer length is %d, want 0", buf.Len())
	}
}

func TestBufferWriteStringReadString(t *testing.T) {
	// Test cases for string operations
	var buf bytes.Buffer
	//测试写入字符串
	input := "test string"
	n, err := buf.WriteString(input)
	if err != nil {
		t.Errorf("WriteString failed: %v", err)
	}
	if n != len(input) {
		t.Errorf("WriteString returned wrong length: got %d, want %d", n, len(input))
	}

	//测试读取字符串
	output, err := buf.ReadString('\n')
	if err != nil {
		t.Errorf("ReadString failed: %v", err)
	}
	if output != input {
		t.Errorf("ReadString returned wrong data: got %q, want %q", output, input)
	}

	//测试重置
	buf.Reset()
	if buf.Len() != 0 {
		t.Errorf("Reset failed: buffer length is %d, want 0", buf.Len())
	}
}

func TestBytesReplace(t *testing.T) {
	// Test cases for bytes.Replace
	input := []byte("hello world")
	old := []byte("world")
	new := []byte("go")
	output := replaceWithMode(input, old, new, ReplaceAll)
	expected := []byte("hello go")
	if !bytes.Equal(output, expected) {
		t.Errorf("Bytes.Replace returned wrong data: got %q, want %q", output, expected)
	}

	// Test cases for bytes.ReplaceAll
	input = []byte("hello world")
	old = []byte("o")
	new = []byte("0")
	output = replaceWithMode(input, old, new, ReplaceAll)
	expected = []byte("hell0 w0rld")
	if !bytes.Equal(output, expected) {
		t.Errorf("Bytes.ReplaceAll returned wrong data: got %q, want %q", output, expected)
	}
}

func TestReplaceWithMode(t *testing.T) {
	tests := []struct {
		s, old, new []byte
		mode        ReplaceMode
		expected    []byte
	}{
		{[]byte("go gopher go"), []byte("go"), []byte("Go"), ReplaceAll, []byte("Go Gopher Go")},
		{[]byte("go gopher go"), []byte("go"), []byte("Go"), ReplaceFirst, []byte("Go gopher go")},
		{[]byte("go gopher go"), []byte("go"), []byte("Go"), ReplaceNone, []byte("go gopher go")},
		{[]byte("hello world"), []byte("x"), []byte("y"), ReplaceAll, []byte("hello world")},
	}

	for _, test := range tests {
		result := replaceWithMode(test.s, test.old, test.new, test.mode)
		if !bytes.Equal(result, test.expected) {
			t.Errorf("ReplaceWithMode(%q, %q, %q, %v) = %q; want %q",
				test.s, test.old, test.new, test.mode, result, test.expected)
		}
	}
}
