package models

type EnhancedVideoResponse struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	EnhancedVideoUri string `json:"EnhancedVideoUri" bson:"EnhancedVideoUri"`
}

type EnhancedVideoNotifyRequest struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces"`
	EnhancedVideoUri   string   `json:"EnhancedVideoUri" bson:"EnhancedVideoUri"`
}

// TODO: change uri back to url and apply binding
