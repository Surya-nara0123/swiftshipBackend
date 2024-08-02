package helperfunction

import (
	"hash/fnv"

	"github.com/google/uuid"
)

func GenerateUniqueInt() int32 {
	u := uuid.New()
	h := fnv.New32a()
	h.Write(u[:])
	return int32(h.Sum32())
}
