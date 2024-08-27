package config

import (
	"time"

	"golang-rpc/internal/grpc/proto"
	"os"
)

var (
	// topic and subscription-related variables
	TopicName           = "/event/GRPCTeste__e"
	ReplayPreset        = proto.ReplayPreset_EARLIEST
	ReplayId     []byte = nil
	Appetite     int32  = 5

	// gRPC server variables
	GRPCEndpoint    = "api.pubsub.salesforce.com:7443"
	GRPCDialTimeout = 5 * time.Second
	GRPCCallTimeout = 5 * time.Second

	// OAuth header variables
	GrantType    string = "password"
	ClientId     string
	ClientSecret string
	Username     string
	Password     string

	// OAuth server variables
	OAuthEndpoint    string
	OAuthDialTimeout = 5 * time.Second
)

func Load() {
	Username = os.Getenv("ORG_USERNAME")
	Password = os.Getenv("ORG_PASSWORD")
	OAuthEndpoint = os.Getenv("OAUTH_ENDPOINT")
	ClientId = os.Getenv("APP_CLIENT_ID")
	ClientSecret = os.Getenv("APP_CLIENT_SECRET")
	OAuthEndpoint = os.Getenv("OAUTH_ENDPOINT")
}
