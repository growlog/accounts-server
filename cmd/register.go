package cmd

import (
    "errors"
    "fmt"
    "os"
    "strconv"

    "github.com/spf13/cobra"

    "github.com/growlog/accounts-server/internal/models"
)

func init() {
    rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
    Use:   "register [FIELDS]",
    Short: "Registers a user account.",
    Long:  `Command used to create a GrowLog account, you begin with a single sign-in identity that has complete access to all GrowLog services and resources in the account. GrowLog accounts are accessed by signing in with an email address and password.`,
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 7 {
          return errors.New("requires the following fields: email, first name, last name, password, tenant id, tenant schema, group id")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        // Get our user arguments.
        email := args[0]
        firstName := args[1]
        lastName := args[2]
        password := args[3]
        tenantIdString := args[4]
        tenantSchema := args[5]
        roleIdString := args[6]

        // Minor modifications.
        tenantId, _ := strconv.ParseInt(tenantIdString, 10, 64)
        roleId, _ := strconv.ParseInt(roleIdString, 10, 64)

        // Load up our `environment variables` from our operating system.
        dbHost := os.Getenv("GROWLOG_ACCOUNT_DB_HOST")
        dbPort := os.Getenv("GROWLOG_ACCOUNT_DB_PORT")
        dbUser := os.Getenv("GROWLOG_ACCOUNT_DB_USER")
        dbPassword := os.Getenv("GROWLOG_ACCOUNT_DB_PASSWORD")
        dbName := os.Getenv("GROWLOG_ACCOUNT_DB_NAME")

        // Initialize and connect our database layer for the command.
        dal := models.InitDataAccessLayer(dbHost, dbPort, dbUser, dbPassword, dbName)

        user, err := dal.CreateUser(email, firstName, lastName, password, tenantId, tenantSchema, roleId)
        if err != nil {
            fmt.Println("Failed registering user!")
            fmt.Println(err)
        } else {
            fmt.Println("User registered with ID #", user.Id)
        }
    },
}
