package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/pb"
	"kubesphere.io/ks-alert/pkg/stderr"
	"time"
)

type ReceiverHandler struct{}

// alert rule
func (server ReceiverHandler) CreateReceiver(ctx context.Context, pbRecvGroup *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {

	recvGroup := ConvertPB2ReceiverGroup(pbRecvGroup)

	v, err := DoTransactionAction(recvGroup, MethodCreate)
	respon := getReceiverGroupResponse(v, err)
	return respon, nil
}

func getReceiverGroupResponse(v interface{}, err error) *pb.ReceiverGroupResponse {
	var resGroup *models.ReceiverGroup
	if v != nil {
		resGroup = v.(*models.ReceiverGroup)
	}

	rg := ConvertReceiverGroup2PB(resGroup)

	var respon = pb.ReceiverGroupResponse{ReceiverGroup: rg}
	respon.Error = stderr.ErrorWrapper(err)
	return &respon
}

func (server ReceiverHandler) DeleteReceiver(ctx context.Context, receiverSpec *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {
	// TODO only delete one receiver in a receiver group
	//recvID := receiverSpec.ReceiverId
	recvGroup := &models.ReceiverGroup{
		ReceiverGroupID: receiverSpec.ReceiverGroupId,
		//Receivers:       &[]models.Receiver{{ReceiverID: recvID}},
	}

	v, err := DoTransactionAction(recvGroup, MethodDelete)

	respon := getReceiverGroupResponse(v, err)
	return respon, nil
}

func (server ReceiverHandler) UpdateReceiver(ctx context.Context, pbRecvGroup *pb.ReceiverGroup) (*pb.ReceiverGroupResponse, error) {

	recvGroup := ConvertPB2ReceiverGroup(pbRecvGroup)

	v, err := DoTransactionAction(recvGroup, MethodUpdate)

	respon := getReceiverGroupResponse(v, err)
	return respon, nil
}

func (server ReceiverHandler) GetReceiver(ctx context.Context, receiverSpec *pb.ReceiverGroupSpec) (*pb.ReceiverGroupResponse, error) {

	recvGroupID := receiverSpec.ReceiverGroupId
	//recvID := receiverSpec.ReceiverId
	recvGroup := &models.ReceiverGroup{
		ReceiverGroupID: recvGroupID,
		//Receivers:       &[]models.Receiver{{ReceiverID: recvID}},
	}

	v, err := DoTransactionAction(recvGroup, MethodGet)
	respon := getReceiverGroupResponse(v, err)
	return respon, nil

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
