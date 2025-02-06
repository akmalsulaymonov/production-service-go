package main

import (
	"context"
	"fmt"

	"github.com/akmalsulaymonov/production-service-go/internal/comment"
	"github.com/akmalsulaymonov/production-service-go/internal/db"
	transportHttp "github.com/akmalsulaymonov/production-service-go/internal/transport/http"
)

// Rin - is going to be responsible for
// the instantiation and startup fot our
// go application
func Run() error {
	fmt.Println("Starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	// comment service
	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	/*
		// post a commnet
		insCmt, _ := cmtService.PostComment(
			context.Background(), comment.Comment{
				Slug:   "manual-test",
				Author: "Abdulaziz",
				Body:   "Do your homework!",
			},
		)
		fmt.Println(insCmt)
		fmt.Println(insCmt.ID)

		// get comment by ID
		fmt.Println(cmtService.GetComment(context.Background(), "601fbfe6-3089-4e00-93a5-b6a27a785e6e"))
	*/

	return nil
}

func main() {
	fmt.Println("Go REST API")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
