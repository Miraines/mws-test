package service

import (
	"context"

	"mws-test/internal/api"
	"mws-test/internal/store"
)

type CatService struct {
	st *store.MemoryStore
}

func NewCatService(st *store.MemoryStore) *CatService {
	return &CatService{st: st}
}

func (s *CatService) ListCats(ctx context.Context) ([]api.Cat, error) {
	return s.st.List(), nil
}

func (s *CatService) CreateCat(ctx context.Context, req *api.NewCat) (*api.Cat, error) {
	c := s.st.Create(*req)
	return &c, nil
}

func (s *CatService) GetCat(
	ctx context.Context,
	params api.GetCatParams,
) (api.GetCatRes, error) {
	c, ok := s.st.Get(params.ID)
	if !ok {
		return &api.GetCatNotFound{}, nil
	}
	return &c, nil
}

func (s *CatService) UpdateCat(
	ctx context.Context,
	req *api.UpdateCat,
	params api.UpdateCatParams,
) (api.UpdateCatRes, error) {
	c, ok := s.st.Update(params.ID, *req)
	if !ok {
		return &api.UpdateCatNotFound{}, nil
	}
	return &c, nil
}

func (s *CatService) DeleteCat(
	ctx context.Context,
	params api.DeleteCatParams,
) (api.DeleteCatRes, error) {
	if !s.st.Delete(params.ID) {
		return &api.DeleteCatNotFound{}, nil
	}
	return &api.DeleteCatNoContent{}, nil
}
