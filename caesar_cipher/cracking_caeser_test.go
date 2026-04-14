package main

import (
	"fmt"
	"testing"
)

func TestCrackCipher(t *testing.T) {
	type testCase struct {
		ciphertext string
		expected   string
	}

	runCases := []testCase{
		{"bcdefghij", "abcdefghi"},
		{"gurdhvpxoebjasbkwhzcfbiregurynmlqbt", "thequickbrownfoxjumpsoverthelazydog"},
		{"guvfvfnybatrefragraprsbezberqngn", "thisisalongersentenceformoredata"},
	}

	submitCases := append(runCases, []testCase{
		{"nynetrrabhtupvcuregrkgznxrfoernxvatpnrfnepvcuregevivny", "alargeenoughciphertextmakesbreakingcaesarciphertrivial"},
		{"serdhraplnanylfvfvfgurjnlgbtb", "frequencyanalysisisthewaytogo"},
		{"sehhusjxehiurqjjuhoijqfbu", "correcthorsebatterystaple"},
		{"nmdsvnsgqddentqehudrhwrdudmdhfgsmhmdsdm", "onetwothreefourfivesixseveneightnineten"},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		actual := crack(test.ciphertext)

		if actual != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      ciphertext: %v
Expecting:   plaintext: %v
Actual:      plaintext: %v
Fail
`, test.ciphertext, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      ciphertext: %v
Expecting:   plaintext: %v
Actual:      plaintext: %v
Pass
`, test.ciphertext, test.expected, actual)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

var withSubmit = true
