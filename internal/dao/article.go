package dao

import (
	. "gin-demo/internal/models/article"
)

func CreateArticle(m *Article) error {
	return Db.Create(m).Error
}

func GetArticle(id uint) (m Article, err error) {
	err = Db.First(&m, id).Error
	return
}

func SaveArticle(m *Article) error {
	return Db.Save(m).Error
}

func DelArticle(id uint) error {
	return Db.Delete(&Article{}, id).Error
}


func ArticleExist(id uint) (bool, error) {
	var count int64
	if err := Db.Model(&Article{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func ArticleAllExist(ids []uint) (bool, error) {
	var count int64
	if err := Db.Model(&Article{}).Where("id IN ?", ids).Count(&count).Error; err != nil {
		return false, err
	}
	return int(count) == len(ids), nil
}

func ArticleList(page, size int) (list []Article, err error) {
	err = Db.Offset((page - 1) * size).Limit(size).Find(&list).Error
	return
}

func ArticleCount() (count int64, err error) {
	err = Db.Model(&Article{}).Count(&count).Error
	return
}

func ArticlePageList(page, size int) (list []Article, err error) {
    err = Db.Offset((page - 1) * size).Limit(size).Find(&list).Error
	return
}

func ArticlePageCount() (count int64, err error) {
	err = Db.Model(&Article{}).Count(&count).Error
	return
}
