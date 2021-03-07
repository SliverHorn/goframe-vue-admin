package service

import (
	"gf-vue-admin/server/app/api/request"
	"testing"
)

var (
	_true  *bool
	_false *bool
)

func init() {
	_true = new(bool)
	*_true = true
	_false = new(bool)
	*_false = false
}

func TestDictionary_Create(t *testing.T) {
	var infos = []*request.CreateDictionary{
		{BaseDictionary: request.BaseDictionary{Name: "1", Type: "1", Desc: "1", Status: _true}},
		{BaseDictionary: request.BaseDictionary{Name: "2", Type: "2", Desc: "2", Status: _false}},
		{BaseDictionary: request.BaseDictionary{Name: "3", Type: "3", Desc: "3", Status: _true}},
		{BaseDictionary: request.BaseDictionary{Name: "4", Type: "4", Desc: "4", Status: _false}},
	}
	for _, info := range infos {
		if err := Dictionary.Create(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestDictionary_First(t *testing.T) {
	infos := []*request.FirstDictionary{
		{Id: 1, Type: ""},
		{Id: 2, Type: ""},
		{Id: 3, Type: ""},
		{Id: 4, Type: ""},
	}
	for _, info := range infos {
		if result, err := Dictionary.First(info); err != nil {
			t.Error(err)
		} else {
			t.Log(result)
			t.Log("success!")
		}
	}
}

func TestDictionary_Update(t *testing.T) {
	infos := []*request.UpdateDictionary{
		{GetById: request.GetById{Id: 1}, BaseDictionary: request.BaseDictionary{Name: "2", Type: "2", Desc: "2", Status: _false}},
		{GetById: request.GetById{Id: 2}, BaseDictionary: request.BaseDictionary{Name: "3", Type: "3", Desc: "3", Status: _true}},
		{GetById: request.GetById{Id: 3}, BaseDictionary: request.BaseDictionary{Name: "4", Type: "4", Desc: "4", Status: _true}},
	}
	for _, info := range infos {
		if err := Dictionary.Update(info); err != nil {
			t.Error(err)
		} else {
			a := &request.FirstDictionary{Id: info.Id}
			if result, err := Dictionary.First(a); err != nil {
				t.Error(err)
			} else {
				t.Log(result)
				t.Log("success!")
			}
		}
	}
}

func TestDictionary_Delete(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
	}
	for _, info := range infos {
		if err := Dictionary.Delete(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestDictionary_GetList(t *testing.T) {
	var infos = []*request.SearchDictionary{
		{Status: _false, Name: "", Type: "", Desc: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Status: nil, Name: "1", Type: "", Desc: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Status: nil, Name: "", Type: "2", Desc: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Status: nil, Name: "", Type: "", Desc: "3", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Status: nil, Name: "", Type: "", Desc: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
	}
	for _, info := range infos {
		list, total, err := Dictionary.GetList(info)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(total)
			t.Log(list)
			t.Log("success!")
		}
	}
}
