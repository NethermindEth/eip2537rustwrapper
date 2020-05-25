package main

// #cgo CFLAGS: -g -Wall
// #include "eip2537_header.h"
// #cgo LDFLAGS: -L${SRCDIR}/libs -leth_pairings -lm -ldl
import "C"
import "errors"

func call(operation int, input []byte) ([]byte, error) {
	iLen := len(input)
	iPtr := C.CBytes(input)
	defer C.free(iPtr)

	oPtr := C.malloc(C.sizeof_char * 256)
	defer C.free(oPtr)
	outLen := uint32(0)
	ePtr := C.malloc(C.sizeof_char * 256)
	defer C.free(ePtr)
	errLen := uint32(0)

	op := C.char(operation)

	cR := C.eip2537_perform_operation(op, (*C.char)(iPtr), C.uint(iLen), (*C.char)(oPtr), (*C.uint)(&outLen), (*C.char)(ePtr), (*C.uint)(&errLen))
	result := int(cR)

	if result == 0 {
		if outLen == 0 {
			return nil, errors.New("function returned zero length")
		}

		output := C.GoBytes(oPtr, C.int(outLen))

		return output, nil
	}

	if errLen == 0 {
		return nil, errors.New("function returned error with no description")
	}

	descr := C.GoStringN((*C.char)(ePtr), C.int(errLen))

	return nil, errors.New(descr)
}

func G1Add(input []byte) ([]byte, error) {
	return call(1, input)
}

func G1Mul(input []byte) ([]byte, error) {
	return call(2, input)
}

func G1Multiexp(input []byte) ([]byte, error) {
	return call(3, input)
}

func G2Add(input []byte) ([]byte, error) {
	return call(4, input)
}

func G2Mul(input []byte) ([]byte, error) {
	return call(5, input)
}

func G2Multiexp(input []byte) ([]byte, error) {
	return call(6, input)
}

func Pair(input []byte) ([]byte, error) {
	return call(7, input)
}

func FpToG1(input []byte) ([]byte, error) {
	return call(8, input)
}

func Fp2ToG2(input []byte) ([]byte, error) {
	return call(9, input)
}
func main() {}
