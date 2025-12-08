package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
	kafkaConsumer *kafka.Reader
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, //[]string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset, //FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua bán chứng khoán
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg
	return &s

	// return &StockInfo{
	// 	Message: msg,
	// 	Type:    typeMsg,
	// }
}

func actionStock(c *gin.Context) {
	s := newStock("Buy", "action")
	jsonBody, _ := json.Marshal(s)

	// create kafka message
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}

	// write message by kafka producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": "action Successfully",
	})
}

// consumer hóng để mua ATC
func RegisterConsumerATC(id int) {
	// group consumer??
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) //"consumer-group"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)
	defer reader.Close()

	fmt.Printf("Consumer (%d) Hong Phien ATC::", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer (%d) error: %v", id, err)
		}

		fmt.Printf("Consumer (%d), hong topic:%v, partition:%v, offset:%v, time:%d %s = %s\n",
			id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
	}
}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// đăng ký 2 consumer để hóng mua ATC (1 và 2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	go RegisterConsumerATC(3)
	go RegisterConsumerATC(4) // không lấy những tin nhắn cũ

	r.Run(":8999")
}
