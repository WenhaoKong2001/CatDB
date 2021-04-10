package CatDB

type DB struct {
}

func Open(filePath string) (*DB, error) {
	return nil, nil
}

//
func (db *DB) Put(key string, value string) error {
	return nil
}

//
func (db *DB) Get(key string) (string, bool, error) {
	return "", false, nil
}

//
func (db *DB) Delete(key string) (bool, error) {
	return false, nil
}
