package routines

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Concurrency struct {
	DB *gorm.DB
	mu sync.Mutex
}

func NewConcurrency(DB *gorm.DB) *Concurrency {
	return &Concurrency{
		DB: DB,
	}
}

func (c *Concurrency) GetConcurrency() {
	ticker := time.NewTicker(4 * time.Minute)

	go func() {
		for range ticker.C {
			c.mu.Lock()
			if err := c.DB.Exec(
				`UPDATE users
            	SET is_blocked = false
            	WHERE id IN (
            	SELECT users_id
            	FROM user_block_infos
            	WHERE user_block_infos.users_id = users.id
            	AND user_block_infos.block_until < NOW()
             );

            DELETE FROM user_block_infos
			WHERE users_id IN (
    		SELECT id
    		FROM users
    		WHERE users.id = user_block_infos.users_id
    		AND users.is_blocked = false
			);`).Error; err != nil {
				fmt.Println(err)
			}

			c.mu.Unlock()

			fmt.Println("worked")
		}
	}()
}
