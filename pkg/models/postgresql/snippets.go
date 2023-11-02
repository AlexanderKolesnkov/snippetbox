package postgresql

import (
	"database/sql"
	"errors"

	"pilrugen.com/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (id, title, content, created, expires)
	VALUES(COALESCE((SELECT MAX(id+1) FROM snippets),1), $1, $2, CURRENT_TIMESTAMP(0), CURRENT_TIMESTAMP(0) + ($3 || ' DAY')::INTERVAL) RETURNING id`

	// _, err := m.DB.Exec(stmt, title, content, expires)
	// if err != nil {
	// 	return 0, err
	// }
	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP(0) AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNotRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP(0) ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []*models.Snippet

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
