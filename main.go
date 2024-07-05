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
  
  cmd2, err2 := exec.Command("python", "hello.py").Output()
    if err2 != nil {
      log.Fatal(err2)
    }
  fmt.Println(string(cmd2))
}
