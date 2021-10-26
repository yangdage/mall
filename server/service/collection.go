package service

import (
	"fmt"
	"mall.com/global"
	"mall.com/models"
	"strconv"
)

type Collection struct {

}

func (c *Collection) AppAdd(param models.AppCollectionParam) int64 {
	key := fmt.Sprintf("user:%s:collect", param.UserId)
	if count := global.Rdb.SAdd(ctx, key, param.ProductId).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Collection) AppDelete(uid string) int64 {
	key := fmt.Sprintf("user:%s:collect", uid)
	pids := global.Rdb.SMembers(ctx, key).Val()
	if count := global.Rdb.SRem(ctx, key, pids).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Collection) AppGetList(uid string) []models.AppCollectionItem {
	var items []models.AppCollectionItem
	key := fmt.Sprintf("user:%s:collect", uid)
	pids := global.Rdb.SMembers(ctx, key).Val()
	productIds := make([]uint, 0)
	for _, pid := range pids {
		id, _ := strconv.Atoi(pid)
		productIds = append(productIds, uint(id))
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&items, productIds)
		return items
	}
	return items
}
