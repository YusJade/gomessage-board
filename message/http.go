package main

import (
	"net/http"
	"time"

	"github.com/YusJade/gomessage-board/message/domain"
	"github.com/YusJade/gomessage-board/message/ports"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HTTPServer struct {
	collection mongo.Collection
}

// (GET /message-board)
func (s HTTPServer) GetMessageBoard(c *gin.Context) {
	logrus.Info("execute HTTPServer.GetMessageBoard")
	defer func() {
		logrus.Info("finished HTTPServer.GetMessageBoard")
	}()

	cur, err := s.collection.Find(c, bson.D{}, options.Find())
	if err != nil {
		logrus.Error("mongo.Collection.Find failed:", err)
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	res := []domain.Message{}
	cur.All(c, &res)
	c.JSON(http.StatusOK, res)
}

// (POST /message-board)
func (s HTTPServer) PostMessageBoard(c *gin.Context) {
	logrus.Info("execute HTTPServer.PostMessageBoard")
	defer func() {
		logrus.Info("finished HTTPServer.PostMessageBoard")
	}()
	var req ports.Message
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("failed to bind json to object: %v | %v", err.Error(), req)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	writeModel := &domain.Message{
		ID:      primitive.NewObjectID(),
		Content: *req.Content,

		Datetime: time.Now().UTC().Format(time.RFC3339),
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("prepare to insert write model:", writeModel)
	res, err := s.collection.InsertOne(c, writeModel)
	// res, err := s.collection.InsertOne(c, bson.M{
	// 	"content":  writeModel.Content,
	// 	"datetime": writeModel.Datetime,
	// })
	if err != nil {
		logrus.Error("failed to insert into collection:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	logrus.Info("success to insert into collection:", res)
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
