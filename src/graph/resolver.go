package graph

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Resolver struct {
	DB *pgxpool.Pool
}
