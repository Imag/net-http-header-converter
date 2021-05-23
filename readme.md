# Quick Net/Http Formatter

# Download
https://github.com/Imag/net-http-header-converter/releases

# Usage
Create a `headers.txt` file in the build directory and place headers inside before running.
Run the build file `header-formatter.exe` in the same directory to create a formatted result

# Example
`headers.txt`
```
accept: /
accept-encoding: gzip, deflate, br
accept-language: en-US,en;q=0.9
content-length: 1189
content-type: multipart/form-data; boundary=----WebKitFormBoundary1zN9VIigLhQgkXMy
origin: https://google.com/
referer: https://google.com/
sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="90', "Google Chrome";v="90"
sec-ch-ua-mobile: ?0
sec-fetch-dest: empty
sec-fetch-mode: cors
sec-fetch-site: same-origin
user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36
x-requested-with: XMLHttpRequest\
```

`Result Output:`
```req.Header.Set("accept", "/")
req.Header.Set("accept-encoding", "gzip, deflate, br")
req.Header.Set("accept-language", "en-US,en;q=0.9")
req.Header.Set("content-length", "1189")
req.Header.Set("content-type", "multipart/form-data; boundary=----WebKitFormBoundary1zN9VIigLhQgkXMy")
req.Header.Set("origin", "https")
req.Header.Set("referer", "https")
req.Header.Set("sec-ch-ua", `"Not A;Brand";v="99", "Chromium";v="90', "Google Chrome";v="90"`)
req.Header.Set("sec-ch-ua-mobile", "?0")
req.Header.Set("sec-fetch-dest", "empty")
req.Header.Set("sec-fetch-mode", "cors")
req.Header.Set("sec-fetch-site", "same-origin")
req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36")
req.Header.Set("x-requested-with", "XMLHttpRequest\")
```
