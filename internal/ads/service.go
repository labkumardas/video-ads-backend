package ads

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) FetchAds() ([]Ad, error) {
	return s.repo.GetAllAds()
}

func (s *Service) LogClick(click Click) {
	_ = s.repo.InsertClick(click) // log error later
}

func (s *Service) GetClickStats() (map[int]int, error) {
	return s.repo.GetClickCounts()
}
