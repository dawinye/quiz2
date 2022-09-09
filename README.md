# quiz2

run go build one.go and go build two.go to build your own executables,
then run ./one 80 and ./two 80 or any other random number to time

https://www.reddit.com/r/golang/comments/vhsw42/usecase_for_using_sortslice_over_sortslicestable/
SliceStable is faster although perhaps it shouldn't be, it uses a form of mergesort which is always O(nlogn), while Sort() should be O(nlogn) in the worst case.
