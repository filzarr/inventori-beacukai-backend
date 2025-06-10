package service

import (
	"context"
	"inventori-beacukai-backend/internal/module/master/entity"
)

func (s *masterService) GetBcDocuments(ctx context.Context, req *entity.GetBcDocumentsReq) (*entity.GetBcDocumentsResp, error) {
	return s.repo.GetBcDocuments(ctx, req)
}

func (s *masterService) GetBcDocument(ctx context.Context, req *entity.GetBcDocumentReq) (*entity.GetBcDocumentResp, error) {
	return s.repo.GetBcDocument(ctx, req)
}

func (s *masterService) CreateBcDocument(ctx context.Context, req *entity.CreateBcDocumentReq) (*entity.CreateBcDocumentResp, error) {
	return s.repo.CreateBcDocument(ctx, req)
}

func (s *masterService) UpdateBcDocument(ctx context.Context, req *entity.UpdateBcDocumentReq) error {
	return s.repo.UpdateBcDocument(ctx, req)
}

func (s *masterService) DeleteBcDocument(ctx context.Context, req *entity.DeleteBcDocumentReq) error {
	return s.repo.DeleteBcDocument(ctx, req)
}
