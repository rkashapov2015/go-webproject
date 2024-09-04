package commands

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/rkashapov2015/webproject/internal/database/models"
	"github.com/rkashapov2015/webproject/internal/tools/security"
	"github.com/uptrace/bun"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func NewUserCommand(db *bun.DB) *cli.Command {
	return &cli.Command{
		Name:  "users",
		Usage: "manage users",
		Subcommands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create user",
				Action: func(c *cli.Context) error {
					var userName string
					var firstName string
					var lastName string
					var email string
					var typeRole string
					fmt.Println("Enter Username")
					_, err := fmt.Fscan(os.Stdin, &userName)
					if err != nil {
						return err
					}

					fmt.Println("Enter First Name")
					_, err = fmt.Fscan(os.Stdin, &firstName)
					if err != nil {
						return err
					}

					fmt.Println("Enter Last Name")
					_, err = fmt.Fscan(os.Stdin, &lastName)
					if err != nil {
						return err
					}

					fmt.Println("Enter email")
					_, err = fmt.Fscan(os.Stdin, &email)
					if err != nil {
						return err
					}

					fmt.Println("Enter type role")
					_, err = fmt.Fscan(os.Stdin, &typeRole)
					if err != nil {
						return err
					}

					fmt.Println("Enter password")
					bytePassword, err := term.ReadPassword(syscall.Stdin)
					if err != nil {
						return err
					}

					byteHash, err := security.GeneratePasswordHash(bytePassword)
					if err != nil {
						return err
					}

					fmt.Printf(
						"Username: %s \nFirst Name: %s \nLast Name: %s \nEmail: %s \nPassword: %s \nPasswordHash: %s \n",
						userName,
						firstName,
						lastName,
						email,
						string(bytePassword),
						string(byteHash),
					)

					values := map[string]interface{}{
						"username":      userName,
						"first_name":    firstName,
						"last_name":     lastName,
						"email":         email,
						"password_hash": string(byteHash),
						"active":        true,
						"created_at":    time.Now(),
						"updated_at":    time.Now(),
					}

					tx, err := db.BeginTx(c.Context, &sql.TxOptions{})
					if err != nil {
						return err
					}
					_, err = db.NewInsert().Model(&values).TableExpr("users").Exec(c.Context)
					if err != nil {
						return err
					}

					user := new(models.User)
					err = db.NewSelect().Model(user).Where("username = ?", userName).Scan(c.Context)
					if err != nil {
						tx.Rollback()
					}
					role := new(models.Role)
					err = db.NewSelect().Model(role).Where("type = ?", typeRole).Scan(c.Context)
					if err != nil {
						tx.Rollback()
					}
					userToRole := map[string]interface{}{
						"user_id": user.ID,
						"role_id": role.ID,
					}
					_, err = db.NewInsert().Model(&userToRole).TableExpr("users_to_roles").Exec(c.Context)
					if err != nil {
						tx.Rollback()
					}

					err = tx.Commit()
					if err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "check",
				Usage: "check user",
				Action: func(c *cli.Context) error {
					userName := strings.Join(c.Args().Slice(), "_")
					if userName == "" {
						return fmt.Errorf("userName is empty string")
					}
					fmt.Println("Enter password")
					bytePassword, err := term.ReadPassword(syscall.Stdin)
					if err != nil {
						return err
					}

					user := new(models.User)
					err = db.NewSelect().Model(user).Where("username = ?", userName).Scan(c.Context)
					if err != nil {
						return err
					}

					resultCheck := security.CheckPassword(string(bytePassword), user.PasswordHash)
					if resultCheck {
						fmt.Println("Password accepted")
					} else {
						fmt.Println("Wrong password")
					}

					return nil
				},
			},
		},
	}
}
