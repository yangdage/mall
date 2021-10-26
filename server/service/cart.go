package service

import (
	"context"
	"fmt"
	"mall.com/global"
	"mall.com/models"
	"strconv"
)

var ctx = context.Background()

type Cart struct {

}

func (c *Cart) AppAdd(pid uint, uid string) int64 {
	key := fmt.Sprintf("user:%s:cart", uid)
	if count := global.Rdb.SAdd(ctx, key, pid).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Cart) AppDelete(pid uint, uid string) int64 {
	key := fmt.Sprintf("user:%s:cart", uid)
	if count := global.Rdb.SRem(ctx, key, pid).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Cart) AppClear(uid string) int64 {
	key := fmt.Sprintf("user:%s:cart", uid)
	pids := global.Rdb.SMembers(ctx, key).Val()
	if count := global.Rdb.SRem(ctx, key, pids).Val(); count > 0 {
		return count
	}
	return 0
}

func (c *Cart) AppGetInfo(uid string) models.AppCartInfo {
	var cartInfo models.AppCartInfo
	key := fmt.Sprintf("user:%s:cart", uid)
	pids := global.Rdb.SMembers(ctx, key).Val()
	productIds := make([]uint, 0)
	for _, pid := range pids {
		id, _ := strconv.Atoi(pid)
		productIds = append(productIds, uint(id))
	}
	if len(productIds) > 0 {
		global.Db.Table("product").Find(&cartInfo.CartItem, productIds)
		for _, item := range cartInfo.CartItem {
			cartInfo.TotalPrice += item.Price
		}
		return cartInfo
	}
	return cartInfo
}


