package config

import (
	"flag"
	"go-one-auth/internal/storage"
)

// Flag struct for parsing from env and cmd args.
type Flag struct {
	DSN     *string
	Storage *string

	HTTP  *string
	HTTPS *string

	FastHTTP  *string
	FastHTTPS *string

	GRPC            *bool
	GRPCCredentials *string
}

// Config struct for storing config values.
type Config struct {
	HTTP  string
	HTTPS string

	FastHTTP  string
	FastHTTPS string

	GRPC                bool
	GRPCWithCredentials string

	DBConfig *storage.Config
}

var f Flag

func init() {
	f.Storage = flag.String("s", "sqlite3", "-s=sqlite3|grpcis")
	f.DSN = flag.String("d", "", "-d=connection_string")

	f.HTTP = flag.String("http", "", "-http=IP:PORT to Run HTTP server")
	f.HTTPS = flag.String("https", "", "-https=IP:PORT to Run HTTPS server")

	f.FastHTTP = flag.String("fhttp", "", "-fhttp=IP:PORT to Run FastHTTP server")
	f.FastHTTPS = flag.String("fhttps", "", "-fhttps=IP:PORT to Run FastHTTPS server")

	f.GRPC = flag.Bool("grpc", false, "-grpc to Run GRPC server")
	f.GRPCCredentials = flag.String("grpc_creds", "", "-grpc_creds to Run GRPC server with creds")
}

func New() *Config {
	flag.Parse()
	return &Config{
		HTTP:  *f.HTTP,
		HTTPS: *f.HTTPS,

		FastHTTP:  *f.FastHTTP,
		FastHTTPS: *f.FastHTTPS,

		GRPC:                *f.GRPC,
		GRPCWithCredentials: *f.GRPCCredentials,

		DBConfig: &storage.Config{
			Type:           *f.Storage,
			DataSourceCred: *f.DSN,
		},
	}
}
