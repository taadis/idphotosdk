# hivision-id-photos-sdk-go

## Usage

get package

```
go get github.com/taadis/hivision-id-photos-sdk-go
```

code example

```
package idphotosdk

import (
	"context"
	"fmt"

    idphotosdk "github.com/taadis/hivision-id-photos-sdk-go"
)

func main(){
	ctx := context.Background()
	client := idphotosdk.NewClient(
		idphotosdk.WithBaseURL("your_base_url"),
		idphotosdk.WithLogEnabled(true),
	)
	req := idphotosdk.NewIdphotoRequest()
	rsp, err := client.Idphoto(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println("got rsp and to do somethings", rsp.Data)
}

```
