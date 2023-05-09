package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Factura struct {
	ID          primitive.ObjectID `bson:"_id"`
	Tipo        string             `bson:"tipo"`
	Numero      int                `bson:"numero"`
	Fecha       string             `bson:"fecha"`
	Cliente     string             `bson:"cliente"`
	Retira      string             `bson:"retira"`
	Rut         string             `bson:"rut"`
	Direccion   string             `bson:"direccion"`
	Email       string             `bson:"email"`
	NombreLibro string             `bson:"nombre"`
	Precio      int                `bson:"precio"`
	Cantidad    int                `bson:"cantidad"`
	Total       int                `bson:"total"`
}
