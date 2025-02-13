//go:build integration
// +build integration

package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/akmalsulaymonov/production-service-go/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		fmt.Println("testing the creation of comment")
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test update comment", func(t *testing.T) {
		fmt.Println("testing the edition of comment")
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		updatedCmt, err := db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug:   "slug updated",
			Author: "author updated",
			Body:   "body updated",
		})
		assert.NoError(t, err)
		assert.Equal(t, "slug updated", updatedCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		fmt.Println("testing the deletion of comment")
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug 1",
			Author: "author 1",
			Body:   "body 1",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
