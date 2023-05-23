package models

type EnhancedVideoResponse struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	EnhancedVideoUrl string `json:"enhancedVideoUrl" bson:"enhancedVideoUrl"`
	Status           string `json:"status" bson:"status"`
	StatusMessage    string `json:"statusMessage" bson:"statusMessage"`
}

type EnhancedVideoNotifyRequest struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces"`
	EnhancedVideoUrl   string   `json:"enhancedVideoUrl" bson:"enhancedVideoUrl"`
}
