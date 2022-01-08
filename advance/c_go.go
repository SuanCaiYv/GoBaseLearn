package advance

/*
void println0(const char *str);
*/
import "C"

func CGo1() {
	C.println0(C.CString("aaa"))
}
