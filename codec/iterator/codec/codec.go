package codec

import (
	"github.com/json-iterator/go"
)

type JsoniterCodec struct {
}

func (c *JsoniterCodec) Decode(data []byte, i interface{}) error {
	return jsoniter.Unmarshal(data, i)
}

func (c *JsoniterCodec) Encode(i interface{}) ([]byte, error) {
	return jsoniter.Marshal(i)
}
