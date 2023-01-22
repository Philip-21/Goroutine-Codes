The sleeping barber problem is a classic inter-process communication and synchronization problem that illustrates the complexities that arise when there are multiple operating system processes.

 - PROBLEM SPECIFICATION
This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
sleep until a new client arrives. So the rules are as follows:
	- if there are no customers, the barber falls asleep in the chair
	- a customer must wake the barber if he is asleep
	- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
	  sits in an empty chair if it's available
	- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
	  and falls asleep if there are none
	- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is empty
	- after the shop is closed and there are no clients left in the waiting area, the barber  goes home

 - The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
semaphores (mutexes) is not needed.
 - Channels are used to solve the problem 