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

// replaceWithMode 根据指定的替换模式替换字节切片中的内容。
// 参数：
//
//	s - 原始字节切片
//	old - 要被替换的字节切片
//	new - 替换成的字节切片
//	mode - 替换模式，支持 ReplaceAll、ReplaceFirst 和 ReplaceNone
//
// 返回值：
//
//	返回替换后的字节切片
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
	output := buf.String()
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

// 测试边界情况
func TestBufferEdgeCases(t *testing.T) {
	var buf bytes.Buffer

	// 测试写入空切片
	n, err := buf.Write([]byte{})
	if err != nil || n != 0 {
		t.Errorf("Writing empty slice failed: got n=%d, err=%v, want n=0, err=nil", n, err)
	}

	// 测试写入nil切片
	n, err = buf.Write(nil)
	if err != nil || n != 0 {
		t.Errorf("Writing nil slice failed: got n=%d, err=%v, want n=0, err=nil", n, err)
	}

	// 测试读取到空切片
	n, err = buf.Read([]byte{})
	if err != nil || n != 0 {
		t.Errorf("Reading into empty slice failed: got n=%d, err=%v, want n=0, err=nil", n, err)
	}

	// 测试读取到nil切片
	n, err = buf.Read(nil)
	if err != nil || n != 0 {
		t.Errorf("Reading into nil slice failed: got n=%d, err=%v, want n=0, err=nil", n, err)
	}
}

// 测试Buffer的其他方法
func TestBufferAdvancedOperations(t *testing.T) {
	var buf bytes.Buffer

	// 测试Grow
	buf.Grow(100)
	input := []byte("test data")
	buf.Write(input)
	if buf.Cap() < 100 {
		t.Errorf("Buffer capacity after Grow is %d, want >= 100", buf.Cap())
	}

	// 测试Truncate
	buf.Truncate(4)
	if buf.String() != "test" {
		t.Errorf("Buffer after Truncate is %q, want %q", buf.String(), "test")
	}

	// 测试Next
	buf.Reset()
	buf.WriteString("hello world")
	first5 := buf.Next(5)
	if string(first5) != "hello" {
		t.Errorf("Buffer.Next(5) = %q, want %q", first5, "hello")
	}
	remaining := buf.String()
	if remaining != " world" {
		t.Errorf("Remaining buffer after Next = %q, want %q", remaining, " world")
	}
}

// 性能测试
func BenchmarkBufferWrite(b *testing.B) {
	data := []byte("hello world")
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Write(data)
		buf.Reset()
	}
}

func BenchmarkBufferWriteString(b *testing.B) {
	data := "hello world"
	var buf bytes.Buffer
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString(data)
		buf.Reset()
	}
}

func BenchmarkReplaceWithMode(b *testing.B) {
	s := []byte("go gopher go gopher go")
	old := []byte("go")
	new := []byte("Go")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		replaceWithMode(s, old, new, ReplaceAll)
	}
}
