package main

import (
    "fmt"
    "time"

   vegeta "github.com/lucas-clemente/quic-go/lib"
  //vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
    // 1. 压测时长&速率
    rate := vegeta.Rate{Freq: 500, Per: time.Second}
    duration := 60 * time.Second

    // 2. 压测接口
    targeter := vegeta.NewStaticTargeter(vegeta.Target{
        Method: "GET",
        URL:    "https://localhost:6121/html" ,
    })

    // 3. 启动压测并收集结果
    var metrics vegeta.Metrics
    attacker := vegeta.NewAttacker()
    for res := range attacker.Attack(targeter, rate, duration, "Qi Lin!") {
        metrics.Add(res)

      // fmt.Printf("%s\n", res.Body)
       // fmt.Printf("%d\n", res.BytesOut)
        //fmt.Printf("%d\n", res.BytesIn)
        //fmt.Printf("%s\n", res.Attack)
        //fmt.Printf("%d\n", res.Code)
       // fmt.Printf("%s\n", res.Error)
        //fmt.Printf("%s\n", res.Latency)
        //fmt.Printf("%s\n", res.Timestamp)
        //fmt.Printf("%d\n", res.Seq)
        //fmt.Printf("============================================\n")




    }

    // 4. 处理结果
    metrics.Close()

    // 5. 打印感兴趣的指标
    //fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
    fmt.Printf("每秒攻击次数: %f\n", metrics.Rate)
    fmt.Printf("总攻击数: %d\n", metrics.Requests)
    fmt.Printf("最大延迟: %s\n", metrics.Latencies.Max)
    fmt.Printf("平均延迟: %s\n", metrics.Latencies.Mean)
    fmt.Printf("延迟总和: %s\n", metrics.Latencies.Total)
    fmt.Printf("成功率: %f\n", metrics.Success)
    fmt.Printf("每秒吞吐量: %f\n", metrics.Throughput)
    fmt.Printf("报表: %s\n", metrics.Histogram)
    fmt.Printf("请求错误后返回的数据: %s\n", metrics.Errors)
    /*fmt.Printf("99th percentile: %s\n", metrics.End)
    fmt.Printf("99th percentile: %s\n", metrics.BytesIn)
    fmt.Printf("99th percentile: %s\n", metrics.BytesOut)
    fmt.Printf("99th percentile: %s\n", metrics.Duration)
    fmt.Printf("99th percentile: %s\n", metrics.Earliest)
    fmt.Printf("99th percentile: %s\n", metrics.Histogram)
    fmt.Printf("99th percentile: %s\n", metrics.Rate)
    fmt.Printf("99th percentile: %s\n", metrics.Requests)
    fmt.Printf("99th percentile: %s\n", metrics.Latest)
    fmt.Printf("99th percentile: %s\n", metrics.StatusCodes)
    fmt.Printf("99th percentile: %s\n", metrics.Success)
    fmt.Printf("99th percentile: %s\n", metrics.Errors)*/
}

