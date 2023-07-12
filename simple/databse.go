package simple

// implementasi multiple binding

type Database struct {
	Name string
}

// biar tidak error: "provider has multiple parameters of type *github.com/dickidarmawansaputra/belajar-go-dependency-injection/simple.Database"
// perlu membuat alias tipe data yang sama
type PostgreSQL Database
type MongoDB Database

type DatabaseRepository struct {
	DatabasePostgreSQL *PostgreSQL
	DatabaseMongoDB    *MongoDB
}

func NewDatabasePostgreSQL() *PostgreSQL {
	return &PostgreSQL{Name: "PostgreSQL"}
}

func NewDatabaseMongoDB() *MongoDB {
	return &MongoDB{Name: "MongoDB"}
}

func NewDatabaseRepository(databasePostgreSQL *PostgreSQL, databaseMongoDB *MongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: databasePostgreSQL, DatabaseMongoDB: databaseMongoDB}
}
