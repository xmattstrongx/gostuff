package main
// piece of code that has several components that run in parallel off to the side
// they admit a heartbeat so that the monitoring code can assure they arent deadlocked or failed.

// write code that monitors the heartbeat so that it occurs within a piece of time for each component


// Q: We have a situation were we want to monitor the health of a set of components. The health of the components is measured in a 30 second heart beat event. If any component fails to satisfy the heartbeat the monitor func should return. How can we acomplish this?

// Function to monitor components health
func monitor(heartBeatPerSeconds int64, c chan int) {
    var heartBeatDuration time.Duration = heartBeatPerSeconds * time.Second
    
    hbs := make(map[string]time.Time)
    
    for {
    select {
    case <- time.After(heartBeatDuration):
        for i, v := range hbs {
            if v.Duration > 30 duration {
                return
            }
        }        
    case id := <- c:
        hbs[id] = time.Now()
    }
}

func myHeartBeatFn(c *Component) {
    return true
}

// Public component
type Component struct {
    ID string
    HeartBeatFn func(c *Component) {}
}

func Setup() {
   c := make(chan int,10)
   fn := func(comp *Component) {
      c <- 0
   }
   
   RunComponents(fn)
   
   monitor(30, c)
}

func RunComponents(fn func(*Component)) {
   for i:=0; i<10; i++ {
      c := Component{
          ID: strconv.Itoa(i),
          HeartBeatFn: fn,
      }
      go c.Run()
   }
}

func (c *Component) Run() {
   for {
      //...
      c.HeartBeatFn(c)
   }
}


