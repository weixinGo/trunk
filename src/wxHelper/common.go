package wxHelper 

import(
    "strings"
)


func checkResponse(msg string, respons string) error{
	if strings.Index(respons, "errcode") != -1 {
		return errors.New(fmt.Sprintf("%v，返回内容：%v", msg, respons))
	}
	return nil 
}