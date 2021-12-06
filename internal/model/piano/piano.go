package piano

import "helloworld/internal/model"

type Create struct {
	model.BaseModel
	Name      string `json:"name,omitempty"`
	Status    int64  `json:"status,omitempty"`
	Type      int64  `json:"type,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	StopTime  int64  `json:"stop_time,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
}
