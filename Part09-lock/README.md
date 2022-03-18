# Road to Go Pro — Locks

Please read the tutorial on [Medium](https://levelup.gitconnected.com/road-to-go-pro-async-part-3-locks-8bf60c476b12).

## Code with race conditions

The functions in `race.go` have race conditions. To run race condition test, use the following command.

```bash
go test -race -run Test_sumUpWithRace

go test -race -run Test_sumUpWithInspector
```

The unit test should fail and show this error in the console.

```
 testing.go:1152: race detected during execution of test
    testing.go:1152: race detected during execution of test
```

## Code without race conditions

```bash
go test -race -run TestTest_sumUpWithLock

go test -race -run TestTest_sumUpWithRWLock
```

The unit test should pass this time.

```
 PASS
ok      github.com/songx23/RoadToGoPro/Part9    0.135s
```