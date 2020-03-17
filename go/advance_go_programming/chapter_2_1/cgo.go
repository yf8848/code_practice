package main

//void SayHello(const char* str);
//void addNum(int num);
import "C"

func main() {
	C.SayHello(C.CString("hello cgo\n"))
	C.addNum(4)
}
