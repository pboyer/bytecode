// Code generated by "stringer -type=binOp"; DO NOT EDIT

package main

import "fmt"

const _binOp_name = "ADDSUBDIVMUL"

var _binOp_index = [...]uint8{0, 3, 6, 9, 12}

func (i binOp) String() string {
	if i < 0 || i >= binOp(len(_binOp_index)-1) {
		return fmt.Sprintf("binOp(%d)", i)
	}
	return _binOp_name[_binOp_index[i]:_binOp_index[i+1]]
}