package handler

import (
	m "blog/model"
	u "blog/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPosts : to get multiple posts data
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var (
		retDatas   []*m.Post
		err        error
		statusCode int
	)

	title := r.URL.Query().Get("title")
	content := r.URL.Query().Get("content")
	status := r.URL.Query().Get("status")
	publishedDate := r.URL.Query().Get("published_date")
	tag := r.URL.Query().Get("tag")

	retDatas, statusCode, err = bh.GetPosts(title, content, status, publishedDate, tag)
	if err != nil {
		u.ReturnResponse(w, statusCode, err.Error(), nil)
		return
	}

	if len(retDatas) == 0 {
		u.ReturnResponse(w, statusCode, u.DataNotExist("post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, "", retDatas)
}

// GetPost : to get data of post
func GetPost(w http.ResponseWriter, r *http.Request) {
	var (
		retData    *m.Post
		err        error
		statusCode int
		id         string
	)

	id = mux.Vars(r)["id"]

	retData, statusCode, err = bh.GetPost(id)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.DataNotExist("post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, "", retData)
}

// AddPost : to add new data of post
func AddPost(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *m.ReqDataPost
		statusCode int
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		u.ReturnResponse(w, http.StatusBadRequest, u.ErrorFailedReadRequest(), nil)
		return
	}

	statusCode, err = bh.AddPost(reqData)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("insert", "post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("insert", "post"), nil)
}

// ModifyPost : to change existing post data
func ModifyPost(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		reqData    *m.ReqDataPost
		statusCode int
	)

	// parse json from request body
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		u.ReturnResponse(w, http.StatusBadRequest, u.ErrorFailedReadRequest(), nil)
		return
	}

	reqData.Id = mux.Vars(r)["id"]
	statusCode, err = bh.ModifyPost(reqData)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("update", "post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("update", "post"), nil)
}

// RemovePost : to remove post data
func RemovePost(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		statusCode int
		id         string
	)

	id = mux.Vars(r)["id"]
	statusCode, err = bh.RemovePost(id)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("delete", "post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("delete", "post"), nil)
}

// PublishPost : to change status post data from draft into publish
func PublishPost(w http.ResponseWriter, r *http.Request) {
	var (
		err        error
		statusCode int
		id         string
	)

	id = mux.Vars(r)["id"]
	statusCode, err = bh.PublishPost(id)
	if err != nil {
		u.ReturnResponse(w, statusCode, u.ErrorFailedExecData("publish", "post"), nil)
		return
	}

	u.ReturnResponse(w, statusCode, u.SuccessExecData("publish", "post"), nil)
}
