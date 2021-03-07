package service

import (
	"github.com/flipped-aurora/gf-vue-admin/server/app/api/request"
	"testing"
)

func TestRecord_Create(t *testing.T) {
	infos := []*request.CreateOperationRecord{
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "1", Path: "1", Agent: "1", Method: "1", Request: "1", Response: "1", ErrorMessage: "1", Status: 1, UserID: 1, Latency: 1}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "2", Path: "2", Agent: "2", Method: "2", Request: "2", Response: "2", ErrorMessage: "2", Status: 2, UserID: 2, Latency: 2}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "3", Path: "3", Agent: "3", Method: "3", Request: "3", Response: "3", ErrorMessage: "3", Status: 3, UserID: 3, Latency: 3}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "4", Path: "4", Agent: "4", Method: "4", Request: "4", Response: "4", ErrorMessage: "4", Status: 4, UserID: 4, Latency: 4}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "5", Path: "5", Agent: "5", Method: "5", Request: "5", Response: "5", ErrorMessage: "5", Status: 5, UserID: 5, Latency: 5}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "6", Path: "6", Agent: "6", Method: "6", Request: "6", Response: "6", ErrorMessage: "6", Status: 6, UserID: 6, Latency: 6}},
		{BaseOperationRecord: request.BaseOperationRecord{Ip: "7", Path: "7", Agent: "7", Method: "7", Request: "7", Response: "7", ErrorMessage: "7", Status: 7, UserID: 7, Latency: 7}},
	}
	for _, info := range infos {
		if err := OperationRecord.Create(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestRecord_First(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
		{Id: 5},
		{Id: 6},
		{Id: 7},
	}
	for _, info := range infos {
		if result, err := OperationRecord.First(info); err != nil {
			t.Error(err)
		} else {
			t.Log(result)
			t.Log("success!")
		}
	}
}

func TestRecord_Delete(t *testing.T) {
	infos := []*request.GetById{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
		{Id: 5},
		{Id: 6},
		{Id: 7},
	}
	for _, info := range infos {
		if err := OperationRecord.Delete(info); err != nil {
			t.Error(err)
		} else {
			t.Log("success!")
		}
	}
}

func TestRecord_Deletes(t *testing.T) {
	info := &request.GetByIds{Ids: []int{1, 2, 3, 4, 5, 6, 7}}
	if err := OperationRecord.Deletes(info); err != nil {
		t.Error(err)
	} else {
		t.Log("success!")
	}
}

func TestRecord_GetList(t *testing.T) {
	infos := []*request.SearchOperationRecord{
		{Path: "", Method: "", Status: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "1", Method: "", Status: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Method: "2", Status: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Method: "", Status: 3, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "4", Method: "", Status: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Method: "5", Status: 0, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Method: "", Status: 6, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
		{Path: "", Method: "", Status: 7, PageInfo: request.PageInfo{Page: 1, PageSize: 10}},
	}
	for _, info := range infos {
		list, total, err := OperationRecord.GetList(info)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(list)
			t.Log(total)
			t.Log("success!")
		}
	}
}
