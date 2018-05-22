package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cmn "github.com/tendermint/tmlibs/common"
)

func TestABCIResults(t *testing.T) {

	// Nil and []byte{} should produce the same hash.
	a := ABCIResult{Code: 0, Data: nil}
	b := ABCIResult{Code: 0, Data: []byte{}}
	require.Equal(t, a.Hash(), a.Hash())
	require.Equal(t, b.Hash(), b.Hash())
	require.Equal(t, a.Hash(), b.Hash())

	c := ABCIResult{Code: 0, Data: []byte("one"),
		Tags: cmn.KVPairs{
			{[]byte("key"), []byte("value")},
		}}
	d := ABCIResult{Code: 14, Data: nil}
	e := ABCIResult{Code: 14, Data: []byte("foo")}
	f := ABCIResult{Code: 14, Data: []byte("bar")}

	// a and b should be the same, don't go in results.
	results := ABCIResults{a, c, d, e, f}

	// Make sure each result hashes properly.
	var last []byte
	for i, res := range results {
		h := res.Hash()
		assert.NotEqual(t, last, h, "%d", i)
		last = h
	}

	// Make sure that we can get a root hash from results and verify proofs.
	root := results.Hash()
	assert.NotEmpty(t, root)

	for i, res := range results {
		proof := results.ProveResult(i)
		valid := proof.Verify(i, len(results), res.Hash(), root)
		assert.True(t, valid, "%d", i)
	}
}

func TestABCIResultTags(t *testing.T) {

	a := ABCIResult{Code: 0, Data: []byte("one"),
		Tags: cmn.KVPairs{
			{[]byte("key"), []byte("value")},
		}}
	b := ABCIResult{Code: 0, Data: []byte("one"),
		Tags: cmn.KVPairs{
			{[]byte("keyv"), []byte("alue")},
		}}

	assert.NotEqual(t, a.Hash(), b.Hash())
}
