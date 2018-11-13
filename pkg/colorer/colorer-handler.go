package colorer

import (
	"log"

	"golang.org/x/net/context"
)

type colorerServer struct {
}

// GetEcho returns the feature at the given point.
func (s *colorerServer) GetColor(ctx context.Context, msg *GetColorRequest) (*GetColorResponse, error) {
	log.Printf("Server colorer called with message (%v)", msg)
        time.Sleep(time.Duration(15) * time.Millisecond)
	return &GetColorResponse{Cold: 0, Hot: 133}, nil
}

func NewServer() ColorerServer {
	s := &colorerServer{}
	return s
}
