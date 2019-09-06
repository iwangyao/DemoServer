package Handler

import (
	"../Global"
	"../NetFrame"
	"../proto/dto"
	"fmt"
)

type HandlerData struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     *Global.ClientState
}
type HandlerFunc interface {
	ReceiveMessage()
}

func NewHandlerData(decode *NetFrame.Decode, msg []byte, _client *Global.ClientState) *HandlerData {
	handlerData := &HandlerData{
		command:    decode.Command,
		bytesStart: decode.ReadPos,
		bytesEnd:   decode.Len + 4,
		messages:   msg,
		client:     _client,
	}
	return handlerData
}

//消息处理中心，处理和分发所有玩家的登录 匹配 战斗信息
func Handler(msg []byte, client *Global.ClientState) {
	var decode NetFrame.Decode
	decode.Read(msg)
	//fmt.Println("不加密", msg[decode.ReadPos:decode.Len+4])

	tmp := string(msg[decode.ReadPos : decode.Len+4])
	//tmp := string(ttt[:])
	fmt.Println(msg[decode.ReadPos : decode.Len+4])
	msg2 := NetFrame.Decrypt(tmp, string(NetFrame.MyKey[:]))
	handlerData := NewHandlerData(&decode, []byte(msg2), client)
	fmt.Println(handlerData.messages)
	var p HandlerFunc

	//HANDLER CENTER
	switch decode.Thetype {
	case int32(DTO.MsgTypes_TYPE_LOGIN):
		{
			p = &Login{handlerData}
			p.ReceiveMessage()
		}
		break

	case int32(DTO.MsgTypes_TYPE_USER):
		break
	case int32(DTO.MsgTypes_TYPE_MATCH):
		{
			p = &Match{handlerData}
			p.ReceiveMessage()
		}
		break
	case int32(DTO.MsgTypes_TYPE_FIGHT):
		{
			p = &Fight{handlerData}
			p.ReceiveMessage()
		}
		break
	default:
		break
	}
}
