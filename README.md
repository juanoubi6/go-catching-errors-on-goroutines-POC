The idea is to test the Errgroup feature that allows us to check if an error was raised in any gouroutine when firing a big amount of them. This is specially usefull when we want to split a big task across many gouroutines and we want to know if any of those small tasks failed.