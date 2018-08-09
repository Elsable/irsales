package mongo

import mgo "gopkg.in/mgo.v2"

type MockDB struct {
}

func (m MockDB) C(name string) Collection {
	return mockCollection{}
}

type mockCollection struct {
}

func (mc mockCollection) Find(query interface{}) *mgo.Query {
	// TODO: Implement mock
	return nil
}

func (mc mockCollection) Count() (n int, err error) {
	// TODO: Implement mock
	return 0, nil
}

func (mc mockCollection) Insert(docs ...interface{}) error {
	return nil
}

func (mc mockCollection) Remove(selector interface{}) error {
	// TODO: Implement mock
	return nil
}

func (mc mockCollection) Update(selector interface{}, update interface{}) error {
	return nil
}
