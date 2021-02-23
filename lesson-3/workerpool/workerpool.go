package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Job struct {
	req     *http.Request
	payload []byte
}

type Worker struct {
	client  *http.Client
	wg      *sync.WaitGroup
	num     int
	jobChan <-chan *Job
}

func main() {
	var numThreads, numJobs, workTime int
	var method, data, contentType string
	flag.IntVar(&numThreads, "t", 1, "threads quantity")
	flag.IntVar(&numJobs, "j", 1, "jobs quantity")
	flag.IntVar(&workTime, "s", 1, "working time (sec)")
	flag.StringVar(&method, "m", "GET", "method")
	flag.StringVar(&data, "d", "", "request body")
	flag.StringVar(&contentType, "ct", "application/json", "test")

	flag.Parse()

	switch method {
	case "GET", "DELETE", "OPTIONS", "POST", "UPDATE", "PATCH", "PUT":
	default:
		log.Fatal("non-existent method")
	}

	isNumJobsPassed := numJobs != 1
	if numJobs > 1 && workTime > 1 {
		log.Fatalln("just one flag:'t' or 'n'")
	}

	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	client := &http.Client{
		Timeout: time.Minute,
	}
	for i := 0; i < 5; i++ {
		worker := NewWorker(i+1, wg, jobChan, client)
		wg.Add(1)
		go worker.Handle()
	}
	start := time.Now()
	for i := 0; i < numJobs; i++ {
		if !isNumJobsPassed && time.Since(start) <= time.Duration(workTime)*time.Second {
			numJobs++
		}
		req, err := http.NewRequest(method, "http://ya.ru", strings.NewReader(data))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Content-Type", contentType)

		jobChan <- &Job{
			req:     req,
			payload: []byte(fmt.Sprintf("mes from start cycle %d", i)),
		}
	}

	close(jobChan)
	wg.Wait()
	fmt.Println(float64(numJobs) / time.Since(start).Seconds())
}

func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {
		//
		time.Sleep(1 * time.Second)
		respWithClient, err := w.client.Do(job.req)
		respWithGet, err := http.Get("http://ya.ru")
		if err != nil {

		}

		log.Printf("Status: %d \n worker %d processing job with payload %s ###RESPFROMCLIENT###", respWithClient.StatusCode, w.num, string(job.payload))
		log.Printf("Status: %d \n worker %d processing job with payload %s ###RESPFROMGET###", respWithGet.StatusCode, w.num, string(job.payload))
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job, client *http.Client) *Worker {
	return &Worker{
		wg:      wg,
		client:  client,
		num:     num,
		jobChan: jobChan,
	}
}
