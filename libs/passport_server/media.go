/* ######################################################################
# Author: (zhengfei@fcadx.cn)
# Created Time: 2018-11-26 16:28:29
# File Name: media.go
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

func SearchMediaByParams(body *services.SearchMediaParams) (r []*types.Media, err error) {
	req := &services.SearchMediaByParamsRequest{Header: BuildCommonHeader(),
		Body: body}
	resp := &services.SearchMediaByParamsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "SearchMediaByParams", req, resp); err != nil {
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

func GetMediaByAppid(appid string) (r *types.Media, err error) {
	req := &services.GetMediaByAppidRequest{Header: BuildCommonHeader(),
		Body: appid}
	resp := &services.GetMediaByAppidResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetMediaByAppid", req, resp); err != nil {
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

func GetMediaById(id int32) (r *types.Media, err error) {
	req := &services.GetMediaByIdsRequest{Header: BuildCommonHeader(),
		Body: []int32{id}}
	resp := &services.GetMediaByIdsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetMediaByIds", req, resp); err != nil {
		err = fmt.Errorf("call error: %v", err)
		return
	}
	if resp.Header.Code != enums.ResponseCode_OK {
		err = fmt.Errorf("call error: server reply code %d", resp.Header.Code)
		return
	}
	medias := resp.Body
	for _, v := range medias {
		r = v
	}

	return
}

func GetMediaByIds(ids []int32) (r map[int32]*types.Media, err error) {
	req := &services.GetMediaByIdsRequest{Header: BuildCommonHeader(),
		Body: ids}
	resp := &services.GetMediaByIdsResponse{}
	conn := Default.Get()
	defer conn.Close()
	if err = conn.Use().(client.XClient).Call(context.Background(), "GetMediaByIds", req, resp); err != nil {
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
