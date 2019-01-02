package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"time"
)

type ReceiverHandler struct{}

// alert rule
func (server ReceiverHandler) CreateReceiver(ctx context.Context, pbRecvGroup *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {

	if pbRecvGroup.ReceiverGroupName == "" {
		return nil, errors.New("the receiver group name must be specified")
	}

	if pbRecvGroup.Receivers == nil || len(pbRecvGroup.Receivers) == 0 {
		return nil, errors.New("the receiver group must contain at least one receiver")
	}

	recvGroup := ConvertPB2ReceiverGroup(pbRecvGroup)

	v, err := DoTransactionAction(recvGroup, ReceiverGroup, MethodCreate)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var respon *models.ReceiverGroup

	if v != nil {
		respon = v.(*models.ReceiverGroup)
	}

	return &pb.ReceiverGroupResponse{
		ReceiverGroup: ConvertReceiverGroup2PB(respon),
	}, nil
}

func (server ReceiverHandler) DeleteReceiver(ctx context.Context, receiverSpec *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {
	recvGroupID := receiverSpec.ReceiverGroupId

	if recvGroupID == "" {
		return nil, errors.New("receiver group id must be specified")
	}

	// TODO only delete one receiver in a receiver group
	//recvID := receiverSpec.ReceiverId
	recvGroup := &models.ReceiverGroup{
		ReceiverGroupID: recvGroupID,
		//Receivers:       &[]models.Receiver{{ReceiverID: recvID}},
	}

	_, err := DoTransactionAction(recvGroup, ReceiverGroup, MethodDelete)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	return &pb.ReceiverGroupResponse{}, nil
}

func (server ReceiverHandler) UpdateReceiver(ctx context.Context, pbRecvGroup *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {

	if pbRecvGroup.ReceiverGroupId == "" {
		return nil, errors.New("the receiver group id must be specified")
	}

	recvGroup := ConvertPB2ReceiverGroup(pbRecvGroup)

	v, err := DoTransactionAction(recvGroup, ReceiverGroup, MethodUpdate)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var respon *models.ReceiverGroup

	if v != nil {
		respon = v.(*models.ReceiverGroup)
	}

	return &pb.ReceiverGroupResponse{
		ReceiverGroup: ConvertReceiverGroup2PB(respon),
	}, nil
}

func (server ReceiverHandler) GetReceiver(ctx context.Context, receiverSpec *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {

	recvGroupID := receiverSpec.ReceiverGroupId

	if recvGroupID == "" {
		return nil, errors.New("receiver group id must be specified")
	}

	//recvID := receiverSpec.ReceiverId
	recvGroup := &models.ReceiverGroup{
		ReceiverGroupID: recvGroupID,
		//Receivers:       &[]models.Receiver{{ReceiverID: recvID}},
	}

	v, err := DoTransactionAction(recvGroup, ReceiverGroup, MethodGet)

	if err != nil {
		glog.Errorln(err.Error())
		return nil, err
	}

	var respon *models.ReceiverGroup

	if v != nil {
		respon = v.(*models.ReceiverGroup)
	}

	return &pb.ReceiverGroupResponse{
		ReceiverGroup: ConvertReceiverGroup2PB(respon),
	}, nil

}

func ConvertPB2ReceiverGroup(pbRecvGroup *pb.ReceiverGroup) *models.ReceiverGroup {
	recvGroup := models.ReceiverGroup{
		ReceiverGroupID:   pbRecvGroup.ReceiverGroupId,
		ReceiverGroupName: pbRecvGroup.ReceiverGroupName,
		Webhook:           pbRecvGroup.Webhook,
		WebhookEnable:     pbRecvGroup.WebhookEnable,
		Description:       pbRecvGroup.Desc,
		UpdatedAt:         time.Now(),
		CreatedAt:         time.Now(),
		Receivers:         ConvertPB2Receiver(&pbRecvGroup.Receivers),
	}

	return &recvGroup
}

func ConvertReceiverGroup2PB(recvGroup *models.ReceiverGroup) *pb.ReceiverGroup {
	if recvGroup == nil {
		return nil
	}

	pbRecvGroup := pb.ReceiverGroup{
		ReceiverGroupId:   recvGroup.ReceiverGroupID,
		ReceiverGroupName: recvGroup.ReceiverGroupName,
		Webhook:           recvGroup.Webhook,
		WebhookEnable:     recvGroup.WebhookEnable,
		Desc:              recvGroup.Description,
		Receivers:         *ConvertReceiver2PB(recvGroup.Receivers),
	}

	return &pbRecvGroup
}

func ConvertPB2Receiver(pbRecvs *[]*pb.Receiver) *[]models.Receiver {
	l := len(*pbRecvs)
	var receivers = make([]models.Receiver, l)
	for i := 0; i < l; i++ {
		r := (*pbRecvs)[i]
		receivers[i] = models.Receiver{
			ReceiverID:   r.ReceiverId,
			ReceiverName: r.ReceiverName,
			Phone:        r.Phone,
			Email:        r.Email,
			Wechat:       r.Wechat,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
	}

	return &receivers
}

func ConvertReceiver2PB(recvs *[]models.Receiver) *[]*pb.Receiver {
	l := len(*recvs)
	var receivers = make([]*pb.Receiver, l)
	for i := 0; i < l; i++ {
		r := (*recvs)[i]
		receivers[i] = &pb.Receiver{
			ReceiverId:   r.ReceiverID,
			ReceiverName: r.ReceiverName,
			Phone:        r.Phone,
			Email:        r.Email,
			Wechat:       r.Wechat,
		}
	}

	return &receivers
}
