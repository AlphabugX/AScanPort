# AScanPort
AScanPort 高速、多线程、全端口、单IP扫描。

## 食用方法

```
  -d string
    	text|json (default "text")
  -o string
    	result.txt (default "result_20220415_172833.txt")
  -t string
    	127.0.0.1 or f5.ink (default "f5.ink")
```

### 支持导出格式:
- text: f5.ink:80
- json: {"f5.ink":"80"}

### 扫描结果

![image](https://user-images.githubusercontent.com/27001865/163555109-bbff021c-c371-4dc3-a493-7bf7188b5043.png)

## AScanPort VS masscan

VPS:腾讯云、200Mbps、上海

![image](masscan_VS_AScanPort.jpg)