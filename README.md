Based on http://www.robotamer.com/code/go/gotamer/gob.html

## Example usage

```go
saveErr := dump.Save("dump_"+region+".json", resp)
dump.Check(saveErr)
datafrom := new(lambda.ListFunctionsOutput)
loadErr := dump.Load("fixtures/dump_eu-central-1.json", datafrom)
dump.Check(loadErr)
fmt.Println("========")
fmt.Println(*datafrom.Functions[0].FunctionArn)
```