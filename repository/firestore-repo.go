package repository

import (
	"context"
	"goTut/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

// NewFirestoreRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "firestore project id"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase/gcp json service account") // initialises database
	client, err := firestore.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebase/gcp json service account")
	client, err := firestore.NewClient(ctx, projectID, opt) // initialises database
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
			return nil, err
		}
		var post entity.Post
		doc.DataTo(&post)
		posts = append(posts, post)
	}
	return posts, nil
}
