package main

import (
	"fmt"
	"time"
)

type Observer interface {
	Update(string)
}

type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers(string)
}

type ExternalSystem struct {
	name      string
	observers []Observer
}

func (es *ExternalSystem) RegisterObserver(o Observer) {
	es.observers = append(es.observers, o)
}

func (es *ExternalSystem) RemoveObserver(o Observer) {
	for i, observer := range es.observers {
		if observer == o {
			es.observers = append(es.observers[:i], es.observers[i+1:]...)
			break
		}
	}
}

func (es *ExternalSystem) NotifyObservers(status string) {
	fmt.Printf("External system %s status changed to %s\n", es.name, status)
	for _, observer := range es.observers {
		observer.Update(status)
	}
}

type Prometheus struct {
	ExternalSystem
}

func (p *Prometheus) Update(status string) {
	fmt.Printf("Prometheus received status update from %s: %s\n", p.name, status)
}

type Elasticsearch struct {
	ExternalSystem
}

func (e *Elasticsearch) Update(status string) {
	fmt.Printf("Elasticsearch received status update from %s: %s\n", e.name, status)
}

func main() {
	prometheus := &Prometheus{
		ExternalSystem: ExternalSystem{
			name: "Prometheus",
		},
	}
	elasticsearch := &Elasticsearch{
		ExternalSystem: ExternalSystem{
			name: "Elasticsearch",
		},
	}

	// 注册观察者
	prometheus.RegisterObserver(elasticsearch)
	elasticsearch.RegisterObserver(prometheus)

	// 模拟外部系统状态变化
	go func() {
		for {
			prometheus.NotifyObservers("UP")
			time.Sleep(5 * time.Second)
			elasticsearch.NotifyObservers("DOWN")
			time.Sleep(5 * time.Second)
			prometheus.NotifyObservers("DOWN")
			time.Sleep(5 * time.Second)
			elasticsearch.NotifyObservers("UP")
			time.Sleep(5 * time.Second)
		}
	}()

	// 阻塞主线程，保持程序运行
	select {}
}
