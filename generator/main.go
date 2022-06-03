package main

import (
	"context"
	"generator/project"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Working Directory not valid! %v", err)
	}
	log.Println("Generate...", wd)
	proj := &project.Project{
		Path: wd,
	}
	ctx := context.Background()
	ctx = project.CtxWithProject(ctx, proj)
	_, err = project.ParseControllerFile(ctx, "controller/index.go")
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Println("done")
}
