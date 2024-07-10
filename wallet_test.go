package utils

import (
	"fmt"
	"strings"
	"testing"
)

func ExamplePrivate2Address() {
	fmt.Println(Private2Address("3fe74272f214ddb0d182a9ae68e3ee844d6f61d04f6e7827aad972e071a31f81"))
	// Output:
	// 0x5750bd75ADC81367D1b7510A57D79ACF777950a1
}

func ExampleChecksumEthereumAddress() {
	fmt.Println(ChecksumEthereumAddress("0x5750bd75ADC81367D1b7510A57D79ACF777950a1")) // checksumed
	fmt.Println(ChecksumEthereumAddress("0x5750bd75adc81367d1b7510a57d79acf777950a1")) // all lower case
	fmt.Println(ChecksumEthereumAddress("0X5750BD75ADC81367D1B7510A57D79ACF777950A1")) // all upper case
	// Output:
	// 0x5750bd75ADC81367D1b7510A57D79ACF777950a1
	// 0x5750bd75ADC81367D1b7510A57D79ACF777950a1
	// 0x5750bd75ADC81367D1b7510A57D79ACF777950a1
}

func TestWallet(t *testing.T) {
	var addresses []string = []string{
		// "0xf611aeb5013fd2c0511c9cd55c7dc5c1140741a6", "0xb5b46f918c2923fc7f26db76e8a6a6e9c4347cf9",
		// "0x6b030ff3fb9956b1b69f475b77ae0d3cf2cc5afa", "0x18248226c16bf76c032817854e7c83a2113b4f06",
		// "0x4a1c3ad6ed28a636ee1751c69071f6be75deb8b8", "0x953a573793604af8d41f306feb8274190db4ae0e",
		// "0xe80761ea617f66f96274ea5e8c37f03960ecc679", "0xd8ad37849950903571df17049516a5cd4cbe55f6",
		// "0x7f45273fd7c644714825345670414ea649b50b16", "0x49b0c695039243bbfeb8ecd054eb70061fd54aa0",
	}

	for i := 0; i < len(addresses); i++ {
		printWalletAddress(addresses[i])
	}
}

func printWalletAddress(addr string) {
	if addr == "" || len(addr) != 42 || addr[:2] != "0x" {
		return
	}

	lower := strings.ToLower(addr)
	checksummed := ChecksumEthereumAddress(lower)

	fmt.Println("\nAddress:", addr)
	fmt.Println("\tchecksum:", checksummed)
	fmt.Println("\tlower:   ", lower)
}
