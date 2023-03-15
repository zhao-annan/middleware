package service

import (
	"github.com/zhao-annan/middleware/domain/model"
	"github.com/zhao-annan/middleware/domain/repository"
)

//定义接口类型

type IMiddleTypeDataService interface {
	AddMiddleType(middleType *model.MiddleType) (int64, error)
	DeleteMiddleType(int64) error
	UpdateMiddleType(middleType *model.MiddleType) error
	FindMiddleTypeByID(int64) (*model.MiddleType, error)
	FindAllMiddleType() ([]model.MiddleType, error)

	//根据ID返回地址
	FindImageVersionByID(int64) (string, error)

	FindVersionByID(int64) (*model.MiddleVersion, error)
	FindAllVersionByTypeID(int64) ([]model.MiddleVersion, error)
}

//注意：返回值的类型

func NewMiddleTypeDataService(repository repository.IMiddleTypeRepository) IMiddleTypeDataService {

	return &MiddleTypeDataService{MiddleTypeRepository: repository}

}

type MiddleTypeDataService struct {
	MiddleTypeRepository repository.IMiddleTypeRepository
}

//func (u *MiddleTypeDataService) UpdateMiddleType(middleType model.MiddleType) error {
//	//TODO implement me
//	panic("implement me")
//}

func (u *MiddleTypeDataService) FindImageVersionByID(i int64) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *MiddleTypeDataService) AddMiddleType(middleType *model.MiddleType) (int64, error) {

	return u.MiddleTypeRepository.CreateMiddleType(middleType)

}

func (u *MiddleTypeDataService) DeleteMiddleType(middleTypeID int64) error {

	return u.MiddleTypeRepository.DeleteMiddleTypeByID(middleTypeID)

}

func (u *MiddleTypeDataService) UpdateMiddleType(middleType *model.MiddleType) error {

	return u.MiddleTypeRepository.UpdateMiddleType(middleType)

}

//查找
func (u *MiddleTypeDataService) FindMiddleTypeByID(middleTypeID int64) (*model.MiddleType, error) {

	return u.MiddleTypeRepository.FindTypeByID(middleTypeID)

}

//查找所有

func (u *MiddleTypeDataService) FindAllMiddleType() ([]model.MiddleType, error) {

	return u.MiddleTypeRepository.FindAll()

}

//根据VersionId查找镜像地址

func (u *MiddleTypeDataService) FindImageVersionById(versionId int64) (string, error) {

	version, err := u.MiddleTypeRepository.FindVersionByID(versionId)

	if err != nil {
		return "", err

	}
	//返回需要的镜像地址
	return version.MiddleDockerImage + ":" + version.MiddleVS, nil

}

//根据versionID  查找单个镜像

func (u *MiddleTypeDataService) FindVersionByID(middleVersionID int64) (
	*model.MiddleVersion, error) {

	return u.MiddleTypeRepository.FindVersionByID(middleVersionID)
}

//根据中间件类型查找对应的所有版本

func (u *MiddleTypeDataService) FindAllVersionByTypeID(middleTypeID int64) ([]model.MiddleVersion, error) {

	return u.MiddleTypeRepository.FindAllVersionByTypeID(middleTypeID)

}
