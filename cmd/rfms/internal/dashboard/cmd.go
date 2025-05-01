package dashboard

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/way-platform/rfms-go/cmd/rfms/v4/internal/dashboard/server"
)

// NewCommand returns a new [cobra.Command] for the rFMS dashboard.
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Serve a local rFMS dashboard",
	}
	port := cmd.Flags().Int("port", 8080, "Port to serve the dashboard on")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context(), *port)
	}
	return cmd
}

func serve(ctx context.Context, port int) error {
	server, err := server.NewServer()
	if err != nil {
		return err
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: server,
	}
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
