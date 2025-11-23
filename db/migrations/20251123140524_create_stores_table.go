package migrations
// goose mysql "user:password@/dbname?parseTime=true" status
// goose mysql "root:@/retail_chain?parseTime=true" status

import (
	"context"
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

const CREATE_STORES_STATEMENT = `CREATE TABLE stores(id int primary key,
																									name varchar(100),
																									description text,
																									store_code varchar(50),
																									tag_line text,
																									INDEX pk_id(id),
																									INDEX idx_store_code(store_code),
																									INDEX idx_name(name)
																									);`

func init() {
	goose.AddMigrationContext(upCreateDbRetailChain, downCreateDbRetailChain)
}

func upCreateDbRetailChain(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(CREATE_STORES_STATEMENT)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}

func downCreateDbRetailChain(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DROP TABLE stores")
	if err != nil{
		log.Fatal(err)
	}
	return nil
}
