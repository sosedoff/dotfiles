package main

import(
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func getUrlContents(url string) string {
  var err error
  var resp *http.Response
  var body []byte

  resp, err = http.Get(url)
  if err != nil {
    return ""
  }

  defer resp.Body.Close()

  body, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    return ""
  }

  return string(body)
}

func parseJson(body string) {
  var versions []interface{}

  err := json.Unmarshal([]byte(body), &versions)
  if err != nil {
    fmt.Println("Failed to parse JSON")
    os.Exit(1)
  }

  if len(versions) == 0 {
    fmt.Println("Gem not found")
    os.Exit(1)
  }

  fmt.Println(versions[0])
}

func main() {
  if len(os.Args) == 1 {
    fmt.Println("Name required")
    os.Exit(1)
  }

  name   := os.Args[1]
  url    := fmt.Sprintf("http://dependenci.com/api/rubygems/%s", name)
  result := getUrlContents(url)

  if len(result) == 0 {
    fmt.Println("Unable to fetch data")
    os.Exit(1)
  }

  parseJson(result)
}
