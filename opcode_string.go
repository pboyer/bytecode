// Code generated by "stringer -type=opCode"; DO NOT EDIT

package main

import "fmt"

const _opCode_name = "PRINTPUSHPOPRETSTOLOADBIN_OPPUSH_IPPUSH_FPJMPHALT"

var _opCode_index = [...]uint8{0, 5, 9, 12, 15, 18, 22, 28, 35, 42, 45, 49}

func (i opCode) String() string {
	if i < 0 || i >= opCode(len(_opCode_index)-1) {
		return fmt.Sprintf("opCode(%d)", i)
	}
	return _opCode_name[_opCode_index[i]:_opCode_index[i+1]]
}