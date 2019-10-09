package router

import (
	"encoding/json"
	"net/http"
	"server/util"
	"fmt"
	"github.com/gorilla/mux"
)

type handler struct {
	utilService util.UtilService
}

func NewRouter(utilServ util.UtilService) http.Handler {
	var r = mux.NewRouter()

	var h = handler{
		utilService: utilServ,
	}

	r.HandleFunc("/reverse", h.reverseString).
		Methods(http.MethodGet)

	return r
}

func (h *handler) reverseString(w http.ResponseWriter, r *http.Request) {

	var response struct {
		Result string `json:"result"`
		Status string `json:"status"`
		Error string `json:"error,omitempty"`
	} 

	err := r.ParseForm()
	if err != nil {
		response.Status = http.StatusText(http.StatusBadRequest)
		response.Error = err.Error()
		answ, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(answ)
		return
	}

	var str = r.Form.Get("str")

	fmt.Printf(" str %v", str)

	result, err := h.utilService.Exec(str)
	
	if err != nil {
		response.Status = getErrorStatus(err.Error())
		response.Error = err.Error()
	}else{
		response.Result = result
	}

	answ, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(answ)
}

func getErrorStatus(err string) string{
	switch err {
	case util.BadInputError:
		return http.StatusText(http.StatusBadRequest)
	case util.BadResponseError:
		return http.StatusText(http.StatusBadRequest)
	default:
		return http.StatusText(http.StatusInternalServerError)
	}
}