package main

import (
	"fmt"
	"github.com/alitto/pond"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	pool := pond.New(10, 100)

	// Register pool metrics collectors

	// Worker pool metrics
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pool_workers_running",
			Help: "Number of running worker goroutines",
		},
		func() float64 {
			return float64(pool.RunningWorkers())
		}))
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pool_workers_idle",
			Help: "Number of idle worker goroutines",
		},
		func() float64 {
			return float64(pool.IdleWorkers())
		}))

	// Task metrics
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "pool_tasks_submitted_total",
			Help: "Number of tasks submitted",
		},
		func() float64 {
			return float64(pool.SubmittedTasks())
		}))
	prometheus.MustRegister(prometheus.NewGaugeFunc(
		prometheus.GaugeOpts{
			Name: "pool_tasks_waiting_total",
			Help: "Number of tasks waiting in the queue",
		},
		func() float64 {
			return float64(pool.WaitingTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "pool_tasks_successful_total",
			Help: "Number of tasks that completed successfully",
		},
		func() float64 {
			return float64(pool.SuccessfulTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "pool_tasks_failed_total",
			Help: "Number of tasks that completed with panic",
		},
		func() float64 {
			return float64(pool.FailedTasks())
		}))
	prometheus.MustRegister(prometheus.NewCounterFunc(
		prometheus.CounterOpts{
			Name: "pool_tasks_completed_total",
			Help: "Number of tasks that completed either successfully or with panic",
		},
		func() float64 {
			return float64(pool.CompletedTasks())
		}))
	http.Handle("/metrics", promhttp.Handler())
	go submitTasks(pool)
	http.ListenAndServe(":8080", nil)
}

func submitTasks(pool *pond.WorkerPool) {
	for i := 0; i < 1000; i++ {
		n := i
		pool.Submit(func() {
			fmt.Printf("Running task #%d\n", n)
			time.Sleep(500 * time.Millisecond)
		})
	}
}
