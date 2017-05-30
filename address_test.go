package usps

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddressVerification(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output, err := usps.AddressVerification(address)
	require.Nil(t, err)
	require.Equal(t, output.Address.Address2, "6406 IVY LN")
}

func TestZipCodeLookup(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var address Address
	address.Address2 = "6406 Ivy Lane"
	address.City = "Greenbelt"
	address.State = "MD"

	output, err := usps.ZipCodeLookup(address)
	require.Nil(t, err)

	require.Equal(t, output.Address.Address2, "6406 IVY LN")
}

func TestCityStateLookup(t *testing.T) {
	var usps USPS
	usps.Username = os.Getenv("USPSUsername")

	var address ZipCode
	address.Zip5 = "90210"

	output, err := usps.CityStateLookup(address)
	require.Nil(t, err)
	require.Equal(t, output.ZipC.Zip5, "90210")
}
