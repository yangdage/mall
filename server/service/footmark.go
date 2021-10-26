package service

import (
	"fmt"
	"mall.com/global"
	"mall.com/models"
	"strconv"
)

type Footmark struct {

}

func (c *Footmark) AppAdd(param models.AppFootmarkParam) int64 {
	key := fmt.Sprintf("user:%s:footmark", param.UserId)
	if count := global.Rdb.LPush(ctx, key, param.ProductId).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Footmark) AppDelete(uid string) int64 {
	key := fmt.Sprintf("user:%s:footmark", uid)
	pids := global.Rdb.LRange(ctx, key, 0, -1).Val()
	for _, id := range pids {
		global.Rdb.LRem(ctx, key, 1, id).Val()
	}
	if count := global.Rdb.LLen(ctx, key).Val(); count == 0 {
		return 1
	}
	return 0
}

func (c *Footmark) AppGetList(uid string) []models.AppFootmarkItem {
	var items []models.AppFootmarkItem
	key := fmt.Sprintf("user:%s:footmark", uid)
	pids := global.Rdb.LRange(ctx, key, 0, -1).Val()
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

