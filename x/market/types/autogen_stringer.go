// Code generated by "stringer -linecomment -output=autogen_stringer.go -type=OrderState,BidState,LeaseState"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OrderOpen-1]
	_ = x[OrderMatched-2]
	_ = x[OrderClosed-3]
}

const _OrderState_name = "ordermatchedclosed"

var _OrderState_index = [...]uint8{0, 5, 12, 18}

func (i OrderState) String() string {
	i -= 1
	if i >= OrderState(len(_OrderState_index)-1) {
		return "OrderState(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _OrderState_name[_OrderState_index[i]:_OrderState_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BidOpen-1]
	_ = x[BidMatched-2]
	_ = x[BidLost-3]
	_ = x[BidClosed-4]
}

const _BidState_name = "openmatchedlostclosed"

var _BidState_index = [...]uint8{0, 4, 11, 15, 21}

func (i BidState) String() string {
	i -= 1
	if i >= BidState(len(_BidState_index)-1) {
		return "BidState(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BidState_name[_BidState_index[i]:_BidState_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LeaseActive-1]
	_ = x[LeaseInsufficientFunds-2]
	_ = x[LeaseClosed-3]
}

const _LeaseState_name = "activeinsufficient-fundsclosed"

var _LeaseState_index = [...]uint8{0, 6, 24, 30}

func (i LeaseState) String() string {
	i -= 1
	if i >= LeaseState(len(_LeaseState_index)-1) {
		return "LeaseState(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _LeaseState_name[_LeaseState_index[i]:_LeaseState_index[i+1]]
}
