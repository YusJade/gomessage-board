package main

import (
	"context"

	"github.com/YusJade/gomessage-board/config"
	"github.com/YusJade/gomessage-board/message/ports"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func init() {
	err := config.NewViperConfig()
	if err != nil {
		logrus.Fatal("failed to load config", err)
	}
}

func main() {
	addr := viper.GetString("http-addr")
	uri := viper.GetString("mongodb-url")
	logrus.SetLevel(logrus.DebugLevel)
	// logrus.Debug(uri)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Fatal("failed to connect with mongodb:", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logrus.Fatal("failed to connect with mongodb:", err)
	} else {
		logrus.Info("success to connect with mongodb.")
	}

	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			logrus.Fatal("failed to disconnect with mongodb:", err)
		}
	}()

	database := client.Database("message-board")
	collection := database.Collection("message")
	apiRouter := gin.New()
	ports.RegisterHandlersWithOptions(apiRouter, HTTPServer{collection: *collection}, ports.GinServerOptions{
		BaseURL: "/api",
	})
	apiRouter.Group("/api")
	if err := apiRouter.Run(addr); err != nil {
		logrus.Fatal("failed to run http server.", err)
	}
}
