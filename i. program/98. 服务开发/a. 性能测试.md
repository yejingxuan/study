# 性能测试

- [性能测试](#性能测试)
  - [一、性能分析指标](#一性能分析指标)
  - [一、压力测试](#一压力测试)
    - [1、go-wrk压测工具](#1go-wrk压测工具)


## 一、性能分析指标

## 一、压力测试

### 1、go-wrk压测工具
压测工具安装：go get github.com/adeven/go-wrk
压测脚本：go-wrk -t=100 -c=100 -n=300 "http://127.0.0.1:9003/disaster/v1/gov/points?menuType=3"