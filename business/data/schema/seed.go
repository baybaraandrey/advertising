package schema

import (
	"github.com/jmoiron/sqlx"
)

func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

const seeds = `
INSERT INTO users (uuid,name,email,phone,roles,password_hash,created,updated) VALUES
('c0a6259e-7ac6-4142-92b3-67573e416a97', 'Rob Pike', 'robpike@example.com', '+380955084565', '{user}','$2b$12$dXm3RasFsIh1EqIRsxsf6.QZFgZVOrLcuNwTzXbOZzOYeY5mvWU.q', '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('489deb46-c944-4741-96d1-8977d5272a48', 'Robert De', 'robert@example.com', '+380955084566', '{user}', '$2b$12$9ZaXHXmABm7bfnr0toaZNusEb9WzV/OQIfbfgKk2JfiuznFimyWu.', '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('3fcd5529-8d3b-4fb1-9c25-0e2591ffa5d1', 'John Doe', 'john@example.com', '+380955084567', '{user}', '$2b$12$.ZR1fA6JAg2KUCiLTCQI1.57wvYwVv48rRvs6dStIuevyveUBjzr6', '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('601b7b6e-a99f-462a-99c1-82dba67f5425', 'Andrew Baybara', 'baybaraandrey@gmail.com', '+380955081131', '{admin}', '$2b$12$gkQ3IuMKY3iphN0jSRYxpO..X4dKXpY2WYPk9gWz8Erlxj50Y4K8e', '2021-03-24 00:00:00', '2021-03-24 00:00:00')
ON CONFLICT DO NOTHING;

INSERT INTO categories (uuid,name,created,updated) VALUES
('b8b86492-82f5-4716-954c-b928b72b9206', 'Cars','2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('ea4859f6-ef40-4a1f-a822-3028da728125', 'Toys', '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('da0fad0b-7841-4241-8716-bd294706c84b', 'Computers', '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('7c3680b2-9309-466a-af85-ed9174548014', 'Software', '2021-03-24 00:00:00', '2021-03-24 00:00:00')
ON CONFLICT DO NOTHING;

INSERT INTO adverts (uuid,user_uuid,category_uuid,title,description,location,price,created,updated) VALUES
('51da7b78-f2b7-4fed-88d1-ece8a3a9a554','c0a6259e-7ac6-4142-92b3-67573e416a97', 'b8b86492-82f5-4716-954c-b928b72b9206', 'Title1', 'Description1', 'Kiev', 101, '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('0ceb43ff-2091-4075-ab0a-c10d4442b7d5','489deb46-c944-4741-96d1-8977d5272a48', 'ea4859f6-ef40-4a1f-a822-3028da728125', 'Title2', 'Description2', 'Kiev', 102, '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('c77115eb-b605-479d-bdb1-83d085a6cc54','3fcd5529-8d3b-4fb1-9c25-0e2591ffa5d1', 'da0fad0b-7841-4241-8716-bd294706c84b', 'Title3', 'Description3', 'Kiev', 103, '2021-03-24 00:00:00', '2021-03-24 00:00:00'),
('d7ba41d8-b5ad-46ea-9354-6af79f81471d','601b7b6e-a99f-462a-99c1-82dba67f5425', '7c3680b2-9309-466a-af85-ed9174548014', 'Title4', 'Description4', 'Kiev', 104, '2021-03-24 00:00:00', '2021-03-24 00:00:00')
ON CONFLICT DO NOTHING;
`
