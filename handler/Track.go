package handler

import (
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ConvertToFixedLengthByteArray(slice []byte) ([32]byte, error) {
	if len(slice) != 32 {
		return [32]byte{}, fmt.Errorf("slice length is not 32")
	}
	var array [32]byte
	copy(array[:], slice)
	return array, nil
}

func Track(c *gin.Context) {
	todo := eth.TraceHash{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	// 将溯源码从字符串转换为字节数组
	traceHash, err := hex.DecodeString(todo.TraceHash)
	if err != nil {
		middlewares.RespError(c, err)
		return
	}
	array, err := ConvertToFixedLengthByteArray(traceHash)
	if err != nil {
		return
	}
	info, err := eth.Track(array)
	if err != nil {
		middlewares.RespError(c, err)
		return
	}
	// 过滤掉空地址
	info.FoundOnes = middlewares.FilterEmptyAddresses(info.FoundOnes)
	middlewares.RespOK(c, info)
}
