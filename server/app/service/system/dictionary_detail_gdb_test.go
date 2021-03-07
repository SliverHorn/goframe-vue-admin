package service

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"testing"
)

func TestDetail_Create(t *testing.T) {
	infos := []*request.CreateDictionaryDetail{
		{BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "1", Status: _true, Value: 1, Sort: 1, DictionaryId: 1}},
		{BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "2", Status: _false, Value: 2, Sort: 2, DictionaryId: 2}},
		{BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "3", Status: _true, Value: 3, Sort: 3, DictionaryId: 3}},
		{BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "4", Status: _false, Value: 4, Sort: 4, DictionaryId: 4}},
	}
	for _, info := range infos {
		if err := DictionaryDetail.Create(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestDetail_First(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
	}
	for _, info := range infos {
		if result, err := DictionaryDetail.First(info); err != nil {
			t.Error(err)
		} else {
			t.Log(result)
			t.Log("success!")
		}
	}
}

func TestDetail_Update(t *testing.T) {
	infos := []*request.UpdateDictionaryDetail{
		{GetById: request.GetById{Id: 1}, BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "2", Status: _false, Value: 2, Sort: 2, DictionaryId: 2}},
		{GetById: request.GetById{Id: 2}, BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "3", Status: _true, Value: 3, Sort: 3, DictionaryId: 3}},
		{GetById: request.GetById{Id: 3}, BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "3", Status: _false, Value: 4, Sort: 4, DictionaryId: 4}},
		{GetById: request.GetById{Id: 4}, BaseDictionaryDetail: request.BaseDictionaryDetail{Label: "5", Status: _false, Value: 5, Sort: 5, DictionaryId: 5}},
	}
	for _, info := range infos {
		if err := DictionaryDetail.Update(info); err != nil {
			t.Error(err)
		} else {
			a := &request.GetById{Id: info.Id}
			if result, err := DictionaryDetail.First(a); err != nil {
				t.Error(err)
			} else {
				t.Log(result)
				t.Log("success!")
			}
		}
	}
}

func TestDetail_Delete(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
	}
	for _, info := range infos {
		if err := DictionaryDetail.Delete(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestDetail_GetList(t *testing.T) {
	infos := []*request.SearchDictionaryDetail{
		{Label: "", Status: nil, Value: 0, Sort: 0, DictionaryId: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Label: "2", Status: nil, Value: 0, Sort: 0, DictionaryId: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Label: "", Status: _true, Value: 0, Sort: 0, DictionaryId: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Label: "", Status: nil, Value: 3, Sort: 0, DictionaryId: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Label: "", Status: nil, Value: 0, Sort: 4, DictionaryId: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Label: "", Status: nil, Value: 0, Sort: 0, DictionaryId: 5, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
	}
	for _, info := range infos {
		list, total, err := DictionaryDetail.GetList(info)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(list)
			t.Log(total)
			t.Log("success!")
		}
	}
}
