# Simple Go Logger
Very Simple and light, level logger for golang.

## Usage

Initialize Logger 

```
import "github.com/adharshmk96/go-logger/gologger"

...
 
logger := gologger.New(gologger.Config{
    ColorLabel: true,
    ColorText:  true,
    LogLevel:   gologger.LogDebug,            
})

...
``` 
Log different levels as required

```
    logger.Debug("this %s logged", "can be")
    logger.Info("this %s logged", "can be")
    logger.Warning("this %s logged", "can be")
    logger.Error("this %s logged", "can be")
    logger.Fatal("this %s logged", "can be")
    logger.Trace("this %s logged", "can be")
```

The level of logging is as follows (low to high), the lower levels will be hidden once we set the log level in the config.
 
    LogDebug
    LogInfo
    LogWarning
    LogError
    LogFatal
    LogTrace