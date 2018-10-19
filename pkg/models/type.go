package models


//type UserRequest struct {
//	UserId   string `json:"userid"`
//	Email    string `json:"email"`
//	ResourceType string `json:"resource_type"`
//	ResourceName string `json:"resource_name"`
//	MetricName   string `json:"metric_name"`
//}

type UserRequest struct {
	UserId   string
	Email    string
	ResourceType string
	ResourceName ResourceName
	MetricName   string
}


type ResourceName struct {
	Workspace string
	Namespace string
	Workload  Workload
	Pod       string
	Container string
}

type Workload struct {
	Kind string
	Name string
}