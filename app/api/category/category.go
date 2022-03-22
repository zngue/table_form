package category

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/zngue/form_common/app/model"
	"github.com/zngue/form_common/app/service"
	"github.com/zngue/go_helper/pkg/api"
	"golang.org/x/sync/errgroup"
)

func ApiBase() service.ICategory {
	return service.NewCategory()
}

// Create  添加数据
func Create(ctx *gin.Context) {
	var req service.CategoryRequest
	var data model.Category
	if err := ctx.ShouldBind(&data); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	req.Data = &data
	err := ApiBase().Add(&req)
	api.DataWithErr(ctx, err, data)
	return
}

// Edit  修改数据
func Edit(ctx *gin.Context) {
	var req service.CategoryRequest
	formMap := ctx.PostFormMap("data")
	updateData := make(map[string]interface{})
	for key, val := range formMap {
		if key != "id" {
			updateData[key] = val
		}
	}
	//判断where 条件id是否存在  自行更改更新条件
	if id, ok := formMap["id"]; ok {
		newID := cast.ToInt(id)
		if newID <= 0 {
			api.Error(ctx, api.Msg("id 不能为空"))
			return
		}
		req.ID = newID
	} else {
		api.DataWithErr(ctx, errors.New("参数错误"), nil)
	}
	req.Data = updateData
	err := ApiBase().Save(&req)
	api.DataWithErr(ctx, err, nil)
}

type ListData struct {
	Count int64             `json:"count"`
	List  *[]model.Category `json:"list"`
}

func List(ctx *gin.Context) {
	var req service.CategoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	var data ListData
	var wg errgroup.Group
	wg.Go(func() error {
		list, err := ApiBase().List(&req)
		data.List = list
		return err
	})
	wg.Go(func() error {
		count, err := ApiBase().Count(&req)
		data.Count = count
		return err
	})
	err := wg.Wait()
	api.DataWithErr(ctx, err, data)
}
func Detail(ctx *gin.Context) {
	var req service.CategoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	list, err := ApiBase().Detail(&req)
	api.DataWithErr(ctx, err, list)
}
func Delete(ctx *gin.Context) {
	var req service.CategoryRequest
	if err := ctx.ShouldBind(&req); err != nil {
		api.Error(ctx, api.Err(err))
		return
	}
	err := ApiBase().Delete(&req)
	api.DataWithErr(ctx, err, nil)
}
