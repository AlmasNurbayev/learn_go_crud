package auth_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/auth"
	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {

	res_sign, err := auth.GenerateJWT("testKey", 1, "sample@sample.com")
	if err != nil {
		assert.FailNow(t, "generate token error")
	}
	//fmt.Fprintln(os.Stdout, "res_sign", res_sign)
	fmt.Fprintln(os.Stdout, "err", err)
	res_parse, err := auth.VerifyJWT(res_sign, "testKey")

	if err != nil {
		assert.FailNow(t, "parsetoken error")
	}
	//fmt.Fprintln(os.Stdout, "=========")
	//fmt.Fprintln(os.Stdout, "res_parse", res_parse)

	assert.Equal(t, res_parse.Id, 1)
	assert.Equal(t, res_parse.Email, "sample@sample.com")

}
