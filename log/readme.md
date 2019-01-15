# logger 


We provided a setting strut for settings of log

```go
type LogSettings struct {
	Output       string `json:"output" yaml:"output" ini:"output"`
	Format       string `json:"format" yaml:"format" ini:"format"`
	Level        string `json:"level" yaml:"level" ini:"level"`
	ReportCaller bool   `json:"reportCaller" yaml:"report-caller" ini:"report-caller"`
}
```

You can use it like this:

```go
func TestLog(t *testing.T) {
	// these ni can be replaced by any struct containing the struct 'LogSettings'
	l := GetLogger(nil)
	l.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	contextLogger := l.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other": "I also should be logged always",
	})
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
	l.Trace("Something very low level.")
	l.Debug("Useful debugging information.")
	l.Info("Something noteworthy happened!")
	l.Warn("You should probably take a look at this.")
	l.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	l.Fatal("Bye.")
	// Calls panic() after logging
	l.Panic("I'm bailing.")

}
```
