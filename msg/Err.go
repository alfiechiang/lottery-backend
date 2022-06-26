package msg

var ErrMsg map[int]string

func Init() {

	msg := make(map[int]string)
	ErrMsg = msg

	ErrMsg[10000] = "帳號錯誤"
	ErrMsg[10001] = "密碼錯誤"

}
