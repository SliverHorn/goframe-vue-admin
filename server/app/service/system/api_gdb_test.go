package service

import (
	"gf-vue-admin/server/app/api/request"
	"testing"
)

func TestApi_Create(t *testing.T) {
	infos := []*request.CreateApi{
		{BaseApi: request.BaseApi{Path: "1", Method: "1", ApiGroup: "1", Description: "1"}},
		{BaseApi: request.BaseApi{Path: "2", Method: "2", ApiGroup: "2", Description: "2"}},
		{BaseApi: request.BaseApi{Path: "3", Method: "3", ApiGroup: "3", Description: "3"}},
		{BaseApi: request.BaseApi{Path: "4", Method: "4", ApiGroup: "4", Description: "4"}},
	}
	for _, info := range infos {
		if err := Api.Create(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestApi_First(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
	}
	for _, info := range infos {
		if result, err := Api.First(info); err != nil {
			t.Error(err)
		} else {
			t.Log(result)
			t.Log("success!")
		}
	}
}

func TestApi_Update(t *testing.T) {
	infos := []*request.UpdateApi{
		{GetById: request.GetById{Id: 1}, BaseApi: request.BaseApi{Path: "2", Method: "2", ApiGroup: "2", Description: "2"}},
		{GetById: request.GetById{Id: 2}, BaseApi: request.BaseApi{Path: "3", Method: "3", ApiGroup: "3", Description: "3"}},
		{GetById: request.GetById{Id: 3}, BaseApi: request.BaseApi{Path: "4", Method: "4", ApiGroup: "4", Description: "4"}},
		{GetById: request.GetById{Id: 4}, BaseApi: request.BaseApi{Path: "5", Method: "5", ApiGroup: "5", Description: "5"}},
	}
	for _, info := range infos {
		if err := Api.Update(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestApi_Delete(t *testing.T) {
	infos := []*request.DeleteApi{
		{Path: "1", Method: "1", GetById: request.GetById{Id: 1}},
		{Path: "2", Method: "2", GetById: request.GetById{Id: 2}},
		{Path: "3", Method: "3", GetById: request.GetById{Id: 3}},
		{Path: "4", Method: "4", GetById: request.GetById{Id: 4}},
	}
	for _, info := range infos {
		if err := Api.Delete(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestApi_GetList(t *testing.T) {
	infos := []*request.SearchApi{
		{Path: "", Description: "", ApiGroup: "", Method: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "1", Description: "", ApiGroup: "", Method: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Description: "2", ApiGroup: "", Method: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Description: "", ApiGroup: "3", Method: "", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Description: "", ApiGroup: "", Method: "4", PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
	}
	for _, info := range infos {
		list, total, err := Api.GetList(info)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(list)
			t.Log(total)
			t.Log("success!")
		}
	}
}

func TestApi_GetAllApi(t *testing.T) {
	if results, err := Api.GetAllApi(); err != nil {
		t.Error(err)
	} else {
		t.Log(results)
		t.Log("success!")
	}
}
