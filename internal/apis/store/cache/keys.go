package cache

import (
	"fmt"
	"time"
)

const (
	QuestionsCacheScheme = "questions:%s:%s:%d"
	CategoryCacheScheme  = "category:%d"
)

func key(language, difficulty string, categoryID int32) string {
	return fmt.Sprintf(QuestionsCacheScheme, language, difficulty, categoryID)
}

func keyList(language string, difficulties []string, categoryIDs []int32) []string {
	keys := make([]string, len(difficulties)+len(categoryIDs))

	for i, d := range difficulties {
		for j, id := range categoryIDs {
			keys[i+j] = key(language, d, id)
		}
	}

	return keys
}

func destinationKey() string {
	return fmt.Sprintf("temp:%d", time.Now().UnixMilli())
}

func categoryKey(categoryID int32) string {
	return fmt.Sprintf(CategoryCacheScheme, categoryID)
}
