package web

import (
	"fmt"
	"math"
	"net"

	log "github.com/sirupsen/logrus"

	"golang.org/x/net/context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/Jaybee233/go-logsink/logsink"
)

// server is used to implement logsink.LogTransferServer.
type server struct {
	hub *hub
}

// SendLine implements logsink.SendLine
func (s *server) SendLine(ctx context.Context, in *pb.LineMessage) (*pb.LineResult, error) {
	numberOfLines.Inc()
	breakAt := viper.GetInt("web.break")
	prio := int32(math.Max(0, math.Min(9, float64(in.Priority))))
	if breakAt == 0 {
		s.broadcastLine(in.Line, prio)
	} else {
		iterations := int(len(in.Line) / breakAt)
		for start := 0; start <= iterations; start++ {
			s.broadcastLine(in.Line[start*breakAt:int32(math.Min(float64((start+1)*breakAt), float64(len(in.Line))))], prio)
		}
	}
	return &pb.LineResult{Result: true}, nil
}

func (s *server) broadcastLine(line string, priority int32) {
	s.hub.broadcast <- []byte(fmt.Sprintf("{ Line: '%s', Priority: %d }", line, priority))
}

func (s *server) run() {
	s.hub = newHub()
	go s.hub.run()
	lis, err := net.Listen("tcp", viper.GetString("web.bind"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterLogTransferServer(srv, s)
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
