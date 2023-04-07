package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	v1 "github.com/ugent-library/people/api/v1"
	"github.com/ugent-library/people/grpc_client"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	getCmd.Flags().String("id", "", "identifier")
	suggestCmd.Flags().String("query", "", "query")
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(allCmd)
	rootCmd.AddCommand(reindexPersonCmd)
	rootCmd.AddCommand(suggestCmd)
}

func openClient(cb func(client v1.PeopleClient) error) error {
	clientConfig := grpc_client.Config{
		Username: config.Api.Username,
		Password: config.Api.Password,
		Host:     config.Api.Host,
		Port:     config.Api.Port,
		Insecure: config.ApiClient.Insecure,
		CaCert:   config.ApiClient.CaCert,
	}
	return grpc_client.Open(clientConfig, cb)
}

var getCmd = &cobra.Command{
	Use: "get [id]",
	Run: func(cmd *cobra.Command, args []string) {

		id, iErr := cmd.Flags().GetString("id")

		if iErr != nil {
			log.Fatal("no id given")
		}
		if id == "" {
			log.Fatal("no id given")
		}

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetPersonRequest{Id: id}
			res, err := c.GetPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			person := res.GetPerson()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				UseProtoNames:   true,
			}

			bytes, err := marshaller.Marshal(person)

			if err != nil {
				return err
			}

			os.Stdout.Write(bytes)

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var allCmd = &cobra.Command{
	Use: "all",
	Run: func(cmd *cobra.Command, args []string) {

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.GetAllPersonRequest{}
			stream, err := c.GetAllPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			marshaller := protojson.MarshalOptions{
				EmitUnpopulated: true,
				// alternative to protobuf field option "json_name"
				UseProtoNames: true,
			}

			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}

				if err != nil {
					if st, ok := status.FromError(err); ok {
						return errors.New(st.Message())
					}

					return err
				}

				if p := res.GetPerson(); p != nil {
					bytes, _ := marshaller.Marshal(p)
					fmt.Printf("%s\n", string(bytes))
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var reindexPersonCmd = &cobra.Command{
	Use: "reindex",
	Run: func(cmd *cobra.Command, args []string) {

		err := openClient(func(c v1.PeopleClient) error {

			req := v1.ReindexPersonRequest{}
			stream, err := c.ReindexPerson(context.Background(), &req)

			if err != nil {
				return err
			}

			stream.CloseSend()

			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}

				if err != nil {
					if st, ok := status.FromError(err); ok {
						return errors.New(st.Message())
					}

					return err
				}

				if m := res.GetMessage(); m != "" {
					fmt.Printf("%s\n", m)
				}
			}

			return nil
		})

		if err != nil {
			log.Fatal(err)
		}
	},
}

var suggestCmd = &cobra.Command{
	Use: "suggest",
	Run: func(cmd *cobra.Command, args []string) {

		query, err := cmd.Flags().GetString("query")
		if err != nil {
			log.Fatal("no query given")
		}
		if query == "" {
			log.Fatal("no query given")
		}

		persons, err := Services().PersonSearchService.Suggest(query)
		if err != nil {
			log.Fatal(err)
		}

		marshaller := protojson.MarshalOptions{
			EmitUnpopulated: true,
			// alternative to protobuf field option "json_name"
			UseProtoNames: true,
		}

		for _, person := range persons {
			bytes, _ := marshaller.Marshal(person)
			os.Stdout.Write(bytes)
			os.Stdout.Write([]byte("\n"))
		}
	},
}
