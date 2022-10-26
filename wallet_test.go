package utils

import (
	"fmt"
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
