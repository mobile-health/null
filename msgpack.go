package null

import "github.com/vmihailenco/msgpack/v5/msgpcode"

var MsgpackNil = []byte{msgpcode.Nil}

func IsMsgpackNil(data []byte) bool {
	if len(data) == 1 && data[0] == msgpcode.Nil {
		return true
	}
	return false
}
