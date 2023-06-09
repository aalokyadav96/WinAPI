package main


import (

	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
    user32DLL				= windows.NewLazyDLL("user32.dll")
    procSystemParamInfo	= user32DLL.NewProc("SystemParametersInfoW")
)

func main()  {
	imagePath, _ := windows.UTF16PtrFromString(`F:\wallp\aper\11.jpg`)
	fmt.Println("[+] Changing background now...")
	procSystemParamInfo.Call(20, 0, uintptr(unsafe.Pointer(imagePath)), 0x001A)

}