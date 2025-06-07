package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func ConnectToDb(connStr string) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("Database connection failed: %v", err.Error())
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Failed to ping the database: %v", err.Error())
	}

	return conn, nil
}
