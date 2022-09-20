package roundrobin_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanhpp/gopher/pkg/test/roundrobin"
)

type TestInterface interface {
	Print()
}

type TestIplm struct {
	v int
}

func (t *TestIplm) Print() {
	log.Println(t.v)
}

func NewTestIplm(i int) TestInterface {
	return &TestIplm{i}
}

func TestRoundRobin(t *testing.T) {
	var (
		sl = []TestInterface{NewTestIplm(1), NewTestIplm(2), NewTestIplm(3)}
	)

	rr := roundrobin.New()
	for i := range sl {
		rr.Add(sl[i])
	}

	for i := 0; i < 10; i++ {
		v, ok := rr.Next()
		assert.True(t, ok)
		testV, ok := v.(*TestIplm)
		assert.True(t, ok)
		testV.Print()
	}
}
