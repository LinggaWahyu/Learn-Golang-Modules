package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	bencmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Lingga",
			request: "Lingga",
		},
		{
			name:    "Wahyu",
			request: "Wahyu",
		},
		{
			name:    "Rochim",
			request: "Rochim",
		},
	}

	for _, bencbencmark := range bencmarks {
		b.Run(bencbencmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bencbencmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Lingga", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Lingga")
		}
	})
	b.Run("Wahyu", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Wahyu")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lingga")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Lingga",
			request:  "Lingga",
			expected: "Hello Lingga",
		},
		{
			name:     "Wahyu",
			request:  "Wahyu",
			expected: "Hello Wahyu",
		},
		{
			name:     "Rochim",
			request:  "Rochim",
			expected: "Hello Rochim",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be '"+test.expected+"'")
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Lingga", func(t *testing.T) {
		result := HelloWorld("Lingga")
		require.Equal(t, "Hello Lingga", result, "Result must be 'Hello Lingga'")
	})

	t.Run("Wahyu", func(t *testing.T) {
		result := HelloWorld("Wahyu")
		require.Equal(t, "Hi Wahyu", result, "Result must be 'Hello Wahyu'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}

	result := HelloWorld("Lingga")
	require.Equal(t, "Hello Lingga", result, "Result must be 'Hello Lingga'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Lingga")
	assert.Equal(t, "Hello Lingga", result, "Result must be 'Hello Lingga'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Lingga")
	require.Equal(t, "Hello Lingga", result, "Result must be 'Hello Lingga'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldOne(t *testing.T) {
	result := HelloWorld("Lingga")

	if result != "Hello Lingga" {
		// error
		t.Error("Result must be Hello Lingga")
	}

	fmt.Println("Test Hello World One Done")
}

func TestHelloWorldTwo(t *testing.T) {
	result := HelloWorld("Lingga Two")

	if result != "Hello Lingga Two" {
		// error
		t.Fatal("Result must be Hello Lingga Two")
	}

	fmt.Println("Test Hello World Two Done")
}
