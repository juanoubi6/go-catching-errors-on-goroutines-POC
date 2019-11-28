package main

import(
  "fmt"
  "errors"
  "golang.org/x/sync/errgroup"
  "time"
  "math/rand"
)

// Many goroutines are fired and each one generates a random number. If that random number
// is equal to the defined error number, an error message will be printed. On the other
// hand, if the error number is not generated, a success message will be printed.
func main() {
  var group errgroup.Group
  var minRange int = 0
  var maxRange int = 10
  var errorValue = 5

  for id := 1; id<10; id++{
    group.Go(func() error{
      rand.Seed(time.Now().UnixNano())
      randomValue := rand.Intn(maxRange - minRange) + minRange
      fmt.Println("The random value obtained by this gouroutine was:", randomValue)
      if randomValue == errorValue{
        return errors.New("Got the error value!!")
      }else{
        return nil
      }
    })
  }

  if err := group.Wait(); err == nil {
    fmt.Println("The error value was not generated, finished successfully")
  }else{
    fmt.Println("An error happened: ", err.Error())
  }
  
}
