# **ezflag** - A easy flag for your CLI


``` go
go get "github.com/ahopo/ezflag"
```
## Usage
``` go
import (
"fmt"
"github.com/ahopo/ezflag"
)
func main(){
    first_name:=ezflag.String("fn","first_name","","The first name of the person.")
    ezflag.Parse()
    fmt.Println("Hi",first_name)
}
```
## Default Commands:
``` ps1
mayapp -h or --help
#to view help
```

#### _Note: sort order of help list arguments will depend on the arrangement of the flag._

## Functions Parameters:
**short** : _the shortcut or a short word to describe a parameter_
```
e.g first_name shortcut to fn : mytool -fn juan
```
**long** : _the whole world of the parameter_
```
e.g first_name : mytool --first_name juan
```

**_default** : _the default form of the arugment._

**description**: _it describe what is the use of the argument._
 
