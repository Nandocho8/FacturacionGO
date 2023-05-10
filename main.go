package main

func main() {
	factura := FacturarHtml()
	println(factura)
	FacturarPdf()


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/factura", corsMiddleware(facturar)).Methods("POST")

}
