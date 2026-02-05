# 🛡️ SentryHTTP-Package

A modern Python package to check HTTP/HTTPS endpoints with status codes, response times, redirects, and TLS monitoring.  
-coded with ❤️ by **Only Human Are Learn**

---

## 🚀 Features

- Check HTTP/HTTPS endpoint status codes (200, 404, 500, etc.)
- Measure response times for URLs
- Detect redirects and final destination URLs
- Validate SSL/TLS certificates
- CLI-friendly and scriptable
- Lightweight and fast, built in Python

---

## 📦 Installation

Install via pip:

```bash
pip install sentryhttp-package
```

Or clone the repo:
```bash
git clone https://github.com/username/sentryhttp-package.git
cd sentryhttp-package
pip install .
```
---

## 💻 Usage

### 🐍 Python Example
```python
from sentryhttp import SentryHTTP

checker = SentryHTTP()
result = checker.check("https://example.com")
print(result)
```
### CLI Example
```bash
sentryhttp https://example.com
sentryhttp --file raw_urls.txt
```
---

## ⚡ Options

```bash
-h, --help       Show help message
-u, --url        Check a single URL
-f, --file       Check multiple URLs from a file
-t, --timeout    Set request timeout (default: 5s)
--no-ssl         Skip SSL/TLS validation
```

---
## 📈 Output Example

```yaml
URL: https://example.com
Status: 200 OK
Response Time: 0.231s
SSL/TLS: Valid (expires in 45 days)
Redirects: None
```
---

## 📝 Contributing
Contributions are welcome! Feel free to submit issues or pull requests.

## 🔗 License
This project is licensed under the MIT License.
