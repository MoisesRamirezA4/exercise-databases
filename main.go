package main

/*
	Servicios web:
		Cliente - servidor
		SPA = single page application
		Internet
			http
				verbos http
					GET: sirve para solicitar informacion al servidor y podemos especificar parametros a traves de query params (?)
					POST: sirve para crear recursos. NO LLEVA QUERY PARAMS PERO LLEVA UN BODY (Payload)
					PUT: sirve para hacer update de recursos
					PATCH: sirve para hacer modificaciones sobre recursos, ES MAS COMUN USAR PUT
					DELETE: sirve para eliminar algun recurso.
				codigos de respuesta http
					https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#2xx_success
					2xx -> 200, 201, 	 : La operacion fue exitosa
					3xx -> 300, 301 .... : Relacionado con la direccion del servidor
					4xx -> 400,401,404   : Errores en la peticion, pero el cliente puede solucionarlo
					5xx -> 500, 503,     : Errores en el lado del servidor

	Diferentes tipos de servios web API:
	API = application proggraming interface

	SOAP            		-> xml
	REST/Restful            -> JSON
	PROTOCOL(Grpc) -> Protocol buffer

	Enpoint:
		GET "google.com/api/v1/movies"

	DB:
		Relacionales
			Tablas, tipos de relaciones : uno a muchos , muchos a muchos
			Columnas: campos que describen el recurso, nombre, edad, direccion etc
			Filas:	recursos como tal
			MySQL, Postgres, oracleDB, MariaDB ...
			SQL: syntax query languange
		No Relacionales (NoSQL):
			Carecen de una estructura rigida.
			MongoDB: utiliza documentos JSON
			Cassandra: maps, partition keys etc
			Elasticsearch: indices + JSON
			Big Query: big tables de Google
			RedisDB: chache
			....
		Teorema de CAP:
			https://en.wikipedia.org/wiki/CAP_theorem
			C : consistencia
			A : Availability , Disponibilidad
			P : Particionamiento
	ORM:
		Object realationship mapping
		mapea structs - tablas
	Frameworks de Go:
		-gin : https://github.com/gin-gonic/gin
		-echo
		-gorillaMux
		...
	Middleware:
		Son funciones que validan y se ejecutan previamente a los handlers para instrucciones repetitivas ej. Validar el token en la peticion

	Peticiones a API's de terceros :
	https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7

	API con gin :
	https://www.youtube.com/watch?v=CzxEUDq9xiQ

	Docker:

		Imagen: Clon de algun servidor (PC)
		Contendor: Es una instancia de una imagen que es el recurso que se corre

	Docker Hub:
	Docker registry (Google):

	docker images  : lista las imagenes de docker que tienes descargadas
	docker ps :  lista los contenedores que estan corriendo, si le pasamos el parametro -a. Nos va a listar los contedores esten corriendo o no.
	docker start 'nombre_del_contendor': inicia un nuevo proceso corriendo ese contenedor
	docker run --name 'nombre_del_contenedor' -d 'nombre_de_imagen': corre un nuevo contenedor usando la imagen que se le pasa como parametro.
	docker stop 'nombre_del_contenedor': detener la ejecucion del contenedor especificado
	docker exec -it 'nombre_del_contenedor' bash: hace una conexion al contenedor especificado y ejecuta el comando especificado (ultimo parametro)
*/

import (
	"fmt"
	_ "log"
	_ "net/http"

	"github.com/epa-datos/first-api-training/entity"
	"github.com/epa-datos/first-api-training/repo"
)

/* func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye %s!", r.URL.Path[1:])
}

func hola() {
	fmt.Println("HOOOla")
} */

func main() {
	repo.Connect()
	repo.NewPost(&entity.Post{
		Name: "primer post",
	})
	repo.NewPost(&entity.Post{
		Name: "segundo post",
	})
	post, err := repo.GetPost(1)
	post2, err := repo.GetPost(2)
	fmt.Println(*post, err)
	fmt.Println(*post2, err)
	posts, err := repo.GetPosts()
	fmt.Println(posts)
}
