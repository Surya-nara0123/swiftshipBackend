package helperfunction

import (
	"hash/fnv"

	"github.com/google/uuid"
)

func GenerateUniqueInt() int64 {
	u := uuid.New()
	h := fnv.New32a()
	h.Write(u[:])
	return int64(h.Sum32())
}
