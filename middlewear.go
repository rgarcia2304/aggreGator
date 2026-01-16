package main 
import(
	"errors"
	"context"
	"database/sql"
	"github.com/rgarcia2304/aggreGator/internal/database"

)
func middlewearLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error{

	return func(s *state, cmd command) error{
		queriedName := sql.NullString{String: s.cfg.Username, Valid: true}
		ctx := context.Background()

		//check if the name exists in the database
		usr, err := s.db.GetUser(ctx, queriedName)
		if err != nil{
			return errors.New("Issue fetching user")
		}
		return handler(s, cmd, usr)
	}
}
