package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 4392

	http.HandleFunc("/get-excel-file", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "could not find resource", http.StatusNotFound)
			return
		} else {
			fileXslx, err := generateXslxFile()
			if err != nil {
				http.Error(w, "There was an internal error:", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Disposition", "attachment; filename=info.xlsx")
			w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheethtml.sheet")
			err = fileXslx.Write(w)
			if err != nil {
				http.Error(w, "xslx file could not be created", http.StatusInternalServerError)
				return
			}
		}

	})

	fmt.Printf("listening at http://localhost:%d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
