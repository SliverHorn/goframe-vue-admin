package service

import "gf-vue-admin/server/library/utils"

var Email = new(email)

type email struct{}

func (e *email) Test() error {
	subject := "test"
	body := "test"
	return utils.Email.Test(subject, body)
}
