package blogApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/logging"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"time"
)

type MessageRestApi struct {
	base.Handler
}

func NewMessageRestApi() *MessageRestApi {
	return &MessageRestApi{}
}

var (
	MessageService service.IMessageService = impl.NewMessageServiceImpl()
)

func (c *MessageRestApi) GetMessages(ctx *gin.Context) {
	messages, err := MessageService.GetMessages()
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetMessagesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Messages:   messages,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *MessageRestApi) AddMessages(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		logging.Logger.Error(err)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	request.Message.IpAddress = ctx.ClientIP()
	request.Message.CreateTime = time.Now().Unix()
	err = MessageService.AddMessage(request.Message)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.AddMessageFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *MessageRestApi) GetAdminMessages(ctx *gin.Context) {
	type Condition struct {
		Current  int64  `json:"current" form:"current"`
		Size     int32  `json:"size" form:"size"`
		Keywords string `json:"keywords" form:"keywords"`
	}
	var condition Condition
	err := ctx.ShouldBind(&condition)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(condition)

	adminMessages, err := MessageService.GetAdminMessages(&pb.CsCondition{
		Current:  condition.Current,
		Size:     condition.Size,
		Keywords: condition.Keywords,
	})
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.GetMessagesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:         pb.Response_ResponseBeginIndex,
		Code:          pb.ResultCode_SuccessOK,
		ServerTime:    time.Now().Unix(),
		AdminMessages: adminMessages,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *MessageRestApi) DeleteMessage(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request.MessageIds)
	err = MessageService.DeleteMessage(request.MessageIds)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.DeleteMessageFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "删除成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
