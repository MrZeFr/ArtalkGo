package artransfer

import (
	"github.com/ArtalkJS/ArtalkGo/model"
)

var ArtalkV1Importer = &_ArtalkV1Importer{
	ImporterInfo: ImporterInfo{
		Name: "artalk_v1",
		Desc: "从 Artalk v1 (PHP) 导入数据",
		Note: "",
	},
}

type _ArtalkV1Importer struct {
	ImporterInfo
}

func (imp *_ArtalkV1Importer) Run(basic *BasicParams, payload []string) {
	err := RequiredBasicTargetSite(basic)
	if err != nil {
		logFatal(err)
		return
	}

	// 读取文件
	jsonStr, jErr := JsonFileReady(payload)
	if jErr != nil {
		logFatal(jErr)
		return
	}

	var aComments []ArtalkV1CommentAFS
	dErr := JsonDecodeFAS(jsonStr, &aComments)
	if dErr != nil {
		logFatal(dErr)
		return
	}

	tp := []model.Artran{}
	for _, tc := range aComments {
		tp = append(tp, model.Artran{
			ID:          tc.ID,
			Rid:         tc.Rid,
			Content:     tc.Content,
			UA:          tc.UA,
			IP:          tc.IP,
			IsCollapsed: tc.IsCollapsed,
			IsPending:   tc.IsPending,
			CreatedAt:   tc.Date,
			UpdatedAt:   tc.Date,
			Nick:        tc.Nick,
			Email:       tc.Email,
			Link:        tc.Link,
			PageKey:     tc.PageKey,
			SiteName:    basic.TargetSiteName,
			SiteUrls:    basic.TargetSiteUrl,
		})
	}

	ImportArtrans(basic, tp)
}

type ArtalkV1CommentAFS struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Nick        string `json:"nick"`
	Email       string `json:"email"`
	Link        string `json:"link"`
	UA          string `json:"ua"`
	PageKey     string `json:"page_key"`
	Rid         string `json:"rid"`
	IP          string `json:"ip"`
	Date        string `json:"date"`
	IsPending   string `json:"is_pending"`
	IsCollapsed string `json:"is_collapsed"`
}