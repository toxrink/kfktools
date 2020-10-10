package utils

import (
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-zookeeper/zk"
)

//TopicModel topic信息
type TopicModel struct {
	Title     string           `json:"title"`
	Partition []PartitionModel `json:"partition"`
}

//PartitionModel topic partition信息
type PartitionModel struct {
	Title string `json:"title"`
	State string `json:"state"`
}

//GetKafkaTopic 获取kafka主题列表
func GetKafkaTopic() []TopicModel {
	zkservers := strings.Split(beego.AppConfig.String("zkservers"), ",")
	conn, _, err := zk.Connect(zkservers, time.Second*5)
	defer conn.Close()
	tree := make([]TopicModel, 0)
	if nil != err {
		logs.Error(err)
		return tree
	}
	topicsPath := "/brokers/topics"
	topics, _, _ := conn.Children(topicsPath)
	for _, topic := range topics {
		partitionsPath := topicsPath + "/" + topic + "/partitions"
		partitions, _, _ := conn.Children(partitionsPath)
		tree2 := make([]PartitionModel, 0)
		for _, partition := range partitions {
			t2 := PartitionModel{}
			t2.Title = partition
			infoBytes, _, _ := conn.Get(partitionsPath + "/" + partition + "/state")
			t2.State = string(infoBytes)
			tree2 = append(tree2, t2)
		}

		t := TopicModel{}
		t.Title = topic
		t.Partition = tree2

		tree = append(tree, t)
	}
	return tree
}
