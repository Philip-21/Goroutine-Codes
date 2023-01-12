 The dining philosophers problem is an example problem often used in concurrent algorithm design to illustrate synchronization issues and techniques for resolving them.


The Dining Philosophers problem is well known in computer science circles.
Five philosophers, numbered from 0 through 4, live in a house where the
table is laid for them; each philosopher has their own place at the table.
Their only difficulty – besides those of philosophy – is that the dish
served is a very difficult kind of spaghetti which has to be eaten with
two forks. There are two forks next to each plate, so that presents no
difficulty. As a consequence, however, this means that no two neighbours
may be eating simultaneously, since there are five philosophers and five forks.
This is a simple implementation of Dijkstra's solution to the "Dining Philosophers" dilemma.

The problem was designed to illustrate the challenges of avoiding deadlock, a system state in which no progress is possible. 