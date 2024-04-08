package service

type DataRepository interface {
	Set(key, value string) error
	Get(key string) (string, error)

}

type MapDataService struct {
	repository         DataRepository
	replicationService DataReplicationService
}

func NewMapDataService(repository DataRepository, replicationService DataReplicationService) *MapDataService {
	return &MapDataService{
		repository:         repository,
		replicationService: replicationService,
	}
}

func (s *MapDataService) Set(key, value string) error {
	if err := s.repository.Set(key, value); err != nil {
		return err
	}

	if err := s.replicationService.ReplicateData(key, value); err != nil {
		return err
	}
	return nil
}

func (s *MapDataService) Get(key string) (string, error) {
	return s.repository.Get(key)
}
