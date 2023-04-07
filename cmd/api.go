package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_server"
)

func init() {
	apiCmd.AddCommand(apiStartCmd)
	apiCmd.AddCommand(apiProxyCmd)
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api [command]",
	Short: "api commands",
}

var apiStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api server",
	Run: func(cmd *cobra.Command, args []string) {

		srv := grpc_server.NewServer(&grpc_server.ServerConfig{
			Logger:              logger,
			PersonService:       Services().PersonService,
			PersonSearchService: Services().PersonSearchService,
			Username:            viper.GetString("api.username"),
			Password:            viper.GetString("api.password"),
			TlsEnabled:          viper.GetBool("api.tls_enabled"),
			TlsServerCert:       viper.GetString("api.tls_server_crt"),
			TlsServerKey:        viper.GetString("api.tls_server_key"),
		})

		addr := fmt.Sprintf("%s:%d", viper.GetString("api.host"), viper.GetInt("api.port"))
		logger.Infof("Listening at %s", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		if err := srv.Serve(listener); err != nil {
			// TODO not everything is a fatal error.
			log.Println(err)
		}
	},
}

/*
POST /api.v1.People/GetAllPerson
POST /api.v1.People/GetPerson with json body '{"id":"<person.id>"}'
*/
var apiProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "start the api proxy server",
	Run: func(cmd *cobra.Command, args []string) {
		openClient(func(c v1.PeopleClient) error {

			// Register gRPC server endpoint
			// Note: Make sure the gRPC server is running properly and accessible
			mux := runtime.NewServeMux()
			err := v1.RegisterPeopleHandlerClient(context.Background(), mux, c)
			if err != nil {
				return err
			}
			addr := fmt.Sprintf("%s:%d", config.ApiProxy.Host, config.ApiProxy.Port)

			// Start HTTP server (and proxy calls to gRPC server endpoint)
			return http.ListenAndServe(addr, mux)
		})
	},
}
