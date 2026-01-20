package main

import(
	"fmt"
	"errors"
	"context"
	"database/sql"
	"github.com/rgarcia2304/aggreGator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddfeeds(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 4{
		return errors.New("Name and URL of feed must be passed")
	}

	ctx := context.Background()


	//connect that user to the feed
	insertFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: sql.NullString{String: cmd.args[2], Valid: true},
		Url: sql.NullString{String: cmd.args[3], Valid: true},
		UserID: user.ID,
	})
	if err != nil{
		return err
	}

	//connect that user to the feed
	_, follow_err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: insertFeed.ID,
	})

	if follow_err != nil{
		return errors.New("Issue with following newly added feed")
	}

	fmt.Println("Feed " + insertFeed.Name.String + " created.")
	
	return nil
}

func handlerListFeeds(s *state, cmd command) error{
	if len(cmd.args) != 2{
		return errors.New("You can not pass arguments to the feeds command")
	}

	ctx := context.Background()
	//check if the user exists in the database
	feedLst, err := s.db.GetFeeds(ctx)
	if err != nil{
		return errors.New("Issue with executing command")
	}
	
	for _, val := range feedLst{
		name := val.Name.String
		url := val.Url.String

		//get the user who created the query 
		usr, err := s.db.GetUserByID(ctx, val.UserID)
		if err != nil{
			return errors.New("The user was not found in the database")
		}

		result := fmt.Sprintf("Name of Feed: %v , Url of Feed: %v, Name of Feed Creator: %v", name, url, usr.Name.String)
		fmt.Println(result)
	}
	return nil

}

func parsePubDate(s string) (time.Time, error) {
    layouts := []string{
        time.RFC1123Z,
        time.RFC1123,
        time.RFC822Z,
        time.RFC822,
        time.RFC3339Nano,
        time.RFC3339,
    }
    var lastErr error
    for _, l := range layouts {
        if t, err := time.Parse(l, s); err == nil {
            return t, nil
        } else {
            lastErr = err
        }
    }
    return time.Time{}, lastErr
}
func scrapeFeeds(s *state) error{
	//fetch the next feed from the database
	feed, err := s.db.GetNextFeedToFetch(context.Background()) 
	if err != nil{
		return errors.New("Could not fetch any feeds")
	}
	fmt.Println(feed.Name.String)
	//mark feed as fetched
	newFeed, err := s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: time.Now(),
		ID: feed.ID,
	})

	feedRsp, err := fetchFeed(context.Background(), newFeed.Url.String)
	
	if err != nil{
		return errors.New("Issue getting the feed response")
	}
	
	fmt.Println("Length of Channel Item")
	fmt.Println(len(feedRsp.Channel.Item))
	fmt.Println("")
	//loop through all the posts in the feed 
	for _, post := range feedRsp.Channel.Item{
	
		pub, err := parsePubDate(post.PubDate)
		if err != nil {
			return fmt.Errorf("bad pubDate %q: %w", post.PubDate, err)
		}
		
		//print("\n")
		//fmt.Println("Now Browsing " + post.Title)
		//fmt.Println("Description " + post.Description)
		//fmt.Println("Link " + post.Link)
		//print("\n")

		newPost, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: post.Title,
			Url: post.Link,
			Description: sql.NullString{String: post.Description, Valid: true},
			PublishedAt: pub,
			FeedID: newFeed.ID,
		})
		if err != nil{
			fmt.Println(err)
		}else{
			fmt.Println("\n")
			fmt.Println(newPost.Title)
			fmt.Println(newPost.Url)
			fmt.Println(newPost.Description.String)
			fmt.Println("\n")
		}
	}
	return  nil

}
