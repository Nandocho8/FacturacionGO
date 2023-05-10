package main

func Facturar (w http.ResponseWriter, r *http.Request){
	// genero conexion
	client := connection(mongoInfo)
	coll := client.Database("ventabookDB").Collection("books")

	// extraigo variables
	vars := mux.Vars(r)
	fmt.Printf("vars %s", vars)
	id := vars["id"]
	objectID, errID := primitive.ObjectIDFromHex(id)
	bodega := vars["bodega"]
	stock, errStock := strconv.Atoi(vars["stock"])

	// Maneja el error si el string no es un ObjectID válido
	if errID != nil {
		http.Error(w, "Invalid id value", http.StatusBadRequest)
		return
	}
	if errStock != nil {
		http.Error(w, "Invalid stock value", http.StatusBadRequest)
	}
	//genero filtro
	filter := bson.M{"_id": objectID, "ubicacion.bodega": bodega}

	update := bson.M{"$inc": bson.M{"ubicacion.$.stock": stock}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents updated: %v\n", result.ModifiedCount)
	defer client.Disconnect(context.Background())
}