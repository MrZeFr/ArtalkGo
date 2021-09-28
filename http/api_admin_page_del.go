package http

import (
	"github.com/ArtalkJS/ArtalkGo/lib"
	"github.com/ArtalkJS/ArtalkGo/model"
	"github.com/labstack/echo/v4"
)

type ParamsAdminPageDel struct {
	Key      string `mapstructure:"key" param:"required"`
	SiteName string `mapstructure:"site_name"`
	SiteID   uint
}

func ActionAdminPageDel(c echo.Context) error {
	if isOK, resp := AdminOnly(c); !isOK {
		return resp
	}

	var p ParamsAdminPageDel
	if isOK, resp := ParamsDecode(c, ParamsAdminPageDel{}, &p); !isOK {
		return resp
	}

	// find site
	if isOK, resp := CheckSite(c, p.SiteName, &p.SiteID); !isOK {
		return resp
	}

	page := model.FindPage(p.Key, p.SiteName)
	if page.IsEmpty() {
		return RespError(c, "page not found.")
	}

	err := lib.DB.Delete(&page).Error
	if err != nil {
		return RespError(c, "page 删除失败")
	}

	// 删除所有相关内容
	var comments []model.Comment
	lib.DB.Where(&model.Comment{PageKey: p.Key, SiteName: p.SiteName}).Find(&comments)

	tx := lib.DB.Begin()
	for _, c := range comments {
		tx.Delete(&c)
	}
	tx.Commit()

	return RespSuccess(c)
}