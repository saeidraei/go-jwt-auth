package main

import (
	"fmt"
	migrate "github.com/saeidraei/go-jwt-auth/db"

	"github.com/saeidraei/go-jwt-auth/implem/gin.server"
	"github.com/saeidraei/go-jwt-auth/implem/jwt.authHandler"
	"github.com/saeidraei/go-jwt-auth/implem/logrus.logger"
	mongoUserRW "github.com/saeidraei/go-jwt-auth/implem/mongo.userRW"
	mysqlUserRW "github.com/saeidraei/go-jwt-auth/implem/mysql.userRW"
	"github.com/saeidraei/go-jwt-auth/implem/user.validator"
	"github.com/saeidraei/go-jwt-auth/infra"
	"github.com/saeidraei/go-jwt-auth/uc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build number and versions injected at compile time, set yours
var (
	Version = "unknown"
	Build   = "unknown"
)

// the command to run the server
var rootCmd = &cobra.Command{
	Use:   "go-realworld-clean",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show build and version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Build: %s\nVersion: %s\n", Build, Version)
	},
}
var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run the migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate.RunMigration()
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(migrationCmd)
	cobra.OnInitialize(infra.CobraInitialization)

	infra.LoggerConfig(rootCmd)
	infra.ServerConfig(rootCmd)
	infra.DatabaseConfig(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Fatal()
	}
}

func run() {
	ginServer := infra.NewServer(
		viper.GetInt("server.port"),
		infra.DebugMode,
	)

	authHandler := jwt.New(viper.GetString("jwt.Salt"))
	routerLogger := logger.NewLogger("TEST",
		viper.GetString("log.level"),
		viper.GetString("log.format"),
	)
	var userRW uc.UserRW
	if viper.GetString("server.userDriver") == "mysql" {
		userRW = mysqlUserRW.New()
	} else {
		userRW = mongoUserRW.New()

	}
	server.NewRouterWithLogger(
		uc.HandlerConstructor{
			Logger:        routerLogger,
			UserRW:        userRW,
			UserValidator: validator.New(),
			AuthHandler:   authHandler,
		}.New(),
		authHandler,
		routerLogger,
	).SetRoutes(ginServer.Router)

	ginServer.Start()
}
