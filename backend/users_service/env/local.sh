export POSTGRES_DB_LOGIN='admin'
export POSTGRES_DB_PASSWORD='pass'
export POSTGRES_DB_HOST='user_db'
export POSTGRES_DB_PORT=5434
export POSTGRES_DB_NAME='user_db'
export MAX_CONNECTIONS=1
export CONNECTION_IDLE_TIME_SEC=10

export POSTGRESQL_URL="postgres://${POSTGRES_DB_LOGIN}:${POSTGRES_DB_PASSWORD}@${POSTGRES_DB_HOST}:${POSTGRES_DB_PORT}/${POSTGRES_DB_NAME}?sslmode=disable"

export GRPC_HOST='0.0.0.0'
export GRPC_PORT=9001
# Заменить на GIN_MODE=release во время деплоя
export GIN_MODE=release


