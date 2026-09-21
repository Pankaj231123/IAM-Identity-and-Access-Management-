package hash

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

func argon2id(argonMemory, argonIterations, argonSaltLength, argonKeyLength uint32, argonParallelism uint8, password string) (string, error) {

}
