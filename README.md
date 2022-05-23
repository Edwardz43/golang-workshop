# Homeworks from Golang workshop

## Homework 1
* 建立新專案 (go mod init)
  * 建立 .env 檔案
  * 撰寫 main.go 讀取 .env 檔案內容
  * 寫簡單測試 (main_test.go)

## Homework 2
* fun Add(i, j int) (int, error)
  * i, j 不能小於 0
* fun Sum(i int) (int, error)
  * i 不能小於 0
  * 0 + 1 + 2 + .... + I
* func main
  * Add (100, -1)
  * Add (100, 100)
  * Sum (100)
  * Sum (-100)

## Homework 3
```go
package main

func main() {
    // Create
    var stack []string

    // Push
    stack = append(stack, "world!")
    stack = append(stack, "Hello ")

    // Output: Hello world!
}
```
## Homework 4
* 撰寫 Email Package
  * 支援 SMTP
  * 支援 AWS SES
* 撰寫 Storage Package
  * 支援 Local Storage
  * 支援 AWS S3


## Homework 5
* 建立一個 go routine,並且執行 30 次
  * 透過 sync.WaitGroup 等待完成執行
  * 每一個 Job 執行時間不同
    * 初始印出 #job_id start job
    * 結束印出 #job_id finish job
## Homework 6
* 寫一個 Sum of Squares
  * 輸入 5
  * 1 * 1 + 2 * 2 + 3 * 3 + 4 * 4 + 5 * 5 = 55
  * 使用一個 int channel

## Homework 7
* 執行 100 次 job 丟到背景處理
  * 結果輸出到 outChan
  * 錯誤輸出到 errChan
  * 確保 Job 正常執行完畢
> 需要 3 個 Channel: outChan, errChan, finishChan