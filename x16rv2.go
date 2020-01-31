package x16rv2

/*
void x16rv2_hash(const char* input, char* output);
#cgo  LDFLAGS: libx16rv2.so
 */
import "C"
import (
	"unsafe"
)

func HashX16RV2(b []byte) [32]byte {
	var blockHash [32]byte
	C.x16rv2_hash((*C.char)(unsafe.Pointer(&b[0])), (*C.char)(unsafe.Pointer(&blockHash[0])))
	return blockHash
}
