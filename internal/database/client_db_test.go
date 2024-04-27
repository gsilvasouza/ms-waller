package database

import (
	"database/sql"
	"testing"

	"github.com/gsilvasouza/ms-waller/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "test",
		Email: "j@j.com",
	}
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuite) TestGet() {
	clientToSave, _ := entity.NewClient("Gabriel", "g@g.com")
	s.clientDB.Save(clientToSave)

	clientDB, err := s.clientDB.Get(clientToSave.ID)
	s.Nil(err)
	s.Equal(clientToSave.ID, clientDB.ID)
	s.Equal(clientToSave.Name, clientDB.Name)
	s.Equal(clientToSave.Email, clientDB.Email)
}
