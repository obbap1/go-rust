package main

// #cgo LDFLAGS: -L/Users/pbaba/projects/go-rust/r/target/debug -lr
// #include <bindings.h>
import "C"
import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("check.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	a := C.CString("merry christmas")
	b := C.hello_world(a)
	defer C.hello_world_free(b)
	// defer C.free(unsafe.Pointer(b))
	fmt.Println(C.GoString(b))
	pprof.WriteHeapProfile(f)
}
