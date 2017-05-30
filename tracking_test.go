package usps

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrackPackage(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	output, err := usps.TrackPackage("9341989949036022338924")
	require.Nil(t, err)
	require.Equal(t, output.TrackInfo.TrackSummary, "The Postal Service could not locate the tracking information for your request. Please verify your tracking number and try again later.")
}
