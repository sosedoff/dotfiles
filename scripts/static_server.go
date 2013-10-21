package main

import(
  "os"
  "fmt"
  "net/http"
  "github.com/jessevdk/go-flags"
)

type ServerOptions struct {
  Port string `short:"b" long:"bind" description:"Server port (5000)"`
  Path string `short:"p" long:"path" description:"Assets directory (current)"`
}

func main() {
  opts := ServerOptions {}

  _, err := flags.ParseArgs(&opts, os.Args)
  if err != nil {
    fmt.Println("Error", err)
    os.Exit(1)
  }

  if len(opts.Port) == 0 {
    opts.Port = "5000"
  }

  if len(opts.Path) == 0 {
    opts.Path, _ = os.Getwd()
  }

  bind := fmt.Sprintf(":%s", opts.Port)
  http.Handle("/", http.FileServer(http.Dir(opts.Path)))

  fmt.Printf("Starting server http://0.0.0.0:%s in %s\n", opts.Port, opts.Path)
  err = http.ListenAndServe(bind, nil)
  if err != nil {
    panic(err)
  }
}
