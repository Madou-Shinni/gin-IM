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
type ANoteRepo interface {
	Create(aNote domain.ANote) error
	Delete(aNote domain.ANote) error
	Update(aNote domain.ANote) error
	Find(aNote domain.ANote) (domain.ANote, error)
	List(page domain.PageANoteSearch) ([]domain.ANote, error)
	Count() (int64, error)
	DeleteByIds(ids request.Ids) error
}

type ANoteService struct {
	repo ANoteRepo
}

func NewANoteService() *ANoteService {
	return &ANoteService{repo: &data.ANoteRepo{}}
}

func (s *ANoteService) Add(aNote domain.ANote) error {
	// 3.持久化入库
	if err := s.repo.Create(aNote); err != nil {
		// 4.记录日志
		logger.Error("s.repo.Create(aNote)", zap.Error(err), zap.Any("domain.ANote", aNote))
		return err
	}

	return nil
}

func (s *ANoteService) Delete(aNote domain.ANote) error {
	if err := s.repo.Delete(aNote); err != nil {
		logger.Error("s.repo.Delete(aNote)", zap.Error(err), zap.Any("domain.ANote", aNote))
		return err
	}

	return nil
}

func (s *ANoteService) Update(aNote domain.ANote) error {
	if err := s.repo.Update(aNote); err != nil {
		logger.Error("s.repo.Update(aNote)", zap.Error(err), zap.Any("domain.ANote", aNote))
		return err
	}

	return nil
}

func (s *ANoteService) Find(aNote domain.ANote) (domain.ANote, error) {
	res, err := s.repo.Find(aNote)

	if err != nil {
		logger.Error("s.repo.Find(aNote)", zap.Error(err), zap.Any("domain.ANote", aNote))
		return res, err
	}

	return res, nil
}

func (s *ANoteService) List(page domain.PageANoteSearch) (response.PageResponse, error) {
	var (
		pageRes response.PageResponse
	)

	data, err := s.repo.List(page)
	if err != nil {
		logger.Error("s.repo.List(page)", zap.Error(err), zap.Any("domain.PageANoteSearch", page))
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

func (s *ANoteService) DeleteByIds(ids request.Ids) error {
	if err := s.repo.DeleteByIds(ids); err != nil {
		logger.Error("s.DeleteByIds(ids)", zap.Error(err), zap.Any("ids request.Ids", ids))
		return err
	}

	return nil
}
