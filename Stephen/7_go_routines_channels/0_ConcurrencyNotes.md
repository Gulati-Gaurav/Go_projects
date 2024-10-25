- Q - Go Scheduler: 
    - Go uses a scheduler that is responsible for the execution of goroutines. 
    - Goroutines need to yield control back to the scheduler at specific points like blocking calls, function calls, etc.
    - By default it works on one cpu core on our machine (even in case of multi core cpu). Can change settings for it to work in multi core. (see [](./SingleCore.jpg)  and  [](./MultiCore.jpg) image)
    - One CPU = concurrency  (context switching).
    - Multi CPU = parallelism

- Q - what are the conditions when the go routine yields the control back to scheduler
    1. Blocking System Calls:
    2. Function Calls:
        Function calls in Go often involve implicit checks for runnable goroutines.
        Before entering a new function, the scheduler might be invoked to see if any other goroutines are ready to run.
    3. Channel Operations:
        Certain channel operations, like unbuffered receives (<-ch) or sends to a full channel (ch <- value), can cause a goroutine to yield.
        The scheduler can then switch to another goroutine if the unbuffered receive cannot proceed immediately or the channel is full.
    4. Defer Statements:
        When a goroutine exits normally (returns from the function), any defer statements are executed in Last-In-First-Out (LIFO) order.
        Between each defer statement execution, the scheduler can potentially be invoked, allowing other goroutines to run.
    5. Garbage Collection:
        The Go runtime may pause goroutines (including the currently running one) to perform garbage collection This ensures efficient memory management without causing memory leaks.
    6. Fairness:
        If a goroutine runs for a very long time without yielding control at any of the points mentioned above, the runtime might preemptive it briefly to allow other goroutines a chance to run.

- Q - Goroutines vs. Threads: 
    Go Routines: Lightweight threads managed by the Go runtime. Go's goroutines are lightweight threads that share the underlying OS threads. This means multiple goroutines can be "running" concurrently, even on a single CPU.
    So doesn't relate os threads with go routines. We can still create multiple go routines with 1 os thread. These go routines will keep sharing the threads (go scheduler does it). 

- Q - say the goroutine 1 doesn't return the control to scheduler unless it is completed. Do the other goroutine have to wait
    In Go, a goroutine cannot completely prevent other goroutines from running, even if it doesn't explicitly yield control back to the scheduler. Here's why:
    Go's Scheduling Model:
    - Go uses a cooperative scheduling model with a non-preemptive approach (in most cases). This means goroutines are expected to cooperate by yielding control at specific points.
    - However, Go's runtime system can still intervene for essential tasks:
        - Garbage Collection: The runtime needs to pause goroutines occasionally to perform garbage collection. This ensures efficient memory management.
        - Fairness: While non-preemptive, Go aims for fairness in scheduling over time. If a goroutine runs for an excessively long period without yielding, the runtime might interrupt it to give other goroutines a chance to run.

