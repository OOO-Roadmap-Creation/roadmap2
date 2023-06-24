package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/OOO-Roadmap-Creation/roadmap2"
	"github.com/OOO-Roadmap-Creation/roadmap2/pkg/repository"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"math/rand"
)

func Listen(repo *repository.Repository) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:29092",
		"group.id":           "myGroup",
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": "false",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic"}, nil)
	defer c.Close()
	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message RECEIVED on %s: %s : %s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			if rand.Float64() > 0.5 {
				fmt.Printf("Panic! Message NOT COMMIT on %s: %s : %s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
				panic("Error when handle message")
			}
			process(msg.Key, msg.Value, repo)
			fmt.Printf("Message PROCESSED on %s: %s : %s\n", msg.TopicPartition, string(msg.Key), string(msg.Value))
			_, err := c.CommitMessage(msg)
			if err != nil {
				return
			}
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

}

func process(key, message []byte, repo *repository.Repository) {
	//dataJson :=
	//{
	//	"eventType":"OBJECT_CREATED", "targetObjectType":"PRE_PUBLISHED_ROADMAP", "targetObjectId":{
	//	"prePublishedRoadmapId":6, "dateOfChange":"2023-06-17T16:17:54.394828"
	//}, "changedAttributes":{
	//	"baseRoadmapId":5, "prePublishedRoadmapId":6, "title":"eiusmod voluptate", "description":"et nulla reprehenderit exercitation ullamc", "dateOfChange":"2023-06-17T16:17:54.394828", "authorId":52, "nodes":[{"id":19, "title":"string title", "description":"string description 3", "priority":2, "parentId":null},{"id":22, "title":"string title", "description":"string description 3", "priority":2, "parentId":null},{"id":20, "title":"string title", "description":"string description 3", "priority":2, "parentId":19},{"id":21, "title":"string title", "description":"string description 3", "priority":2, "parentId":20},{"id":23, "title":"consequat eu aute 23", "description":"reprehenderit 23", "priority":10, "parentId":22},{"id":24, "title":"consequat eu aute", "description":"reprehenderit", "priority":10, "parentId":23}]}}

	var roadmap roadmap
	err := json.Unmarshal([]byte(message), &roadmap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Event Type:", roadmap.EventType)
	fmt.Println("Target Object Type:", roadmap.TargetObjectType)
	fmt.Println("Target Object ID:", roadmap.TargetObjectID)
	fmt.Println("Changed Attributes:", roadmap.ChangedAttributes)
	new := roadmap2.PublishedRoadmap{
		Id:            roadmap.TargetObjectID.PrePublishedRoadmapID,
		Version:       1,
		Visible:       true,
		Title:         roadmap.ChangedAttributes.Title,
		Description:   roadmap.ChangedAttributes.Description,
		DateOfPublish: roadmap.ChangedAttributes.DateOfChange,
	}
	old, err := repo.PR.GetById(roadmap.TargetObjectID.PrePublishedRoadmapID)
	if err != nil {
		fmt.Println(err)
		repo.PR.Create(new)
		return
	}
	fmt.Println(old)
	new.Version = old.Version
	if new != old {
		new.Version++
		repo.PR.Create(new)
	}

}

func Recoverer(maxPanics, id int, f func(repo *repository.Repository), repo *repository.Repository) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HERE", id)
			fmt.Println(err)
			if maxPanics == 0 {
				panic("TOO MANY PANICS")
			} else {
				go Recoverer(maxPanics-1, id, f, repo)
			}
		}
	}()
	f(repo)
}
