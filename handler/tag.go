package handler

import (
	m "blog/model"
	u "blog/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetTags : to get multiple tags data
func GetTags(w http.ResponseWriter, r *http.Request) {
	var (
		retDatas   []*m.Tag
		err        error
		statusCode int
	)

	tag := r.URL.Query().Get("tag")

	retDatas, statusCode, err = bh.GetTags(tag)
	if err != nil {
		u.ReturnResponse(w, statusCode, err.Error(), nil)
		return
	}

	if len(retDatas) == 0 {
		u.ReturnResponse(w, statusCode, u.DataNotExist("tag"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, "", retDatas)
}

// GetTag : to get data of tag
func GetTag(w http.ResponseWriter, r *http.Request) {
	var (
		retData    *m.Tag
		err        error
		statusCode int
		id         string
	)

	id = mux.Vars(r)["id"]

	retData, statusCode, err = bh.GetTagById(id)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.DataNotExist("tag"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, "", retData)
}

// AddTag : to add new data of tag
func AddTag(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *m.ReqDataTag
		statusCode int
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		u.ReturnResponse(w, http.StatusBadRequest, u.ErrorFailedReadRequest(), nil)
		return
	}

	statusCode, err = bh.AddTag(reqData)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("insert", "tag"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("insert", "tag"), nil)
}

// ModifyTag : to change existing tag data
func ModifyTag(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *m.ReqDataTag
		statusCode int
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		u.ReturnResponse(w, http.StatusBadRequest, u.ErrorFailedReadRequest(), nil)
		return
	}

	reqData.Id = mux.Vars(r)["id"]
	statusCode, err = bh.ModifyTag(reqData)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("update", "tag"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("update", "tag"), nil)
}

// RemoveTag : to remove tag data
func RemoveTag(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		statusCode int
		id         string
	)

	id = mux.Vars(r)["id"]
	statusCode, err = bh.RemoveTag(id)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("delete", "tag"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("delete", "tag"), nil)
}
