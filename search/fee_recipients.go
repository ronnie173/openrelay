package search

import (
	"encoding/json"
	"github.com/notegio/openrelay/affiliates"
	"net/http"
)

func FeeRecipientHandler(affiliateService affiliates.AffiliateService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryObject := r.URL.Query()
		errs := []ValidationError{}
		pageInt, perPageInt, err := getPages(queryObject)
		if err != nil {
			errs = append(errs, ValidationError{err.Error(), 1001, "page"})
		}
		affiliates, err := affiliateService.List()
		if err != nil {
			returnError(w, err, 500)
			return
		}
		if len(errs) > 0 {
			returnErrorList(w, errs)
			return
		}
		//total, page, per_page int, records
		pagedResult := GetPagedResult(len(affiliates), pageInt, perPageInt, affiliates[(pageInt - 1) * perPageInt:pageInt * perPageInt])
		response, err := json.Marshal(pagedResult)
		if err != nil {
			returnError(w, err, 500)
		} else {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		}
	}
}
