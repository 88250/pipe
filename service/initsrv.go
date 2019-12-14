// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/88250/pipe/model"
	"github.com/88250/pipe/theme"
	"github.com/88250/pipe/util"
	"github.com/jinzhu/gorm"
)

// Init service.
var Init = &initService{
	mutex: &sync.Mutex{},
}

type initService struct {
	mutex  *sync.Mutex
	inited bool
}

// PlatformStatus represents platform status.
type PlatformStatus struct {
	Version string `json:"version"`
	Locale  string `json:"locale"`
	Inited  bool   `json:"inited"`
}

func (srv *initService) Inited() bool {
	if srv.inited {
		return true
	}

	status, err := srv.Status()
	if err != nil {
		return false
	}

	return status.Inited
}

func (srv *initService) Status() (platformStatus *PlatformStatus, err error) {
	platformStatus = &PlatformStatus{
		Version: model.Version,
		Locale:  "zh_CN",
	}

	localeSetting := &model.Setting{}
	if err = db.Where("`name` = ? AND `value` IS NOT NULL AND `blog_id` = ?", model.SettingNameI18nLocale, uint(1)).
		Find(localeSetting).Error; nil != err {
		if gorm.ErrRecordNotFound == err {
			err = nil
			return
		}

		msg := "checks platform init status failed: " + err.Error()
		logger.Errorf(msg)

		return platformStatus, errors.New(msg)
	}

	srv.inited, platformStatus.Inited = true, true
	platformStatus.Locale = localeSetting.Value

	return
}

func (srv *initService) initBlog(tx *gorm.DB, admin *model.User, blogID uint64) error {
	if err := initBlogAdmin(tx, admin, blogID); nil != err {
		return err
	}
	if err := initSystemSettings(tx, blogID); nil != err {
		return err
	}
	if err := initThemeSettings(tx, blogID); nil != err {
		return err
	}
	if err := initBasicSettings(tx, admin, blogID); nil != err {
		return err
	}
	if err := initPreferenceSettings(tx, blogID); nil != err {
		return err
	}
	if err := initSignSettings(tx, blogID); nil != err {
		return err
	}
	if err := initI18nSettings(tx, blogID); nil != err {
		return err
	}
	if err := initFeedSettings(tx, blogID); nil != err {
		return err
	}
	if err := init3rdStatistic(tx, blogID); nil != err {
		return err
	}
	if err := initAd(tx, blogID); nil != err {
		return err
	}
	if err := initStatisticSettings(tx, blogID); nil != err {
		return err
	}
	if err := helloWorld(tx, admin, blogID); nil != err {
		return err
	}

	return nil
}

func (srv *initService) InitBlog(blogAdmin *model.User) error {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	user := User.GetUserByName(blogAdmin.Name)
	if nil != user && nil != User.GetOwnBlog(user.ID) {
		return nil
	}

	blogID := util.CurrentMillisecond()
	tx := db.Begin()
	if err := srv.initBlog(tx, blogAdmin, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

func (srv *initService) InitPlatform(platformAdmin *model.User) error {
	if srv.inited {
		return nil
	}

	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	blogID := uint64(1)

	saCount := 0
	db.Model(&model.Correlation{}).Where("`id1` = ? AND `type` = ? AND `int1` = ? AND `blog_id` = ?",
		blogID, model.CorrelationBlogUser, model.UserRoleBlogAdmin, blogID).
		Count(&saCount)
	if 0 < saCount {
		srv.inited = true

		return nil
	}

	tx := db.Begin()
	if err := srv.initBlog(tx, platformAdmin, blogID); nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	srv.inited = true

	return nil
}

func initBlogAdmin(tx *gorm.DB, admin *model.User, blogID uint64) error {
	admin.Locale = "zh_CN"
	admin.TotalArticleCount = 1 // article "Hello, World!"

	exist := &model.User{}
	tx.Where("`name` = ?", admin.Name).First(exist)
	admin.ID = exist.ID
	admin.CreatedAt = exist.CreatedAt

	if err := tx.Save(admin).Error; nil != err {
		return err
	}

	blogUser := &model.Correlation{
		ID1:    blogID,
		ID2:    admin.ID,
		Type:   model.CorrelationBlogUser,
		Int1:   model.UserRoleBlogAdmin,
		Int2:   1, // article "Hello, World!"
		BlogID: blogID,
	}
	if err := tx.Create(blogUser).Error; nil != err {
		return err
	}

	return nil
}

func helloWorld(tx *gorm.DB, admin *model.User, blogID uint64) error {
	content := `![Hello](` + util.ImageSize(util.RandImage(), 768, 432) + `)

Pipe 博客平台已经初始化完毕，可在管理后台 - 设置 - 基础设置中调整更多细节。如果需要导入已有博客文章，请参考文档 [Hexo/Jekyll/Markdown 文件导入](https://hacpai.com/article/1498490209748)。

另外，出于安全考虑请尽快完成如下操作：

1. 登录[社区](https://hacpai.com)
2. 在社区[个人设置 - B3](https://hacpai.com/settings/b3) 中更新 B3 Key
3. 在 Pipe 管理后台 - 设置 - 账号中也进行同样的 B3 Key 更新

最后，如果你觉得 Pipe 很赞，请到[项目主页](https://github.com/88250/pipe)给颗星鼓励一下 :heart:`

	now := time.Now()
	article := &model.Article{
		AuthorID:     admin.ID,
		Title:        "世界，你好！",
		Tags:         "Pipe",
		Content:      content,
		Path:         "/hello-world",
		Status:       model.ArticleStatusOK,
		Topped:       false,
		Commentable:  true,
		CommentCount: 1,
		BlogID:       blogID,
	}
	article.CreatedAt = now
	article.UpdatedAt = now
	article.PushedAt = now
	article.ID = util.CurrentMillisecond()
	if err := tx.Create(article).Error; nil != err {
		return err
	}

	tag := &model.Tag{
		Title:        "Pipe",
		ArticleCount: 1,
		BlogID:       blogID,
	}
	if err := tx.Create(tag).Error; nil != err {
		return err
	}

	articleTagRel := &model.Correlation{
		ID1:    article.ID,
		ID2:    tag.ID,
		Type:   model.CorrelationArticleTag,
		BlogID: blogID,
	}
	if err := tx.Create(articleTagRel).Error; nil != err {
		return err
	}

	if err := Archive.ArchiveArticleWithoutTx(tx, article); nil != err {
		return err
	}

	comment := &model.Comment{
		ArticleID: article.ID,
		AuthorID:  admin.ID,
		Content:   "相信积累后必然会有收获 :smile:",
		BlogID:    blogID,
	}
	comment.CreatedAt = now
	comment.UpdatedAt = now
	comment.PushedAt = now
	comment.ID = util.CurrentMillisecond()
	if err := tx.Create(comment).Error; nil != err {
		return err
	}

	return nil
}

func initSystemSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategorySystem,
		Name:     model.SettingNameSystemVer,
		Value:    model.Version,
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initThemeSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryTheme,
		Name:     model.SettingNameThemeName,
		Value:    theme.DefaultTheme,
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initBasicSettings(tx *gorm.DB, blogAdmin *model.User, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicBlogURL,
		Value:    model.Conf.Server + util.PathBlogs + "/" + blogAdmin.Name,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicBlogSubtitle,
		Value:    "记录精彩的程序人生",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicBlogTitle,
		Value:    blogAdmin.Name + " 的博客",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicCommentable,
		Value:    "true",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicFooter,
		Value:    model.SettingBasicFooterDefault,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicHeader,
		Value:    model.SettingBasicHeaderDefault,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicMetaDescription,
		Value:    "小而美的 golang 博客平台",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicMetaKeywords,
		Value:    "Pipe,golang,博客,开源",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicNoticeBoard,
		Value:    model.SettingBasicBasicNoticeBoardDefault,
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicFaviconURL,
		Value:    "https://img.hacpai.com/pipe.ico",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryBasic,
		Name:     model.SettingNameBasicLogoURL,
		Value:    "https://img.hacpai.com/pipe192.png",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initPreferenceSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListPageSize,
		Value:    strconv.Itoa(model.SettingPreferenceArticleListPageSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListWindowSize,
		Value:    strconv.Itoa(model.SettingPreferenceArticleListWindowSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceArticleListStyle,
		Value:    strconv.Itoa(model.SettingPreferenceArticleListStyleDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostCommentArticleListSize,
		Value:    strconv.Itoa(model.SettingPreferenceMostCommentArticleListSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostUseTagListSize,
		Value:    strconv.Itoa(model.SettingPreferenceMostUseTagListSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceMostViewArticleListSize,
		Value:    strconv.Itoa(model.SettingPreferenceMostViewArticleListSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRecentCommentListSize,
		Value:    strconv.Itoa(model.SettingPreferenceRecentCommentListSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryPreference,
		Name:     model.SettingNamePreferenceRecommendArticleListSize,
		Value:    strconv.Itoa(model.SettingPreferenceRecommendArticleListSizeDefault),
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initSignSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategorySign,
		Name:     model.SettingNameArticleSign,
		Value:    model.SettingArticleSignDefault,
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initI18nSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryI18n,
		Name:     model.SettingNameI18nLocale,
		Value:    "zh_CN",
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryI18n,
		Name:     model.SettingNameI18nTimezone,
		Value:    "Asia/Shanghai",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initFeedSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryFeed,
		Name:     model.SettingNameFeedOutputMode,
		Value:    strconv.Itoa(model.SettingFeedOutputModeValueAbstract),
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func init3rdStatistic(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryThirdStatistic,
		Name:     model.SettingNameThirdStatisticBaidu,
		Value:    "",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initStatisticSettings(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticArticleCount,
		Value:    "1", // article "Hello, World!"
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticCommentCount,
		Value:    "1", // comment of article "Hello, World!"
		BlogID:   blogID}).Error; nil != err {
		return err
	}
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryStatistic,
		Name:     model.SettingNameStatisticViewCount,
		Value:    "0",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}

func initAd(tx *gorm.DB, blogID uint64) error {
	if err := tx.Create(&model.Setting{
		Category: model.SettingCategoryAd,
		Name:     model.SettingNameAdGoogleAdSenseArticleEmbed,
		Value:    "",
		BlogID:   blogID}).Error; nil != err {
		return err
	}

	return nil
}
