package cmd

import (
    "errors"
    "log"
    // "strconv"
    // "time"

    "github.com/spf13/cobra"

    "github.com/growlog/accounts-server/internal/utils"
)

func init() {
    rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
    Use:   "verify [FIELDS]",
    Short: "Verifies access token",
    Long:  `Command used to check access token and see if its valid and if not then print the reason why.`,
    Args: func(cmd *cobra.Command, args []string) error {
        if len(args) < 1 {
          return errors.New("requires the following fields: accessToken")
        }
        return nil
    },
    Run: func(cmd *cobra.Command, args []string) {
        // Get our user arguments.
        accessToken := args[0]

        claims, err := utils.VerifyAccessTokenString(accessToken)
        log.Printf("------------ CLAIMS ------------")
        log.Printf("%v\n", claims)
        log.Printf("------------ ERRORS ------------")
        log.Printf("%v\n", err)
        log.Printf("--------------------------------")
    },
}
