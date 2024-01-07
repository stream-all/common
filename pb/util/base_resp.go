package util

import (
	"github.com/stream-all/common/pb"
	"github.com/stream-all/common/pb/error"
)

func OkBaseResp() *pb.BaseResp {
	return &pb.BaseResp{
		Code: error.CodeOk,
		Msg:  "ok",
	}
}

func BaseResp(code int32, msg string) *pb.BaseResp {
	return &pb.BaseResp{
		Code: code,
		Msg:  msg,
	}
}

func ParseBaseResp(baseResp *pb.BaseResp) (int32, string) {
	if baseResp == nil {
		return error.CodeUnknown, "baseResp is nil"
	}
	return baseResp.Code, baseResp.Msg
}
