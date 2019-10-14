package internal // github.com/growlog/accounts-server/internal

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/growlog/rpc/protos"

	"github.com/growlog/accounts-server/internal/controllers"
	"github.com/growlog/accounts-server/internal/models"
)

type AccountServer struct {
	webServerAddress string
	dal *models.DataAccessLayer
	grpcServer *grpc.Server
}

// Function will construct the Mikapod IAM application.
func InitAccountServer(dbHost, dbPort, dbUser, dbPassword, dbName, webServerAddress string) (*AccountServer) {

	// Initialize and connect our database layer for the entire application.
    dbInstance := models.InitDataAccessLayer(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Create our app's models if they haven't been created previously.
    dbInstance.CreateUserTable(false)

	// Create our application instance.
 	return &AccountServer{
		webServerAddress: webServerAddress,
		dal: dbInstance,
		grpcServer: nil,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the Mikapod IAM application.
func (app *AccountServer) RunMainRuntimeLoop() {
	// Open a TCP server to the specified localhost and environment variable
    // specified port number.
    lis, err := net.Listen("tcp", app.webServerAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Initialize our gRPC server using our TCP server.
    grpcServer := grpc.NewServer()

    // Save reference to our application state.
    app.grpcServer = grpcServer

    // For debugging purposes only.
    log.Printf("gRPC server is running.")

    // Block the main runtime loop for accepting and processing gRPC requests.
    pb.RegisterAccountServer(grpcServer, &controllers.AccountServer{
        // DEVELOPERS NOTE:
        // We want to attach to every gRPC call the following variables...
        DAL: app.dal,
    })
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *AccountServer) StopMainRuntimeLoop() {
	// Finish any RPC communication taking place at the moment before
    // shutting down the gRPC server.
    app.grpcServer.GracefulStop()

	// Shutdown our connection with our database.
	app.dal.Shutdown()
}
