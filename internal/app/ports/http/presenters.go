package http

//type saveEventRequest struct {
//	Module string `json:"module"`
//	Type   string `json:"type"`
//	Event  string `json:"event"`
//	Name   string `json:"name"`
//	Data   struct {
//		Action string `json:"action"`
//	} `json:"data"`
//}

func EventErrorResponse(status string, err error) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["status"] = status
	resp["error"] = err

	return resp
}

func EventSuccessResponse(status string) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["status"] = status
	resp["error"] = nil

	return resp
}
