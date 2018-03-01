package navutil_test

import (
	"fmt"
	"math"

	"github.com/navcoin/navutil"
)

func ExampleAmount() {

	a := navutil.Amount(0)
	fmt.Println("Zero Satoshi:", a)

	a = navutil.Amount(1e8)
	fmt.Println("100,000,000 Satoshis:", a)

	a = navutil.Amount(1e5)
	fmt.Println("100,000 Satoshis:", a)
	// Output:
	// Zero Satoshi: 0 NAV
	// 100,000,000 Satoshis: 1 NAV
	// 100,000 Satoshis: 0.001 NAV
}

func ExampleNewAmount() {
	amountOne, err := navutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := navutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := navutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := navutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 NAV
	// 0.01234567 NAV
	// 0 NAV
	// invalid bitcoin amount
}

func ExampleAmount_unitConversions() {
	amount := navutil.Amount(44433322211100)

	fmt.Println("Satoshi to kNAV:", amount.Format(navutil.AmountKiloNAV))
	fmt.Println("Satoshi to NAV:", amount)
	fmt.Println("Satoshi to MilliNAV:", amount.Format(navutil.AmountMilliNAV))
	fmt.Println("Satoshi to MicroNAV:", amount.Format(navutil.AmountMicroNAV))
	fmt.Println("Satoshi to Satoshi:", amount.Format(navutil.AmountSatoshi))

	// Output:
	// Satoshi to kNAV: 444.333222111 kNAV
	// Satoshi to NAV: 444333.222111 NAV
	// Satoshi to MilliNAV: 444333222.111 mNAV
	// Satoshi to MicroNAV: 444333222111 μNAV
	// Satoshi to Satoshi: 44433322211100 Satoshi
}
