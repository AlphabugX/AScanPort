# AScanPort
AScanPort 高速、多线程、全端口、单IP扫描。

## 食用方法

```
Usage of AScanPort:
  -check int
        MaxCheck:Connect check the maximum number (default 1)
  -format string
        Result format: text=>ip:port,json=>{"ip":"port"} (default "text")
  -h string
        Target:f5.ink|114.67.111.74|114.67.111.74/28|114.67.111.74-80|114.67.111.74-114.67.111.80|114.67.111.* (default "f5.ink")
  -out string
        result.txt
  -t int
        Maximum threads (default 10000)
  -time int
        timeout:3 seconds (default 2)
```

### 支持导出格式:
- text: f5.ink:80
- json: {"f5.ink":"80"}

### 扫描结果
```
./AScanPort_linux_amd64 -h f5.ink -time 1 -t 10000 -check 2
```
![image](AScanPort.jpg)

## AScanPort VS masscan

VPS:腾讯云、200Mbps、上海

![image](masscan_VS_AScanPort.jpg)

## 项目背景与更新主线

AScanPort是对标masscan的速度，暂未增加SYN模式

- 2022年4月20日 10:42:51关于增加fuzz、poc、爆破、获取特征的需求，AScanPort不再更新。后期另外成立一个项目，欢迎各路大佬多多指教。