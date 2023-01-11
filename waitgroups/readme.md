A repo for displaying differnent Scenerios for using wait groups in a go routine 
- A WaitGroup waits for a collection of goroutines to finish. 

- The WaitGroup type of sync package, is used to wait for the program to finish all goroutines launched from the main function. It uses a counter that specifies the number of goroutines, and Wait blocks the execution of the program until the WaitGroup counter is zero.

- The Add method is used to add a counter to the WaitGroup.

- The Done method of WaitGroup is scheduled using a defer statement to decrement the WaitGroup counter.

- The Wait method of the WaitGroup type waits for the program to finish all goroutines.

- The Wait method is called inside the main function, which blocks execution until the WaitGroup counter reaches the value of zero and ensures that all goroutines are executed.