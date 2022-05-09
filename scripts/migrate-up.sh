export POSTGRESQL_URL='postgres://husanmusa:pass@localhost:5432/pro_book?sslmode=disable'

migrate -database ${POSTGRESQL_URL} -path migrations up