package article

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/clairejyu/go-blog/internal/app/blog/user"
	"github.com/clairejyu/go-blog/internal/pkg/db"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
)

var (
	sqlDb    *sql.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	var err error
	// Open connection to the test database.
	// Do NOT import fixtures in a production database!
	// Existing data would be deleted.
	db.Init()
	InitDB(db.D)
	user.InitDB(db.D)

	fixtures, err = testfixtures.New(
		testfixtures.Database(db.SqlDB),    // You database connection
		testfixtures.Dialect("postgresql"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("fixtures"), // the directory containing the YAML files
	)
	if err != nil {
		fmt.Println("testfixtures.New error", err)
	}

	os.Exit(m.Run())
}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		fmt.Println("prepareTestDatabase error", err)
	}
}

func TestGetById(m *testing.T) {
	prepareTestDatabase()
	GetById("1")
}
