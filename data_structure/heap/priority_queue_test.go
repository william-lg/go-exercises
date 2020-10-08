package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayMaxPQ(t *testing.T) {
	var pq ArrayMaxPQ
	pq = make([]int, 0)

	pq.Insert(1)
	assert.Equal(t, 1, pq.Maximum())
	pq.Insert(2)
	assert.Equal(t, 2, pq.Maximum())
	pq.Insert(4)
	pq.Insert(5)
	assert.Equal(t, 5, pq.Maximum())

	v := pq.ExtractMax()
	assert.Equal(t, 5, v)
	assert.Equal(t, 4, pq.Maximum())

	pq.Insert(7)
	pq.Insert(8)
	pq.IncreaseKey(len(pq)-1, 10)
	assert.Equal(t, 10, pq.Maximum())
	pq.IncreaseKey(1, 11)
	assert.Equal(t, 11, pq.Maximum())
}
