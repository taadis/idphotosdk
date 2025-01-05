# idphotossdk

hivision id photo sdk.

## Usage

get package

```
go get github.com/taadis/idphotosdk
```

code example

```
package main

import (
	"context"
	"fmt"

	"github.com/taadis/idphotosdk"
)

func main(){
	ctx := context.Background()
	client := idphotosdk.NewClient(
		idphotosdk.WithBaseURL("your_base_url"),
		idphotosdk.WithLogEnabled(true),
	)
	req := idphotosdk.NewIdphotoRequest()
	req.InputImageBase64 = "your_image_base64"
	rsp, err := client.Idphoto(ctx, req)
	if err != nil {
		panic(err)
	}
	fmt.Println("got rsp and to do somethings", rsp.Data)
}

```
