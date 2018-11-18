package jwttoken

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/robbert229/jwt"
)

func JWT() {
	secret := "ThisIsMySuperSecret"
	algorithm := jwt.HmacSha256(secret)

	id := 3610300032039
	str := strconv.Itoa(id)
	md5 := md5.New()
	io.WriteString(md5, str)
	encrypted := string(md5.Sum(nil))
	fmt.Printf("%v", encrypted)
	claims := jwt.NewClaim()
	claims.Set("Role", "Admin")

	claims.Set(encrypted, "101231")
	claims.SetTime("exp", time.Now().Add(time.Minute))

	token, err := algorithm.Encode(claims)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Token: %s\n", token)

	if algorithm.Validate(token) != nil {
		panic(err)
	}

	loadedClaims, err := algorithm.Decode(token)
	if err != nil {
		panic(err)
	}

	role, err := loadedClaims.Get("Role")
	if err != nil {
		panic(err)
	}

	roleString, ok := role.(string)
	if !ok {
		panic(err)
	}

	if strings.Compare(roleString, "Admin") == 0 {
		//user is an admin
		fmt.Println("User is an admin")
	}
}
