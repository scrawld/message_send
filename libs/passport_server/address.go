/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-11-26 16:31:36
# File Name: address.go
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

func GetAddress(userId int32) (r []*types.MpUserAddress, err error) {
	req := &services.SearchUserAddressByParamsRequest{Header: BuildCommonHeader(),
		Body: &services.SearchUserAddressParams{UserId: userId}}
	resp := &services.SearchUserAddressByParamsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "SearchUserAddress", req, resp); err != nil {
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

func CreateAddress(address *types.MpUserAddress) (r int32, err error) {
	req := &services.CreateUserAddressRequest{Header: BuildCommonHeader(), Body: address}
	resp := &services.CreateUserAddressResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "CreateUserAddress", req, resp); err != nil {
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

func UpdateAddress(address *types.MpUserAddress) (err error) {
	req := &services.UpdateUserAddressRequest{Header: BuildCommonHeader(), Body: address}
	resp := &services.UpdateUserAddressResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "UpdateUserAddress", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	return
}
