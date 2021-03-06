package omg

import (
	"encoding/json"
)

const (
	NotFound     = 404
	InternalErr  = 503
	AlreadyExist = 504
	InvalidParam = 505
	CreateErr    = 506
)

func JsonFormatErr() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "json format error",
		Code:    InternalErr,
	}
	out, _ := json.Marshal(s)
	return out
}
func InternalError() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Internal error",
		Code:    InternalErr,
	}
	out, _ := json.Marshal(s)
	return out
}

func UnknownReq() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Unknown request params",
		Code:    NotFound,
	}
	out, _ := json.Marshal(s)
	return out
}

func BucketNotFound() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Bucket not found",
		Code:    NotFound,
	}
	out, _ := json.Marshal(s)
	return out
}

// File not exist
func FileNotFound() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "File not found",
		Code:    NotFound,
	}

	out, _ := json.Marshal(s)
	return out
}

// File already exist
func FileAlreadyExist() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "File already exist",
		Code:    AlreadyExist,
	}

	out, _ := json.Marshal(s)
	return out
}

func InvalidParams() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: " Invalid params",
		Code:    InvalidParam,
	}
	out, _ := json.Marshal(s)
	return out
}

func CreateRecordFailed() []byte {
	s := ResponseCommon{
		Version: Version,
		SeqNum:  1,
		From:    "omigad",
		To:      "client",
		Type:    "omigad",
		Number:  "XXXX-XXXX-XXXX-XXXX",
		Message: "Failed to create record for the file manager",
		Code:    CreateErr,
	}
	out, _ := json.Marshal(s)
	return out
}

func notFound() []byte {
	s := ResNotFound{
		Status: 404,
		Msg:    "not support request.",
	}

	out, _ := json.Marshal(s)
	return out
}
