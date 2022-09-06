# Spider github repositories

---

This project is an API that gets data information about api repositories from github



The process of getting data and loading it in data.json file is repeated every 10 minutes, the process to complete takes 2s



#### Config

* In your code, in processing file:

```go
// change url and add your url of github api repos
const (
  url = "your url of api github repos"
)


