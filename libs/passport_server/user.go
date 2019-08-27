/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-11-26 16:26:57
# File Name: user.go
# Description:
####################################################################### */

package passport_server

import (
	"context"
	"fmt"

	"github.com/smallnest/rpcx/client"
	"gitlab.com/feichi/fcad_thrift/libs/go/enums"
	services "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_services"
	types "gitlab.com/feichi/fcad_thrift/libs/go/fcmp_passport_types"
)

func GetUser(token string) (r *types.MpUser, err error) {
	req := &services.GetUserRequest{Header: BuildCommonHeaderWithToken(token)}
	resp := &services.GetUserResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetUser", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	r = resp.Body
	return
}

func GetUserById(id int32) (r *types.MpUser, err error) {
	req := &services.GetUserByIdsRequest{Header: BuildCommonHeader(), Body: []int32{id}}
	resp := &services.GetUserByIdsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetUserByIds", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	users := resp.Body
	for _, v := range users {
		r = v
	}

	return
}

func GetUserByIds(ids []int32) (r map[int32]*types.MpUser, err error) {
	req := &services.GetUserByIdsRequest{Header: BuildCommonHeader(), Body: ids}
	resp := &services.GetUserByIdsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetUserByIds", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	r = resp.Body
	return
}

func Login(token string, user *types.MpUser) (r *types.MpUser, rtoken string, err error) {
	req := &services.LoginRequest{Header: BuildCommonHeaderWithToken(token), Body: user}
	resp := &services.LoginResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "Login", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	r = resp.Body
	rtoken = resp.Header.Metadata["token"]
	return
}

func Logout(token string) (err error) {
	req := &services.LogoutRequest{Header: BuildCommonHeaderWithToken(token)}
	resp := &services.LogoutResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "Logout", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	return
}

func UpdateUser(token string, user *types.MpUser) (err error) {
	req := &services.UpdateUserRequest{Header: BuildCommonHeaderWithToken(token), Body: user}
	resp := &services.UpdateUserResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "UpdateUser", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	return

}

func CreateFormid(formid *types.MpUserFormid) (err error) {
	req := &services.CreateUserFormidRequest{Header: BuildCommonHeader(), Body: formid}
	resp := &services.CreateUserFormidResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "CreateUserFormid", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	return
}

func GetFormidByUserId(id int32) (r string, err error) {
	req := &services.GetUserFormidByUidsRequest{Header: BuildCommonHeader(), Body: []int32{id}}
	resp := &services.GetUserFormidByUidsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetUserFormidByUids", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	formIds := resp.Body
	for _, v := range formIds {
		r = v.Formid
	}
	return
}

func GetFormidByUserIds(ids []int32) (r map[int32]string, err error) {
	req := &services.GetUserFormidByUidsRequest{Header: BuildCommonHeader(), Body: ids}
	resp := &services.GetUserFormidByUidsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetUserFormidByUids", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	r = map[int32]string{}
	formIds := resp.Body
	for _, v := range formIds {
		r[v.UserId] = v.Formid
	}
	return
}
