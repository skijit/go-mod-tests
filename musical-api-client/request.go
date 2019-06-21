package request

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func DoSomething() {

	endpt := "http://localhost:5000/api/v1/submitted-events"
	// resp, body, errs := gorequest.New().
	_, body, _ := gorequest.New().
		Post(endpt).
		Send(`{"pageSize": 1, "pageNum": -1, "sortField": "submittedDt", "isSortAsc": false, "isPending": true}`).
		End()

	fmt.Println(body)

}
