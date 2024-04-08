package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
// Необходимо написать функцию для параллельного выполнения заданий в n параллельных горутинах:
// n - количество запускаемх горутин
// m - количество ошибок по достижению m нужно остановить работу

1.  количество создаваемых горутин не должно зависеть от числа заданий, т.е.
функция должна запускать n горутин для конкурентной обработки заданий
и, возможно, еще несколько вспомогательных горутин;

2. функция должна останавливать свою работу, если произошло m ошибок;
после завершения работы функции (успешного или из-за превышения m)
не должно оставаться работающих горутин.

*/

type Task func() error

func Run(tasks []Task, n, m int) error {
	var errorCount int32
	var wg sync.WaitGroup

	tasksChan := make(chan Task)
	doneChan := make(chan struct{})

	go func() {
		for i := 0; i < len(tasks); i++ {
			fmt.Printf("task %d is %#v\n", i, tasks[i])
			tasksChan <- tasks[i]
		}
		close(tasksChan)
	}()

	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, tasksChan, &wg, &errorCount, doneChan)
	}
	wg.Wait()

	fmt.Printf("Общее число ошибок %d\n", errorCount)
	// если произошло m ошибок то выполняем закрытие всех горутин
	if errorCount == int32(m) {
		close(doneChan)
		return fmt.Errorf("ErrErrorsLimitExceeded. Limit is %d", errorCount)
	}
	return nil
}

func worker(id int, jobs <-chan Task, wg *sync.WaitGroup, ec *int32, done chan struct{}) {
	defer wg.Done()
	for job := range jobs {
		select {
		case <-done:
		default:
			err := job() // Здесь я запускаю задачу
			fmt.Printf("Worker %d выполнил задачу %v => ", id, job)
			if err != nil {
				fmt.Printf("err in job: [%v]\n", job)
				atomic.AddInt32(ec, 1)
			}

		}
	}
}

func main() {
	n := 2
	//tasksCount := 5
	//tasks := make([]Task, 0)

	// for i := 0; i < tasksCount; i++ {
	// 	taskSleep := time.Millisecond * time.Duration(rand.Intn(100))

	// 	tasks = append(tasks, func() error {
	// 		//fmt.Printf("[%d] TaskSleepTime %d\n", i, taskSleep)
	// 		time.Sleep(taskSleep)
	// 		return nil
	// 	})
	// }
	tasks := []Task{
		func() error {
			fmt.Printf("sum: %d\n", 2+2)
			//return nil
			return fmt.Errorf("err in sum")
		},
		func() error {
			fmt.Printf("mul: %d\n", 2*5)
			return nil
		},
		func() error {
			fmt.Printf("sub: %d\n", 10-2)
			//return nil
			return fmt.Errorf("err in sub")
		},
		// добавляем здесь любые другие задачи
	}

	Run(tasks, n, 2)

}
