package main

import(
	"fmt"
	"errors"
	"context"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error{
	if len(cmd.args) != 2  && len(cmd.args) != 3{
		return errors.New("Need to pass correct number of paraemeters either none for default or one for varying post return")
	}
	 

	if len(cmd.args) != 3{
		limit := int32(2)
		ctx := context.Background()

		//connect that user to the feed
		posts, err := s.db.GetPost(ctx, limit)
		if err != nil{
			return err
		}
		for _, post := range posts{
			fmt.Println(post.Title)
		res := fmt.Sprintf("Post Title: %v \n Post Description %v \n Post Link %v \n", post.Title, post.Description.String, post.Url)
		fmt.Println(res)	
		}

	}else {
		limit32, err := strconv.ParseInt(cmd.args[2], 10, 32)
		if err != nil{
			return errors.New("The limit passed must be an int")
		}
		//limit := int(limit32)
		ctx := context.Background()

		//connect that user to the feed
		posts, err := s.db.GetPost(ctx, int32(limit32))
		if err != nil{
		return err
		}
		for _, post := range posts{
			fmt.Println(post.Title)
		res := fmt.Sprintf("Post Title: %v \n Post Description %v \n Post Link %v \n", post.Title, post.Description.String, post.Url)
		fmt.Println(res)
		}
	}
	
	return nil	
}

	

