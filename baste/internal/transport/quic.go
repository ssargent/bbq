package transport

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"

	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
)

type QUIC struct {
	logger   *zap.Logger
	addr     string
	certPath string
	keyPath  string
}

func (s *QUIC) ListenAndServe(ctx context.Context) error {
	s.logger.Info("quic starting", zap.String("addr", s.addr))

	cert, err := tls.LoadX509KeyPair(s.certPath, s.keyPath)
	if err != nil {
		return fmt.Errorf("quic: failed to load certificates %w", err)
	}

	tlsCfg := tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"go-quickly"},
		ClientAuth:   tls.NoClientCert,
	}

	ln, err := quic.ListenAddr(s.addr, &tlsCfg, nil)
	if err != nil {
		return fmt.Errorf("quic: failed to listen %w", err)
	}

	s.logger.Info("quic listening", zap.String("addr", s.addr))

	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		conn, err := ln.Accept(ctx)
		switch err {
		case nil:
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.connection(ctx, conn)
			}()
		case context.Canceled:
			return err
		default:
			s.logger.Error("accept", zap.String("addr", s.addr), zap.Error(err))
		}
	}
}

func (s *QUIC) connection(ctx context.Context, conn quic.Connection) {
	addressAttribute := zap.String("remote-addr", conn.RemoteAddr().String())
	s.logger.Info("connection", addressAttribute)
	defer func() {
		conn.CloseWithError(0, "bye")
		s.logger.Info("closed", addressAttribute)
	}()

	stream, err := conn.OpenStream()
	if err != nil {
		s.logger.Error("error opening stream", addressAttribute, zap.Error(err))
	}

	s.stream(ctx, stream)
}

func (s *QUIC) stream(ctx context.Context, stream quic.Stream) {
	defer stream.Close()

	//stream.
}
