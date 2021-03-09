# go-concurrency-basic-wait-group

## Descripción
- En este ejemplo implementamos el uso de Wait-Group. Realizamos varias peticiones http de tipo GET a distintas URL's que se encuentran en un arreglo.

- La función main posee un handleFunc que utiliza checkStatus. 

- Creamos una función checkStatus para comprobar el estado de la url (200, 500, 429, etc). 

- Creamos una variable de tipo waitgroup, recorremos las URL del arreglo y hacemos la comprobación del status. Por cada comprobación de la URLle seteamos al waitGroup.Done() para avisar que finalizó el proceso.

- Una vez terminado el recorrido del arreglo, finalizamos con el waitGroup.Wait()

De esta manera con el wait groupp nos aseguramos que se muestren todos los resultados. Sin el wait group, sería imposible ya que la función checkStatus termina la ejecución antes de que se hayan recibido las respuestas http. y por ende no mostraría ningún resultado. 


## Test

:heavy_check_mark: En la terminal
``` 
go run main.go
``` 

 :heavy_check_mark: Solicitamos las petición
```
http://localhost:8080/
```

![image](https://user-images.githubusercontent.com/32901911/110538300-a8fcf580-8102-11eb-954d-242afc35384d.png)
