package main

// #cgo windows LDFLAGS: -lnidaqmx
// #cgo windows CFLAGS: -I'C:/NI/ExternalCompilerSupport/C/include'
// #cgo windows LDFLAGS: "-LC:/NI/ExternalCompilerSupport/C/lib32/msvc"
// #include <stdio.h>
// #include "NIDAQmx.h"
import "C"

import "fmt"

func main() {
	var version C.uInt32
	C.DAQmxGetSysNIDAQMajorVersion(&version)
	fmt.Print(version)
}


