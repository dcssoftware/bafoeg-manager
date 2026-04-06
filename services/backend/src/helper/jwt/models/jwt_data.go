package jwtmodels

type JwtDataModel struct {
	UserUUID  string   `json:"userUuid"`
	SessionID string   `json:"sessionUuid"`
	Scopes    []string `json:"scopes"`
}

func NewJwtDataModel(uuid, sessionID string, scopes []string) *JwtDataModel {
	return &JwtDataModel{
		UserUUID:  uuid,
		SessionID: sessionID,
		Scopes:    scopes,
	}
}

type RefreshDataModel struct {
	SessionID string `json:"sessionUuid"`
}

func NewRefreshDataModel(sessionID string) *RefreshDataModel {
	return &RefreshDataModel{
		SessionID: sessionID,
	}
}
