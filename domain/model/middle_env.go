package model

type MiddleEnv struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`

	//关联的环境变量ID

	MiddleID int64 `json:"middle_id"`

	//环境变量的key
	EnvKey string `json:"env_key"`

	//环境变量Value

	EnvValue string `json:"env_value"`
}
