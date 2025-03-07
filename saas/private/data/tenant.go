package data

import (
	"context"
	"errors"
	"fmt"
	kitgorm "github.com/go-saas/kit/pkg/gorm"
	v1 "github.com/go-saas/kit/saas/api/tenant/v1"
	"github.com/go-saas/kit/saas/private/biz"
	"github.com/google/uuid"
	"github.com/goxiaoy/go-eventbus"
	gg "gorm.io/gorm"
)

type TenantRepo struct {
	*kitgorm.Repo[biz.Tenant, string, v1.ListTenantRequest]
}

func NewTenantRepo(eventbus *eventbus.EventBus, data *Data) biz.TenantRepo {
	res := &TenantRepo{}
	res.Repo = kitgorm.NewRepo[biz.Tenant, string, v1.ListTenantRequest](data.DbProvider, eventbus, res)
	return res
}

func (g *TenantRepo) GetDb(ctx context.Context) *gg.DB {
	ret := GetDb(ctx, g.DbProvider)
	return ret
}

func (g *TenantRepo) BuildDetailScope(withDetail bool) func(db *gg.DB) *gg.DB {
	return func(db *gg.DB) *gg.DB {
		if withDetail {
			return db.Preload("Conn").Preload("Features")
		}
		return db
	}
}
func (g *TenantRepo) DefaultSorting() []string {
	return []string{
		"-created_at",
	}
}

func (g *TenantRepo) BuildFilterScope(q *v1.ListTenantRequest) func(db *gg.DB) *gg.DB {
	search := q.Search
	filter := q.Filter
	return func(db *gg.DB) *gg.DB {
		ret := db
		if search != "" {
			ret = ret.Where("name like ?", fmt.Sprintf("%%%v%%", search))
		}
		if filter == nil {
			return ret
		}
		ret = ret.Scopes(kitgorm.BuildStringFilter("`id`", filter.Id))
		ret = ret.Scopes(kitgorm.BuildStringFilter("`name`", filter.Name))
		ret = ret.Scopes(kitgorm.BuildStringFilter("`region`", filter.Region))
		return ret
	}
}

func (g *TenantRepo) FindByIdOrName(ctx context.Context, idOrName string) (*biz.Tenant, error) {
	//parse
	if idOrName == "" {
		return nil, nil
	}
	//parse uuid
	id, err := uuid.Parse(idOrName)
	var tDb biz.Tenant
	if err == nil {
		//id
		err = g.GetDb(ctx).Model(&biz.Tenant{}).Scopes(g.BuildDetailScope(true)).Where("id = ?", id.String()).First(&tDb).Error
	} else {
		err = g.GetDb(ctx).Model(&biz.Tenant{}).Scopes(g.BuildDetailScope(true)).Where("name = ?", idOrName).First(&tDb).Error
	}
	if err != nil {
		if errors.Is(err, gg.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tDb, err
}

func (g *TenantRepo) UpdateAssociation(ctx context.Context, entity *biz.Tenant) error {
	if entity.Conn != nil {
		if err := g.GetDb(ctx).Model(entity).Association("Conn").Replace(entity.Conn); err != nil {
			return err
		}
	} else {
		if err := g.GetDb(ctx).Model(entity).Association("Conn").Clear(); err != nil {
			return err
		}
	}
	if entity.Features != nil {
		if err := g.GetDb(ctx).Model(entity).Association("Features").Replace(entity.Features); err != nil {
			return err
		}
	} else {
		if err := g.GetDb(ctx).Model(entity).Association("Features").Clear(); err != nil {
			return err
		}
	}
	return nil
}
