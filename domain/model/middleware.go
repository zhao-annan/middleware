package model

type Middleware struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment"`

	//中间件的名称
	MiddleName string `json:"middle_name"`

	//中间件创建的命名空间

	MiddleNamespace string `json:"middle_namespace"`

	//中间件的类型

	MiddleTypeID int64 `json:"middle_type_id"`

	//中间件的版本

	MiddleVersionID int64 `json:"middle_version_id"`

	//中间件的端口

	MiddlePort []MiddlePort `gorm:"ForeignKey:MiddleID" json:"middle_port"`

	//默认生成的账号密码
	MiddleConfig MiddleConfig `gorm:"ForeignKey:MiddleID" json:"middle_config"`

	//环境变量
	MiddleEnv []MiddleEnv `gorm:"ForeignKey:MiddleID" json:"middle_env"`

	//中间件的cpu
	MiddleCpu float32 `json:"middle_cpu"`

	//中间件的内存
	MiddleMemory float32 `json:"middle_memory"`

	//中间件的存储

	MiddleStorage []MiddleStorage `json:"middle_storage"`
}
