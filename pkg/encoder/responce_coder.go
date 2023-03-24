package encoder

import (
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/golang/protobuf/proto"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Reason string `json:"reason"`
	Message string `json:"message"`
}

func RespEncoder(w http.ResponseWriter, r *http.Request, i interface{}) error {
	codec := encoding.GetCodec("json")
	messageMap := make(map[string]interface{})
	messageStr,  _ := codec.Marshal(i.(proto.Message))
	fmt.Println(messageStr)
	_ = codec.Unmarshal(messageStr, &messageMap)

	if len(messageMap) == 1 {
		for _, v := range messageStr{
			i = v
		}
	}

	resp := Response{
		Code:    200,
		Reason:  "",
		Message: "",
	}

	if msg, ok := messageMap["message"]; ok {
		i = msg
	}

	message, err := codec.Marshal(i)
	_ = json.Unmarshal(message, &resp.Message)
	if err != nil {
		return err
	}

	data, err := codec.Marshal(resp)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

	return nil
}