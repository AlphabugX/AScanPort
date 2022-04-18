# AScanPort
AScanPort 高速、多线程、全端口、单IP扫描。

## 食用方法

```
Usage of AScanPort:
  -check int
        MaxCheck:Connect check the maximum number (default 1)
  -d string
        Result format: text=>ip:port,json=>{"ip":"port"} (default "text")
  -h string
        127.0.0.1 or f5.ink (default "f5.ink")
  -o string
        result.txt (default "result_20220418_135557.txt")
  -t int
        Maximum threads (default 14000)
  -time int
        timeout:3 seconds (default 3)
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