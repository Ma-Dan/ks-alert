package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/models"
	"time"
)

type AlertRuleHandler struct{}

// alert rule
func (server AlertRuleHandler) CreateAlertRule(ctx context.Context, ruleGroup *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {

	if len(ruleGroup.AlertRules) == 0 || ruleGroup.AlertRuleGroupName == "" {
		return &pb.AlertRuleGroupResponse{
			Error: &pb.Error{
				Code: pb.Error_INVALID_PARAM,
				Text: "invalid param",
			},
		}, nil
	}

	if ruleGroup.ResourceTypeId != "" {
		r, _ := models.GetResourceType(&models.ResourceType{ResourceTypeID: ruleGroup.ResourceTypeId})
		if r == nil || r.ResourceTypeID == "" {
			return &pb.AlertRuleGroupResponse{
				Error: &pb.Error{
					Code: pb.Error_INVALID_PARAM,
					Text: "resource type is not exist",
				},
			}, nil
		}

		alertRuleGroup, err := models.CreateAlertRuleGroup(ConvertPB2AlertRuleGroup(ruleGroup))
		var pErr *pb.Error
		if err != nil {
			pErr = &pb.Error{
				Code: pb.Error_ACCESS_DENIED,
				Text: err.Error(),
			}
		} else {
			pErr = &pb.Error{
				Code: pb.Error_SUCCESS,
				Text: "success",
			}
		}
		return &pb.AlertRuleGroupResponse{
			Error:          pErr,
			AlertRuleGroup: ConvertAlertRuleGroup2PB(alertRuleGroup),
		}, nil
	}

	return nil, nil
}

func (server AlertRuleHandler) UpdateAlertRule(ctx context.Context, alertRule *pb.AlertRuleGroup) (*pb.AlertRuleGroupResponse, error) {
	// check alert_rule_group_id is exist
	if alertRule.AlertRuleGroupId == "" || alertRule.AlertRuleGroupName == "" {
		return &pb.AlertRuleGroupResponse{
			Error: &pb.Error{
				Code: pb.Error_INVALID_PARAM,
				Text: "invalid param",
			},
		}, nil
	}

	err := models.UpdateAlertRuleGroup(ConvertPB2AlertRuleGroup(alertRule))

	var pErr *pb.Error
	if err != nil {
		pErr = &pb.Error{
			Code: pb.Error_ACCESS_DENIED,
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Code: pb.Error_SUCCESS,
			Text: "success",
		}
	}
	return &pb.AlertRuleGroupResponse{
		Error: pErr,
	}, nil
}

func (server AlertRuleHandler) GetAlertRule(ctx context.Context, alertRuleSpec *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {

	groupID := alertRuleSpec.AlertRuleGroupId

	// means to get supported alert rule for the resource type
	typeID := alertRuleSpec.ResourceTypeId
	//systemRule := alertRuleSpec.SystemRule

	if groupID == "" || typeID == "" {
		return &pb.AlertRuleGroupResponse{
			Error: &pb.Error{
				Code: pb.Error_INVALID_PARAM,
				Text: "invalid param",
			},
		}, nil
	}

	ruleGroup, err := models.GetAlertRuleGroup(alertRuleSpec)

	var pErr *pb.Error
	if err != nil {
		pErr = &pb.Error{
			Code: pb.Error_ACCESS_DENIED,
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Code: pb.Error_SUCCESS,
			Text: "success",
		}
	}
	return &pb.AlertRuleGroupResponse{
		Error:          pErr,
		AlertRuleGroup: ConvertAlertRuleGroup2PB(ruleGroup),
	}, nil
}

func (server AlertRuleHandler) DeleteAlertRule(ctx context.Context, alertRuleSpec *pb.AlertRuleGroupSpec) (*pb.AlertRuleGroupResponse, error) {

	groupID := alertRuleSpec.AlertRuleGroupId

	// means to get supported alert rule for the resource type
	typeID := alertRuleSpec.ResourceTypeId
	//systemRule := alertRuleSpec.SystemRule

	if groupID == "" || typeID == "" {
		return &pb.AlertRuleGroupResponse{
			Error: &pb.Error{
				Code: pb.Error_INVALID_PARAM,
				Text: "invalid param",
			},
		}, nil
	}

	err := models.DeleteAlertRuleGroup(alertRuleSpec)

	var pErr *pb.Error
	if err != nil {
		pErr = &pb.Error{
			Code: pb.Error_ACCESS_DENIED,
			Text: err.Error(),
		}
	} else {
		pErr = &pb.Error{
			Code: pb.Error_SUCCESS,
			Text: "success",
		}
	}
	return &pb.AlertRuleGroupResponse{
		Error: pErr,
	}, nil
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
				SystemRule:       ptr.SystemRule,
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
				SystemRule:             ptr.SystemRule,
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
