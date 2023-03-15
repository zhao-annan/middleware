package handler

import (
	"context"
	log "github.com/asim/go-micro/v3/logger"
	"github.com/zhao-annan/common"
	"github.com/zhao-annan/middleware/domain/model"
	"github.com/zhao-annan/middleware/domain/service"
	middleware "github.com/zhao-annan/middleware/proto/middleware"
	"strconv"
)

type MiddlewareHandler struct {
	//注意这里的类型是 IMiddlewareDataService 接口类型
	MiddlewareDataService service.IMiddlewareDataService
	MiddleTypeDataService service.IMiddleTypeDataService
}

// Call is a single request handler called via client.Call or the generated client code
func (e *MiddlewareHandler) AddMiddleware(ctx context.Context, info *middleware.MiddlewareInfo, rsp *middleware.Response) error {
	log.Info("Received *middleware.AddMiddleware request")
	middleModel := &model.Middleware{}
	if err := common.SwapTo(info, middleModel); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
		//调用其他的服务处理数据
		//根据给定的id查找镜像地址
	} else {
		imageAddress, err := e.MiddleTypeDataService.FindImageVersionByID(info.MiddleVersionId)
		if err != nil {
			common.Error(err)
			rsp.Msg = err.Error()
			return err
		}
		//赋值
		info.MiddleDockerImageVersion = imageAddress
		//在k8s中创建资源
		if err := e.MiddlewareDataService.CreateToK8s(info); err != nil {
			common.Error(err)
			rsp.Msg = err.Error()
			return err
		} else {
			//从数据库中插入数据
			middleID, err := e.MiddlewareDataService.AddMiddleware(middleModel)
			if err != nil {
				common.Error(err)
				rsp.Msg = err.Error()
				return err
			}
			rsp.Msg = "中间件添加成功" + strconv.FormatInt(middleID, 10)
			common.Info(rsp.Msg)
		}

	}

	return nil
}

func (e *MiddlewareHandler) DeleteMiddleware(ctx context.Context, req *middleware.MiddlewareId, rsp *middleware.Response) error {
	log.Info("Received *middleware.DeleteMiddleware request")

	middleModel, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	//从k8s中  删除middleware
	if err := e.MiddlewareDataService.DeleteFromK8s(middleModel); err != nil {

		common.Error(err)

		rsp.Msg = err.Error()

		return err

	}

	return nil
}

func (e *MiddlewareHandler) UpdateMiddleware(ctx context.Context, req *middleware.MiddlewareInfo, rsp *middleware.Response) error {
	log.Info("Received *middleware.UpdateMiddleware request")

	//更新到k8s
	if err := e.MiddlewareDataService.UpdateToK8s(req); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	//更新到数据库
	//先查出来数据
	middleModel, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return nil
	}
	common.SwapTo(req, middleModel)
	//执行更新操作
	if err := e.MiddlewareDataService.UpdateMiddleware(middleModel); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return nil
	}
	return nil
}

func (e *MiddlewareHandler) FindMiddlewareByID(ctx context.Context, req *middleware.MiddlewareId, rsp *middleware.MiddlewareInfo) error {
	log.Info("Received *middleware.FindMiddlewareByID request")
	middleModel, err := e.MiddlewareDataService.FindMiddlewareByID(req.Id)
	if err != nil {
		common.Error(err)
		return err
	}

	if err := common.SwapTo(middleModel, rsp); err != nil {
		common.Error(err)
		return err
	}

	return nil
}

func (e *MiddlewareHandler) FindAllMiddleware(ctx context.Context, req *middleware.FindAll, rsp *middleware.AllMiddleware) error {
	log.Info("Received *middleware.FindAllMiddleware request")

	middlewares, err := e.MiddlewareDataService.FindAllMiddleware()

	if err != nil {
		common.Error(err)
		return err
	}
	//对middleware格式进行转化
	for _, v := range middlewares {

		middleModel := &middleware.MiddlewareInfo{}

		if err := common.SwapTo(v, middleModel); err != nil {

			common.Error(err)
			return err
		}
		rsp.MiddlewareInfo = append(rsp.MiddlewareInfo, middleModel)

	}

	return nil
}

//查同一种类型下的所有的中间件
func (e *MiddlewareHandler) FindAllMiddlewareByTypeID(ctx context.Context, req *middleware.FindAllByTypeId, rsp *middleware.AllMiddleware) error {

	log.Info("Received *middleware.FindAllMiddlewareBytypeID")

	allmiddleWare, err := e.MiddlewareDataService.FindAllMiddlewareByTypeID(req.TypeId)

	if err != nil {
		common.Error(err)

		return err

	}

	for _, v := range allmiddleWare {
		middleware := &middleware.MiddlewareInfo{}

		if err := common.SwapTo(v, middleware); err != nil {

			common.Error(err)
			return err

		}
		rsp.MiddlewareInfo = append(rsp.MiddlewareInfo, middleware)

	}
	return nil

}

//根据id查找中间件类型信息

func (e *MiddlewareHandler) FindMiddleTypeByID(ctx context.Context, req *middleware.MiddleTypeId, rsp *middleware.MiddleTypeInfo) error {

	middleTypeModel, err := e.MiddleTypeDataService.FindMiddleTypeByID(req.Id)

	if err != nil {
		common.Error(err)

		return err
	}

	if err := common.SwapTo(middleTypeModel, rsp); err != nil {
		common.Error(err)

		return err

	}

	return nil

}

//添加中间件类型
func (e *MiddlewareHandler) AddMiddleType(ctx context.Context, req *middleware.MiddleTypeInfo, rsp *middleware.Response) error {

	middleTypeInfo := &model.MiddleType{}

	if err := common.SwapTo(req, middleTypeInfo); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	//将转换的结构体添加到数据库

	id, err := e.MiddleTypeDataService.AddMiddleType(middleTypeInfo)

	if err != nil {

		common.Error(err)

		rsp.Msg = err.Error()
		return err

	}
	rsp.Msg = "添加成功了middletypeID" + strconv.FormatInt(id, 10)

	return nil

}

//删除中间件
func (u *MiddlewareHandler) DeleteMiddleTypeByID(ctx context.Context, req *middleware.MiddleTypeId, rsp *middleware.Response) error {

	err := u.MiddleTypeDataService.DeleteMiddleType(req.Id)

	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	return nil

}

//更新中间件类型
func (u *MiddlewareHandler) UpdateMiddleType(ctx context.Context, req *middleware.MiddleTypeInfo, rsp *middleware.Response) error {

	middleTypemodel, err := u.MiddleTypeDataService.FindMiddleTypeByID(req.Id)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	if err := common.SwapTo(req, middleTypemodel); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	//正式地来更新数据

	if err := u.MiddleTypeDataService.UpdateMiddleType(middleTypemodel); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}

	return nil

}

//查找所有的类型
func (u *MiddlewareHandler) FindAllMiddleType(ctx context.Context, req *middleware.FindAll, rsp *middleware.AllMiddleType) error {

	allMiddleType, err := u.MiddleTypeDataService.FindAllMiddleType()
	if err != nil {
		common.Error(err)

		return err
	}

	for _, v := range allMiddleType {
		middleTypeModel := &middleware.MiddleTypeInfo{}
		if err := common.SwapTo(v, middleTypeModel); err != nil {

			common.Error(err)
			return err
		}

		rsp.MiddleTypeInfo = append(rsp.MiddleTypeInfo, middleTypeModel)
	}
	return nil
}
