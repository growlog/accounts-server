package cmd

import (
    "errors"
    "log"
    "os"
    // "strconv"

    "github.com/spf13/cobra"

    "github.com/growlog/accounts-server/internal/models"
    "github.com/growlog/accounts-server/internal/utils"
)

func init() {
    rootCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
    Use:   "login [FIELDS]",
    Short: "Generate credentials for valid email and password.",
    Long:  `Command used to grant access to our system for the valid email and password pair.`,
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 2 {
          return errors.New("requires the following fields: email, password")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        // Get our user arguments.
        email := args[0]
        password := args[1]

        // Load up our `environment variables` from our operating system.
        dbHost := os.Getenv("GROWLOG_ACCOUNT_DB_HOST")
        dbPort := os.Getenv("GROWLOG_ACCOUNT_DB_PORT")
        dbUser := os.Getenv("GROWLOG_ACCOUNT_DB_USER")
        dbPassword := os.Getenv("GROWLOG_ACCOUNT_DB_PASSWORD")
        dbName := os.Getenv("GROWLOG_ACCOUNT_DB_NAME")

        // Initialize and connect our database layer for the command.
        dal := models.InitDataAccessLayer(dbHost, dbPort, dbUser, dbPassword, dbName)

        user, err := dal.FindUserByEmail(email)
        if err != nil {
            log.Fatalf("could not find user, found error: %v", err)
        }
        if user == nil {
            log.Fatalf("could not find user with email: %v", email)
        }

        // Verify our inputted password is valid.
        var isCorrectPassword bool = utils.CheckPasswordHash(password, user.PasswordHash.String)
        if isCorrectPassword == false {
            log.Fatalf("could not login into account with email and password: %v %v", email, password)
        }

        // Generate our access token.
        tokenString, err := utils.GenerateAccessToken(user.Id, 0, 15)

        // If there is an error in creating the JWT return an internal server error
        if err != nil {
            log.Fatalf("could not generate JWT token because: %v",err)
        }
        log.Printf("ACCESS TOKEN: %v\n\n", tokenString)

        // isValidJWTToken, validationError := utils.VerifyAccessTokenString(tokenString)
        // log.Printf("\n%v %v\n\n", isValidJWTToken, validationError)
    },
}
