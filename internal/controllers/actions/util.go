package actions

import (
	"AirAccountGateway/internal/models/webapi/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

func walletApiInvoker[T any](ctx *gin.Context, walletApi func() (*http.Response, error), success func(v *T)) {
	//v := new(T)
	//success(v)
	// @todo: 调用钱包API, change to your own wallet api
	if resp, err := walletApi(); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusFailedDependency, &msg)
		return
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			if body, err := io.ReadAll(resp.Body); err == nil {
				v := new(T)
				if err := json.Unmarshal(body, v); err != nil {
					msg := err.Error()
					response.Fail(ctx, http.StatusUnprocessableEntity, &msg)
				} else {
					success(v)
				}
			} else {
				msg := err.Error()
				response.Fail(ctx, http.StatusUnprocessableEntity, &msg)
				return
			}
		} else {
			msg := fmt.Sprintf("downstream code: %d", resp.StatusCode)
			response.Fail(ctx, http.StatusUnprocessableEntity, &msg)
			fmt.Println(msg)
			return
		}
	}
}

const (
	SrcSms = 0
)

// idToHash id转换为hash
func idToHash(rawId string, src int) string {
	data := []byte(rawId)
	// @todo: 根据src选择不同的hash算法
	hash := crypto.Keccak256Hash(data)
	return hash.Hex()
}
