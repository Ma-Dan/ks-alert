package handler

import (
	"context"
	"kubesphere.io/ks-alert/pkg/models"
	"kubesphere.io/ks-alert/pkg/pb"
	"kubesphere.io/ks-alert/pkg/stderr"
	"time"
)

type AlertRuleHandler struct{}

// alert rule
func (h AlertRuleHandler) CreateAlertRule(ctx context.Context, ruleGroup *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {
	v, err := DoTransactionAction(ConvertPB2AlertRuleGroup(ruleGroup), MethodCreate)
	respon := getAlertRuleGroupResponse(v, err)
	return respon, nil
}

func getAlertRuleGroupResponse(v interface{}, err error) *pb.AlertRuleGroupResponse {
	var ruleGroup *models.AlertRuleGroup
	if v != nil {
		ruleGroup = v.(*models.AlertRuleGroup)
	}

	arg := ConvertAlertRuleGroup2PB(ruleGroup)

	var respon = pb.AlertRuleGroupResponse{AlertRuleGroup: arg}
	respon.Error = stderr.ErrorWrapper(err)

	return &respon
}

func (h AlertRuleHandler) UpdateAlertRule(ctx context.Context, ruleGroup *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {
	v, err := DoTransactionAction(ConvertPB2AlertRuleGroup(ruleGroup), MethodUpdate)

	respon := getAlertRuleGroupResponse(v, err)
	return respon, nil
}

func (h AlertRuleHandler) GetAlertRule(ctx context.Context, alertRuleSpec *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {

	ruleGroup := models.AlertRuleGroup{
		AlertRuleGroupID: alertRuleSpec.AlertRuleGroupId,
		ResourceTypeID:   alertRuleSpec.ResourceTypeId,
		SystemRule:       alertRuleSpec.SystemRule,
	}

	v, err := DoTransactionAction(&ruleGroup, MethodGet)

	respon := getAlertRuleGroupResponse(v, err)
	return respon, nil
}

func (h AlertRuleHandler) DeleteAlertRule(ctx context.Context, alertRuleSpec *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {
	ruleGroup := models.AlertRuleGroup{
		AlertRuleGroupID: alertRuleSpec.AlertRuleGroupId,
		ResourceTypeID:   alertRuleSpec.ResourceTypeId,
		SystemRule:       alertRuleSpec.SystemRule,
	}

	v, err := DoTransactionAction(&ruleGroup, MethodDelete)
	respon := getAlertRuleGroupResponse(v, err)
	return respon, nil
}

func ConvertPB2AlertRuleGroup(pbRuleGroup *pb.AlertRuleGroup) *models.AlertRuleGroup {
	ruleGroup := &models.AlertRuleGroup{
		AlertRuleGroupID:   pbRuleGroup.AlertRuleGroupId,
		AlertRuleGroupName: pbRuleGroup.AlertRuleGroupName,
		AlertRules:         ConvertPB2AlertRules(pbRuleGroup.AlertRules),
		Description:        pbRuleGroup.Desc,
		SystemRule:         pbRuleGroup.SystemRule,
		ResourceTypeID:     pbRuleGroup.ResourceTypeId,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
	}

	return ruleGroup
}

func ConvertAlertRuleGroup2PB(ruleGroup *models.AlertRuleGroup) *pb.AlertRuleGroup {
	if ruleGroup == nil {
		return nil
	}
	pbRuleGroup := &pb.AlertRuleGroup{
		AlertRuleGroupId:   ruleGroup.AlertRuleGroupID,
		AlertRuleGroupName: ruleGroup.AlertRuleGroupName,
		AlertRules:         ConvertAlertRules2PB(ruleGroup.AlertRules),
		SystemRule:         ruleGroup.SystemRule,
		Desc:               ruleGroup.Description,
		ResourceTypeId:     ruleGroup.ResourceTypeID,
	}

	return pbRuleGroup
}

func ConvertAlertRules2PB(alertRules []*models.AlertRule) []*pb.AlertRule {
	if alertRules != nil {
		l := len(alertRules)
		var pbAlertRules = make([]*pb.AlertRule, l)
		for i := 0; i < l; i++ {
			ptr := alertRules[i]
			pbAlertRules[i] = &pb.AlertRule{
				AlertRuleId:      ptr.AlertRuleID,
				AlertRuleName:    ptr.AlertRuleName,
				AlertRuleGroupId: ptr.AlertRuleGroupID,
				MetricName:       ptr.MetricName,
				ConditionType:    ptr.ConditionType,
				Threshold:        ptr.Threshold,
				Period:           ptr.Period,
				Unit:             ptr.Unit,
				ConsecutiveCount: ptr.ConsecutiveCount,
				Enable:           ptr.Enable,
				RepeatSend: &pb.RepeatSend{
					InitRepeatSendInterval: ptr.InitRepeatSendInterval,
					RepeatSendType:         pb.RepeatSendType(ptr.RepeatSendType),
					MaxRepeatSendCount:     ptr.MaxRepeatSendCount,
				},
				PreferSeverity:    ptr.PerferSeverity,
				InhibitRuleEnable: ptr.InhibitRule,
			}
		}

		return pbAlertRules
	}

	return nil
}

func ConvertPB2AlertRules(pbAlertRules []*pb.AlertRule) []*models.AlertRule {
	if pbAlertRules != nil {
		l := len(pbAlertRules)
		var alertRules = make([]*models.AlertRule, l)
		for i := 0; i < l; i++ {
			ptr := pbAlertRules[i]
			alertRules[i] = &models.AlertRule{
				AlertRuleID:            ptr.AlertRuleId,
				AlertRuleName:          ptr.AlertRuleName,
				AlertRuleGroupID:       ptr.AlertRuleId,
				MetricName:             ptr.MetricName,
				ConditionType:          ptr.ConditionType,
				PerferSeverity:         ptr.PreferSeverity,
				Threshold:              ptr.Threshold,
				Period:                 ptr.Period,
				Unit:                   ptr.Unit,
				ConsecutiveCount:       ptr.ConsecutiveCount,
				InhibitRule:            ptr.InhibitRuleEnable,
				Enable:                 ptr.Enable,
				RepeatSendType:         int32(ptr.RepeatSend.RepeatSendType),
				InitRepeatSendInterval: ptr.RepeatSend.InitRepeatSendInterval,
				MaxRepeatSendCount:     ptr.RepeatSend.MaxRepeatSendCount,
				UpdatedAt:              time.Now(),
				CreatedAt:              time.Now(),
			}
		}
		return alertRules
	}

	return nil
}
