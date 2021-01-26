// helper_test.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2021 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package rpofs //nolint:testpackage

import (
	"testing"

	"github.com/vault-thirteen/tester"
)

func Test_isSetOfUniqueSymbols(t *testing.T) {
	type TestData struct {
		Set            []rune
		ResultExpected bool
	}

	tests := []TestData{}

	// Test #1. Empty Set.
	tests = append(tests, TestData{
		Set:            []rune{},
		ResultExpected: false,
	})

	// Test #2. Set with a single Item
	tests = append(tests, TestData{
		Set:            []rune{'A'},
		ResultExpected: true,
	})

	// Test #3. Set with unique Symbols
	tests = append(tests, TestData{
		Set:            []rune{'J', 'A', 'C', 'K'},
		ResultExpected: true,
	})

	// Test #4. Set with non-unique Symbols
	tests = append(tests, TestData{
		Set:            []rune{'J', 'A', 'C', 'K', 'C'},
		ResultExpected: false,
	})

	var (
		aTest  = tester.New(t)
		result bool
	)

	for _, test := range tests {
		result = isSetOfUniqueSymbols(test.Set)
		aTest.MustBeEqual(result, test.ResultExpected)
	}
}
