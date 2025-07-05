package service

import (
	"cudo/test/dao"
	"cudo/test/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func GetFraudTransaction(gin *gin.Context) {
	resultChan := make(chan interface{}, 3)
	chDuratiunTask1 := make(chan int64, 1)
	chDuratiunTask2 := make(chan int64, 1)
	chDuratiunTask3 := make(chan int64, 1)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		t := time.Now()
		wg.Done()
		data := freqCheck()
		t2 := time.Now()
		chDuratiunTask1 <- t2.Sub(t).Milliseconds()
		resultChan <- data
	}()

	go func() {
		t := time.Now()
		wg.Done()
		data := amountCheck()
		t2 := time.Now()
		chDuratiunTask2 <- t2.Sub(t).Milliseconds()
		resultChan <- data
	}()

	go func() {
		t := time.Now()
		wg.Done()
		data := patternCheck()
		t2 := time.Now()
		chDuratiunTask3 <- t2.Sub(t).Milliseconds()
		resultChan <- data
	}()

	wg.Wait()

	var results []interface{}

	for ch := range resultChan {
		results = append(results, ch)
	}


	transactionResponse := dto.CheckFraud{
		Transaction: results,
		ProcessMetadata: dto.ProcessMetadata{
			TotalTransactionAnalyzed: 0,
			DurationMs:               0,
			ParallelTasks: dto.ParallelTasks{
				FrequencyAnalysisDuration: ,
				AmountAnalysisDuration:    0,
				PatternAnalysisDuration:   0,
			},
		},
	}

	gin.JSON(http.StatusOK, transactionResponse)

}

func freqCheck() interface{} {
	var result interface{}
	transModels, err := dao.GetFreqCheck()
	if err != nil {
		return result
	}

	// TODO: add logic for check high freq

	return transModels

}

func amountCheck() interface{} {
	// TODO: add logic for check high freq

	return nil
}

func patternCheck() interface{} {
	// TODO: add logic for check high freq

	return nil
}

type TaskResult struct {
	Data      interface{}
	StartTime time.Time
	EndTime   time.Time
}
