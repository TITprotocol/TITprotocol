// Copyright 2017 The go-TITprotocol Authors
// This file is part of the go-TITprotocol library.
//
// The go-TITprotocol library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-TITprotocol library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-TITprotocol library. If not, see <http://www.gnu.org/licenses/>.

package math

import (
	"math/big"
	"testing"
	a "math"
	"fmt"
)

// BigPow returns a ** b as a big integer.
func bigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

func Test_aaaa(t *testing.T) {
	var (
		baseBlockRewardfirstYear  = big.NewInt(0.4243e+17) // 550000/30.0/43200=0.4243827160493827
		baseBlockRewardSecondYear = big.NewInt(12.11e+16)  // 550000/30.0/43200=0.4243827160493827
		baseBlockRewardThirdYear  = big.NewInt(34.05e+16)  // 550000/30.0/43200=0.4243827160493827
		baseBlockRewardForthYear  = big.NewInt(85.74e+16)  // 550000/30.0/43200=0.4243827160493827
		baseBlockRewardFivthYear  = big.NewInt(193.1e+16)  // 550000/30.0/43200=0.4243827160493827
		baseBlockRewardSixthYear  = big.NewInt(388.6e+16)  // 550000/30.0/43200=0.4243827160493827
	)
	h := big.NewInt(64800000)
	ret := new(big.Int).Div(h, big.NewInt(1296000))
	fmt.Println("header.Number is ", h, "cyle is ", ret.String())
	if ret == big.NewInt(0) {
		fmt.Println("baseBlockRewardfirstYear.Number is ", baseBlockRewardfirstYear)
		}

	retInt64 := ret.Int64()
	ab := big.NewInt(0)
	if retInt64 > 0 {
		if retInt64 < 12 { // first year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardfirstYear, big.NewInt(10)), bigPow(int64(a.Pow(110, float64(ret.Int64()))), ret.Int64()-3))
		} else if ((retInt64 >= 12) && (retInt64 < 24)) { // second year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardSecondYear, big.NewInt(10)), bigPow(int64(a.Pow(109, float64(ret.Int64()-12))), ret.Int64()-15))
		} else if ((retInt64 >= 24) && (retInt64 < 36)) { // third year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardThirdYear, big.NewInt(10)), bigPow(int64(a.Pow(108, float64(ret.Int64()-24))), ret.Int64()-27))
		} else if ((retInt64 >= 36) && (retInt64 < 48)) { // fourth year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardForthYear, big.NewInt(10)), bigPow(int64(a.Pow(107, float64(ret.Int64()-36))), ret.Int64()-39))
		} else if ((retInt64 >= 48) && (retInt64 < 60)) { // fivth year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardFivthYear, big.NewInt(10)), bigPow(int64(a.Pow(106, float64(ret.Int64()-48))), ret.Int64()-51))
		} else if ((retInt64 >= 60) && (retInt64 < 72)) { // sixth year
			ab = new(big.Int).Mul(new(big.Int).Mul(baseBlockRewardSixthYear, big.NewInt(10)), bigPow(int64(a.Pow(105, float64(ret.Int64()-60))), ret.Int64()-63))
		} else {
			ab = baseBlockRewardfirstYear
		}
	}
	t.Errorf("a := %v\n", ab.String())
}
