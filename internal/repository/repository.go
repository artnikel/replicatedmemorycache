package repository

type MapDataRepository struct {
	dataMap map[string]string
}

func NewMapDataRepository() *MapDataRepository {
	return &MapDataRepository{
		dataMap: make(map[string]string),
	}
}

func (r *MapDataRepository) Set(key, value string) error {
	r.dataMap[key] = value
	return nil
}

func (r *MapDataRepository) Get(key string) (string, error) {
	value, ok := r.dataMap[key]
	if !ok {
		return "", nil
	}
	return value, nil
}
