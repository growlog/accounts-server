package cmd

import (
    "errors"
    "log"
    "strconv"
    "time"

    "github.com/spf13/cobra"

    "github.com/growlog/accounts-server/internal/utils"
)

func init() {
    rootCmd.AddCommand(authorizeCmd)
}

var authorizeCmd = &cobra.Command{
    Use:   "authorize [FIELDS]",
    Short: "Generate custom credentials.",
    Long:  `Command used to grant access to our system by the inputted custom values.`,
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 3 {
          return errors.New("requires the following fields: userId, thingId, duration")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        // Get our user arguments.
        userIdString := args[0]
        thingIdString := args[1]
        durationString := args[2]

        // Minor modifications.
        userId, _ := strconv.ParseInt(userIdString, 10, 64)
        thingId, _ := strconv.ParseInt(thingIdString, 10, 64)
        duration, _ := strconv.ParseInt(durationString, 10, 64)

        // Generate our access token.
        tokenString, err := utils.GenerateAccessToken(userId, thingId, time.Duration(duration))

        // If there is an error in creating the JWT return an internal server error
        if err != nil {
            log.Fatalf("could not generate JWT token because: %v",err)
        }
        log.Printf("ACCESS TOKEN: %v\n\n", tokenString)
    },
}
