package main

import rand2 "math/rand"
// 生成测试所用的随机数据，包括随机的行数，随机的username 和随机的information

var (
	strConst =  map[int]string{1:"root", 2:"user", 3:"guest"}
	// 数据库中相应的username字段的三个类型，都是用在测试场所。
	// 实际的函数参数应该考虑运用接口的方式进行扩展，使得程序的可扩展性更强
)

// randomInt Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand2.Intn(max-min)
}

// randomString Generate a random string of Ascii(21 - 126) chars with len
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(21, 126))
	}
	return string(bytes)
}

// randInfo Generate info randomly
func randInfo() string {

	str := randomString(randomInt(3, 100))
	return str
}

// randUserName Generate user_name randomly
func randUser() string {
	str := strConst[randomInt(1, 4)]
	return str
}