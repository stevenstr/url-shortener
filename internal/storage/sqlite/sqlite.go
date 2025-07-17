package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	// По сути путь к функции. Проще врапить ошибки и информативнее при дебагинге
	// op - operation или fn - function
	// иногда добавляется в логгер
	const op = "storage.sqlite.New"

	// создаем БД
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// готовим запрос
	stmt, err := db.Prepare(`CREATE TABLE IF NOT EXISTS url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	)`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// создаем таблицу и индекс
	if _, err = stmt.Exec(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
