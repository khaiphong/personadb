package persondb

import (

)

// DB is a PersonaDB database.
type DB struct {
    // Need 64-bit alignment.
    seq uint64

    // Session.
    s *session

}

func openDB(s *session) (*DB, error) {}

func (db *DB) Close() error {}
