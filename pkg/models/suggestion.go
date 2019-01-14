package models

import (
	"fmt"
	"github.com/carmanzhang/ks-alert/pkg/utils/dbutil"
	"strings"
	"time"
)

type Suggestion struct {
	AlertConfigID string `gorm:"primary_key;"`
	ResourceID    string `gorm:"primary_key;"`
	AlertRuleID   string `gorm:"primary_key;"`
	// json format message
	Message   string    `gorm:"type:text;"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}

func UpdateSuggestion(suggestion *Suggestion) (*Suggestion, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// may be there is no suggestions
	sugg, err := GetSuggestion(suggestion)

	if err != nil {
		return nil, err
	}

	var sql string
	var mes = strings.Trim(suggestion.Message, "")

	if sugg != nil {
		if mes != "" {
			// update mode, since suggestion exist and comming message not null
			sql = fmt.Sprintf("UPDATE suggestions SET message='%s', updated_at='%v' "+
				"WHERE alert_config_id='%s' AND alert_rule_id='%s' AND resource_id='%s'",
				suggestion.Message, time.Now(),
				suggestion.AlertConfigID, suggestion.AlertRuleID, suggestion.ResourceID)
		} else {
			// delete mode, since suggestion exist and comming message is null
			sql = fmt.Sprintf("DELETE from suggestions WHERE alert_config_id='%s' AND alert_rule_id='%s' AND resource_id='%s'",
				suggestion.AlertConfigID, suggestion.AlertRuleID, suggestion.ResourceID)
		}

	} else {
		// create mode, since suggestion does not exist
		sql = fmt.Sprintf("INSERT INTO suggestions (alert_config_id, alert_rule_id, resource_id, message, created_at, updated_at) "+
			"VALUES ('%s', '%s', '%s', '%s', '%v', '%v')",
			suggestion.AlertConfigID, suggestion.AlertRuleID, suggestion.ResourceID,
			suggestion.Message, time.Now(), time.Now())
	}

	fmt.Println(sql)

	err = db.Exec(sql).Error

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return suggestion, nil
}

func GetSuggestion(sugSpec *Suggestion) (*Suggestion, error) {

	db, err := dbutil.DBClient()

	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	// get suggestion by alert_config_id and alert_rule_id and resource_id
	var suggestion Suggestion

	//err = db.Debug().Raw("SELECT * from suggestions WHERE alert_config_id=? AND alert_rule_id=? AND resource_id=?",
	//	sugSpec.AlertConfigID, sugSpec.AlertRuleID, sugSpec.ResourceID).Scan(&suggestion).Error

	b := db.Where("alert_config_id=? AND alert_rule_id=? AND resource_id=?",
		sugSpec.AlertConfigID, sugSpec.AlertRuleID, sugSpec.ResourceID).First(&suggestion).RecordNotFound()

	if b {
		return nil, nil
	}

	err = db.Error
	if err != nil {
		return nil, Error{Text: err.Error(), Code: DBError}
	}

	return &suggestion, nil
}
