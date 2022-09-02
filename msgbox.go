package msgbox

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	// 对话框类型常量
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	//默认按钮类型常量
	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300

	//更多请查看windows api 网址：https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-messageboxw
)

// 立即加载模式
func MessageBox(caption, text string, style uintptr) int {

	user32, _ := syscall.LoadLibrary("user32.dll")
	messageBox, _ := syscall.GetProcAddress(user32, "MessageBoxW")
	defer syscall.FreeLibrary(user32)

	t, _ := syscall.UTF16PtrFromString(text)
	c, _ := syscall.UTF16PtrFromString(caption)
	ret, _, callErr := syscall.SyscallN(
		messageBox,
		0, // api参数1，父窗口句柄，为0无父窗口
		//api参数2，对话框内容文本，utf16指针格式文本
		uintptr(unsafe.Pointer(t)),
		//api参数3，对话框标题，utf16指针格式文本
		uintptr(unsafe.Pointer(c)),
		style, //api参数4，对话框类型
	)
	if callErr != 0 {
		panic(fmt.Sprintf("%s failed: %v", "Call MessageBox", callErr))
	}
	return int(ret)
}

// 懒加载模式
func MsgBox(caption, text string, style uintptr) int {

	user32 := syscall.NewLazyDLL("user32.dll")
	msgBox := user32.NewProc("MessageBoxW")

	t, _ := syscall.UTF16PtrFromString(text)
	c, _ := syscall.UTF16PtrFromString(caption)
	ret, _, _ := msgBox.Call(
		0, // api参数1，父窗口句柄，为0无父窗口
		//api参数2，对话框内容文本，utf16指针格式文本
		uintptr(unsafe.Pointer(t)),
		//api参数3，对话框标题，utf16指针格式文本
		uintptr(unsafe.Pointer(c)),
		style, //api参数4，对话框类型
	)
	return int(ret)
}
