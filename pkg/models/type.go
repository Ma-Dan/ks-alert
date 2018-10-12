package models


type UserRequest struct {
	UserId   string `json:"userid"`
	Email    string `json:"email,omitempty"`
	ResourceType string `json:"resource_type"`
	ResourceName string `json:"resource_name"`
	MetricName   string `json:"metric_name"`
}
