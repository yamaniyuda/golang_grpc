package main

import (
	"context"
	"fmt"
	"golang_grpc/student"
	"google.golang.org/grpc"
	"log"
	"time"
)

func getDataStudentByEmail(client student.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := student.Student{Email: email}
	data, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error when get student by email", err.Error())
	}

	fmt.Println(data)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("Error in dial")
	}
	defer conn.Close()
	client := student.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "budi@gmail.com")
}
