package transport

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"sync"

	"github.com/quic-go/quic-go"
	"go.uber.org/zap"
)

const QuicProtos = "baste-bbq-service"

type QUIC struct {
	logger      *zap.Logger
	addr        string
	certPath    string
	keyPath     string
	temperature func()
	metadata    func()
}

func (s *QUIC) ListenAndServe(ctx context.Context) error {
	s.logger.Info("quic starting", zap.String("addr", s.addr))

	/*	cert, err := tls.LoadX509KeyPair(s.certPath, s.keyPath)
		if err != nil {
			return fmt.Errorf("quic: failed to load certificates %w", err)
		}
	*/
	tlsCfg := generateTLSConfig()

	ln, err := quic.ListenAddr(s.addr, tlsCfg, nil)
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

func (s *QUIC) Connect() (quic.Connection, error) {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{QuicProtos},
	}
	conn, err := quic.DialAddr(s.addr, tlsConf, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
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

	s.controlStream(ctx, conn, stream)
}

func (s *QUIC) controlStream(ctx context.Context, conn quic.Connection, stream quic.Stream) {
	defer stream.Close()

	/*
		Listen to the stream

		Messages will tell us to accept streams for temp data or metadata
	*/
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{QuicProtos},
	}
}
