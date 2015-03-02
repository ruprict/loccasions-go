package app

import (
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}
	http.Error(w, "Sorry, an error occurred", http.StatusInternalServerError)

	/*	if err, ok := err.(httpError); ok {
			http.Error(w, err.Error(), err.Status)
			return
		}

		if err, ok := err.(validationFailure); ok {
			renderJSON(w, err, http.StatusBadRequest)
			return
		}

		if isErrSqlNoRows(err) {
			http.NotFound(w, r)
			return
		}

		logError(err)

		http.Error(w, "Sorry, an error occurred", http.StatusInternalServerError)
	*/
}
