package models

type EnhancedVideoResponse struct {
	UserId               string `json:"userId" bson:"userId"`
	RequestId            string `json:"requestId" bson:"requestId"`
	EnhancedVideoUrl     string `json:"enhancedVideoUrl" bson:"enhancedVideoUrl"`
	EnhancedVideoQuality string `json:"enhancedVideoQuality" bson:"enhancedVideoQuality"`
	Status               string `json:"status" bson:"status"`
	StatusMessage        string `json:"statusMessage" bson:"statusMessage"`
}

type EnhancedVideoNotifyRequest struct {
	UserId               string `json:"userId" bson:"userId"`
	RequestId            string `json:"requestId" bson:"requestId"`
	EnhancedVideoUrl     string `json:"enhancedVideoUrl" bson:"enhancedVideoUrl"`
	EnhancedVideoQuality string `json:"enhancedVideoQuality" bson:"enhancedVideoQuality"`
	Status               string `json:"status" bson:"status"`
}

type NotificationInterfacesRequest struct {
	NotificationInterfaces []string `json:"notificationInterfaces" bson:"notificationInterfaces" binding:"required,are-notification-interfaces-valid"`
}
