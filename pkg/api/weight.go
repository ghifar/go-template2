package api

type WeightService interface {
}

type WeightRepository interface {
}

type weightService struct {
	storage WeightRepository
}

func NewWeightService(repo WeightRepository) WeightService {
	return &weightService{
		storage: repo,
	}
}
