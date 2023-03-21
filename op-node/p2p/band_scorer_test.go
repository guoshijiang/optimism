package p2p

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestBandScorer_ParseDefault tests the [BandScorer.Parse] function
// on the default band scores cli flag value.
func TestBandScorer_ParseDefault(t *testing.T) {
	// Create a new band scorer.
	bandScorer := NewBandScorer()
	require.NoError(t, bandScorer.Parse("-40:graylist;-20:restricted;0:nopx;20:friend;"))

	// Validate the [BandScorer] internals.
	require.ElementsMatch(t, bandScorer.bands, []scorePair{
		{band: "graylist", threshold: -40},
		{band: "restricted", threshold: -20},
		{band: "nopx", threshold: 0},
		{band: "friend", threshold: 20},
	})
}

// TestBandScorer_BucketCorrectly tests the [BandScorer.Bucket] function
// on a variety of scores.
func TestBandScorer_BucketCorrectly(t *testing.T) {
	// Create a new band scorer.
	bandScorer := NewBandScorer()
	require.NoError(t, bandScorer.Parse("-40:graylist;-20:restricted;0:nopx;20:friend;"))

	// Let's validate that the [BandScorer] sorts the bands correctly.
	require.Equal(t, bandScorer.bands, []scorePair{
		{band: "graylist", threshold: -40},
		{band: "restricted", threshold: -20},
		{band: "nopx", threshold: 0},
		{band: "friend", threshold: 20},
	})

	// Validate the [BandScorer] internals.
	require.Equal(t, bandScorer.Bucket(-100), "graylist")
	require.Equal(t, bandScorer.Bucket(-40), "graylist")
	require.Equal(t, bandScorer.Bucket(-39), "restricted")
	require.Equal(t, bandScorer.Bucket(-20), "restricted")
	require.Equal(t, bandScorer.Bucket(-19), "nopx")
	require.Equal(t, bandScorer.Bucket(0), "nopx")
	require.Equal(t, bandScorer.Bucket(1), "friend")
	require.Equal(t, bandScorer.Bucket(20), "friend")
	require.Equal(t, bandScorer.Bucket(21), "friend")
}

// TestBandScorer_BucketInverted tests the [BandScorer.Bucket] function
// on a variety of scores, in descending order.
func TestBandScorer_BucketInverted(t *testing.T) {
	// Create a new band scorer.
	bandScorer := NewBandScorer()
	require.NoError(t, bandScorer.Parse("20:friend;0:nopx;-20:restricted;-40:graylist;"))

	// Let's validate that the [BandScorer] sorts the bands correctly.
	require.Equal(t, bandScorer.bands, []scorePair{
		{band: "graylist", threshold: -40},
		{band: "restricted", threshold: -20},
		{band: "nopx", threshold: 0},
		{band: "friend", threshold: 20},
	})

	// Validate the [BandScorer] internals.
	require.Equal(t, bandScorer.Bucket(-100), "graylist")
	require.Equal(t, bandScorer.Bucket(-40), "graylist")
	require.Equal(t, bandScorer.Bucket(-39), "restricted")
	require.Equal(t, bandScorer.Bucket(-20), "restricted")
	require.Equal(t, bandScorer.Bucket(-19), "nopx")
	require.Equal(t, bandScorer.Bucket(0), "nopx")
	require.Equal(t, bandScorer.Bucket(1), "friend")
	require.Equal(t, bandScorer.Bucket(20), "friend")
	require.Equal(t, bandScorer.Bucket(21), "friend")
}

// TestBandScorer_ParseEmpty tests the [BandScorer.Parse] function
// on an empty string.
func TestBandScorer_ParseEmpty(t *testing.T) {
	// Create a band scorer on an empty string.
	bandScorer := NewBandScorer()
	require.NoError(t, bandScorer.Parse(""))

	// Validate the [BandScorer] internals.
	require.Len(t, bandScorer.bands, 0)
}

// TestBandScorer_ParseWhitespace tests the [BandScorer.Parse] function
// on a variety of whitespaced strings.
func TestBandScorer_ParseWhitespace(t *testing.T) {
	// Create a band scorer on an empty string.
	bandScorer := NewBandScorer()
	require.NoError(t, bandScorer.Parse("  ;  ;  ;  "))

	// Validate the [BandScorer] internals.
	require.Len(t, bandScorer.bands, 0)
}
