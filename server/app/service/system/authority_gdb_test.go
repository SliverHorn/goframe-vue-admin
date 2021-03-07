package service

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"testing"
)

func TestAuthority_Create(t *testing.T) {
	infos := []*request.CreateAuthority{
		{BaseAuthority: request.BaseAuthority{ParentId: "0", AuthorityId: "1", AuthorityName: "1", DefaultRouter: "1"}},
		{BaseAuthority: request.BaseAuthority{ParentId: "0", AuthorityId: "2", AuthorityName: "2", DefaultRouter: "2"}},
		{BaseAuthority: request.BaseAuthority{ParentId: "0", AuthorityId: "3", AuthorityName: "3", DefaultRouter: "3"}},
		{BaseAuthority: request.BaseAuthority{ParentId: "0", AuthorityId: "4", AuthorityName: "4", DefaultRouter: "4"}},
	}
	for _, info := range infos {
		if err := Authority.Create(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}
