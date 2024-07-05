package main
import (
  "os/exec"
  "fmt"
  "log"
)

func main () {
  cmd, err := exec.Command("bash", "hello.sh").Output()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(cmd))
}
