package main

import (
	"awesomeProject/proto"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

type grpcServer struct {}


func main() {

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen at port tcp:4040, error: %v", err)
	}
	grpcSrv := grpc.NewServer()
	proto.RegisterDetectionCRUDServer(grpcSrv, &grpcServer{})
	reflection.Register(grpcSrv)

	if e := grpcSrv.Serve(listener); e != nil {
		panic(e)
	}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			createQueryType(
				createDetectionType(),
			),
		),
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		GraphiQL:true,
	})
	http.Handle("/graphql", CorsMiddleware(handler))
	log.Println("GraphQL API started at http://localhost:8091/graphql")
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w,r)
	})
}

type DetectionMetadata struct {
	Id          string `json:"id"`
	XCoordinate string `json:"x_coordinate"`
	YCoordinate string `json:"y_coordinate"`
	BodyPart    string `json:"bodyPart"`
	Timestamp   string `json:"timestamp"`
}

func createQueryType(detectionType *graphql.Object) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"detection": &graphql.Field{
			Type: detectionType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("fetching detection metadata with id: %d", v)
				return fetchDetectionFromElastic(v)
			},
		},
	}}
}

func createDetectionType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "DetectionMetadata",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"x_coordinate": &graphql.Field{
				Type: graphql.String,
			},
			"y_coordinate": &graphql.Field{
				Type: graphql.String,
			},
			"bodyPart": &graphql.Field{
				Type: graphql.String,
			},
			"timestamp": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func (s *grpcServer) Get(ctx context.Context, request *proto.DetectionRequest) (*proto.DetectionResponse, error) {
	detection, err := fetchDetectionFromElastic(int(request.GetId()))
	if err != nil {
		log.Fatalf("failed to fetch detection, error: %v", err)
	}
	protoDetection := proto.Detection{Id:detection.Id, XCoordinate:detection.XCoordinate,YCoordinate:detection.YCoordinate,BodyPart:detection.BodyPart,Timestamp:detection.Timestamp}
	return &proto.DetectionResponse{Detection:&protoDetection},nil
}

func fetchDetectionFromElastic(id int) (*DetectionMetadata, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:9200/detection/stream/_search=id:%d", id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s: %s", "Error: ", resp.Status)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error by parsing data")
	}
	result := DetectionMetadata{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.New("error by unmarshal data")
	}
	return &result, nil
}
