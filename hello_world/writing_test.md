# Golang learn by testing

## Writing test

Writing test in `Golang` is like writing a function with a few rules.

- It needs to be in a file with a name `xxx_test.go`
- The test function must start with the word `Test`
- The test function take one arguments only `t *testing.T`
- The `t` in the type `*testing.T` is an **hook** to the testing framework. We can also use it with for example : `t.Fail()`.
