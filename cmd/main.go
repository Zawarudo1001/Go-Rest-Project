package main

type salary int
type sample struct {
	name   string
	salary int
}

type database interface {
	isIn(s sample) (bool, error)
	add(s sample) error
	remove(s sample) error
	update(s sample) error
}

type databaseImpl1 struct {
	records map[string]sample
}

type databaseImpl2 struct {
	records []sample
}

func newDbImpl1() *databaseImpl1 {
	return &databaseImpl1{
		records: make(map[string]sample),
	}
}

func newDbImpl2() *databaseImpl2 {
	return &databaseImpl2{
		records: make([]sample, 0),
	}
}

func isInImpl1(db *databaseImpl1, s sample) (bool, error) {
	_, exists := db.records[s.name]
	return exists, nil
}

func isInImpl2(db *databaseImpl2, s sample) (bool, error) {
	for _, record := range db.records {
		if record.name == s.name {
			return true, nil
		}
	}
	return false, nil
}

func addImpl1(db *databaseImpl1, s sample) error {
	db.records[s.name] = s
	return nil
}

func addImpl2(db *databaseImpl2, s sample) error {
	db.records = append(db.records, s)
	return nil
}

func removeImpl1(db *databaseImpl1, s sample) error {
	delete(db.records, s.name)
	return nil
}

func removeImpl2(db *databaseImpl2, s sample) error {
	for i, record := range db.records {
		if record.name == s.name {
			db.records = append(db.records[:i], db.records[i+1:]...)
			break
		}
	}
	return nil
}

func updateImpl1(db *databaseImpl1, s sample) error {
	db.records[s.name] = s
	return nil
}

func updateImpl2(db *databaseImpl2, s sample) error {
	for i, record := range db.records {
		if record.name == s.name {
			db.records[i] = s
			break
		}
	}
	return nil
}

func main() {

}
