# usermodel

Given a task, return the user-configured LLM model by executing [`llm-manager`](https://github.com/xyproto/llm-manager), or else return the string in the `DefaultModel` variable (currently `gemma2:2b`).

### Example use

```go
package main

import (
    "fmt"

    "github.com/xyproto/usermodel"
)

func main() {
    fmt.Println(usermodel.GetCodeCompletionModel())
}
```

### Exported functions

```go
func AvailableTasks() []Task
func GetChatModel() string
func GetCodeCompletionModel() string
func GetTestModel() string
func GetTextGenerationModel() string
func GetToolUseModel() string
func GetTranslationModel() string
func GetVisionModel() string
func Get(task Task) string
```

### General info

* Version: 1.0.1
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
