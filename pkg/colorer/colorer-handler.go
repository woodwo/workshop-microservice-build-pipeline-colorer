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
	return &GetColorResponse{Cold: 0, Hot: 50}, nil
}

func NewServer() ColorerServer {
	s := &colorerServer{}
	return s
}
