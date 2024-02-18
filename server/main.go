package main

import (
	"context"
	"encoding/json"
	"golang_grpc/student"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sync"
)

type DataStudentServer struct {
	student.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*student.Student
}

func (d *DataStudentServer) FindStudentByEmail(ctx context.Context, student *student.Student) (*student.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}

	return nil, nil
}

func (d *DataStudentServer) loadData() {
	data, err := os.ReadFile("data/datas.json")
	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error in unmarshal data json", err.Error())
	}
}

func newServer() *DataStudentServer {
	s := DataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error in listen", err.Error())
	}
	grpcServer := grpc.NewServer()
	student.RegisterDataStudentServer(grpcServer, newServer())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error when serve grpc", err.Error())
	}
}
