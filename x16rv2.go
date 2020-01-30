package x16rv2

/*
void x16rv2_hash(const char* input, char* output);
#cgo  LDFLAGS: libx16rv2.so
 */
import "C"
import (
	"unsafe"
)

func X16RV2(data []byte) []byte {
	output := make([]byte, 32)
	C.x16rv2_hash((*C.char)(unsafe.Pointer(&data[0])), (*C.char)(unsafe.Pointer(&output[0])))
	return output
}
