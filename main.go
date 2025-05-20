package main

import (
    "context"
    "fmt"
    "log"

    "vpc-orchestrator/cmd"
)

func main() {
    if err := cmd.Execute(context.Background()); err != nil {
        log.Fatalf("Error: %v", err)
    }
}