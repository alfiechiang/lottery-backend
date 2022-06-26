package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	Cost int
}

//Make 加密方法
func (b *Bcrypt) Make(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, b.Cost)
}

//Check 检查方法
func Check(hashedPassword, password string) error {

	b1 := []byte(hashedPassword)
	b2 := []byte(password)

	return bcrypt.CompareHashAndPassword(b1, b2)
}

func HashMake(password string) string {
	var b Bcrypt
	pwd, _ := b.Make([]byte(password))
	pwd1 := string(pwd)
	return pwd1
}
