CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created TIMESTAMP NOT NULL,
    expires TIMESTAMP NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);

-- Добавляем несколько тестовых записей
INSERT INTO snippets (id, title, content, created, expires) VALUES (
    COALESCE((SELECT MAX(id+1) FROM snippets),1),
    'Не имей сто рублей',
    E'Не имей сто рублей,\nа имей сто друзей.',
    CURRENT_TIMESTAMP(0),
    CURRENT_TIMESTAMP(0) + INTERVAL '365' DAY
);
 
INSERT INTO snippets (id, title, content, created, expires) VALUES (
    COALESCE((SELECT MAX(id+1) FROM snippets),1),
    'Лучше один раз увидеть',
    E'Лучше один раз увидеть,\nчем сто раз услышать.',
    CURRENT_TIMESTAMP(0),
    CURRENT_TIMESTAMP(0) + INTERVAL '365' DAY
);
 
INSERT INTO snippets (id, title, content, created, expires) VALUES (
    COALESCE((SELECT MAX(id+1) FROM snippets),1),
    'Не откладывай на завтра',
    E'Не откладывай на завтра,\nчто можешь сделать сегодня.',
    CURRENT_TIMESTAMP(0),
    CURRENT_TIMESTAMP(0) + INTERVAL '7' DAY
);

SELECT * FROM snippets;

CREATE USER weblocalhost;
GRANT SELECT, INSERT, UPDATE ON snippets TO weblocalhost;