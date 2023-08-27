package service

import (
	"gin-IM/internal/data"
	"gin-IM/internal/domain"
	"gin-IM/pkg/request"
	"gin-IM/pkg/response"
	"github.com/Madou-Shinni/go-logger"
	"go.uber.org/zap"
)

// 定义接口
type ArticleRepo interface {
	Create(article domain.Article) error
	Delete(article domain.Article) error
	Update(article domain.Article) error
	Find(article domain.Article) (domain.Article, error)
	List(page domain.PageArticleSearch) ([]domain.Article, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type ArticleService struct {
	repo ArticleRepo
}

func NewArticleService() *ArticleService {
	return &ArticleService{repo: &data.ArticleRepo{}}
}

func (s *ArticleService) Add(article domain.Article) error {
	// 3.持久化入库
	if err := s.repo.Create(article); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(article)", zap.Error(err), zap.Any("domain.Article", article))
		return err
	}

	return nil
}

func (s *ArticleService) Delete(article domain.Article) error {
	if err := s.repo.Delete(article); err != nil {
		logger.Error("s.repo.Delete(article)", zap.Error(err), zap.Any("domain.Article", article))
		return err
	}

	return nil
}

func (s *ArticleService) Update(article domain.Article) error {
	if err := s.repo.Update(article); err != nil {
		logger.Error("s.repo.Update(article)", zap.Error(err), zap.Any("domain.Article", article))
		return err
	}

	return nil
}

func (s *ArticleService) Find(article domain.Article) (domain.Article, error) {
	res, err := s.repo.Find(article)

	if err != nil {
		logger.Error("s.repo.Find(article)", zap.Error(err), zap.Any("domain.Article", article))
		return res, err
	}

	return res, nil
}

func (s *ArticleService) List(page domain.PageArticleSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageArticleSearch", page))
		return pageRes, err
	}

	count, err := s.repo.Count()
	if err != nil {
		logger.Error("s.repo.Count()", zap.Error(err))
		return pageRes, err
	}

	pageRes.Data = data
	pageRes.Total = count

	return pageRes, nil
}

func (s *ArticleService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
