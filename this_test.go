package this

import (
	"testing"

	"github.com/zeebo/assert"
)

type thisTest struct{}

func (thisTest) method() string  { return This() }
func (thisTest) method2() string { return thisTest{}.method() }
func (thisTest) method3() string { return thisTest{}.method2() }

func (*thisTest) pmethod() string { return This() }

func TestThis(t *testing.T) {
	assert.Equal(t, This(),
		"github.com/zeebo/this.TestThis")
	assert.Equal(t, thisTest{}.method(),
		"github.com/zeebo/this.thisTest.method")
	assert.Equal(t, thisTest{}.method2(),
		"github.com/zeebo/this.thisTest.method")
	assert.Equal(t, thisTest{}.method3(),
		"github.com/zeebo/this.thisTest.method")
	assert.Equal(t, new(thisTest).pmethod(),
		"github.com/zeebo/this.(*thisTest).pmethod")
}

func BenchmarkThis(b *testing.B) {
	b.Run("Direct", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			This()
		}
	})

	b.Run("Inlined", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method()
		}
	})

	b.Run("InlinedTwice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method2()
		}
	})

	b.Run("InlinedThrice", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			thisTest{}.method3()
		}
	})
}
