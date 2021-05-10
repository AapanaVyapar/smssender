package helpers

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"sync/atomic"
	"time"
)

func ContextError(ctx context.Context) error {

	switch ctx.Err() {
	case context.Canceled:
		log.Println("Request Canceled")
		return status.Error(codes.DeadlineExceeded, "Request Canceled")
	case context.DeadlineExceeded:
		log.Println("DeadLine Exceeded")
		return status.Error(codes.DeadlineExceeded, "DeadLine Exceeded")
	default:
		return nil
	}
}
func processUniqueBytes() [5]byte {
	var b [5]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return b
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize objectid package with crypto.rand.Reader: %v", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

type UniqueId [12]byte

func (id UniqueId) Hex() string {
	return hex.EncodeToString(id[:])
}

func (id UniqueId) String() string {
	return fmt.Sprintf("ObjectID(%q)", id.Hex())
}

func NewUniqueIdFromTimestamp() UniqueId {
	var b [12]byte

	timestamp := time.Now()

	processUnique := processUniqueBytes()
	objectIDCounter := readRandomUint32()

	binary.BigEndian.PutUint32(b[0:4], uint32(timestamp.Unix()))
	copy(b[4:9], processUnique[:])
	putUint24(b[9:12], atomic.AddUint32(&objectIDCounter, 1))

	return b
}
