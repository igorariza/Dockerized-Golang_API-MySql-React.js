# Dockerized Golang-API MySql React.js |<img src="https://user-images.githubusercontent.com/18409088/130371581-5ab1814a-a489-4a23-ac2c-e780a416aa3f.png" alt="docker logo" width="32">|<img src="https://raw.githubusercontent.com/Delta456/Delta456/master/img/golang.png" alt="go logo" width="38">|<img src="https://user-images.githubusercontent.com/18409088/129100018-c75e1ca5-3c0d-4f2a-949a-2d376aae09be.png" alt="react native logo" width="24">|<img src="https://user-images.githubusercontent.com/18409088/134431364-4f3f8f74-4fb7-4558-a659-e53cbf0ebfd4.png" alt="mysql logo" width="38">




#### Init:

- `docker-compose build` 

### Backend
Script para extraer información del dataset ubicado en la API de FUT21 y ser agregarda a base de datos para la consulta y modificación.
`Bases de datos MySQL`. <br>
Url de la API:
https://www.easports.com/fifa/ultimate-team/api/fut/item
Se itera en hasta el numero de pagina que retorna el response del endpoint y guarda los jugadores en la base de datos, se puede utilizar "page" para solicitar un
número de página específico a la API (por ejemplo ?page=1). <br>
Los jugadores están en el response en el campo `ítems`, de estos interesa almacenar el nombre, la posición en que juega, su nacionalidad y el nombre del
equipo al que pertenece.

### API REST 

GET `/api/v1/players` <br>
Busca los jugadores que contengan el String en los campos del nombre del jugador, ya sea una coincidencia parcial o total, y sin importar si es mayúscula o
minúscula. <br>
El order puede ser asc o desc y define el orden a partir del nombre alfabéticamente, por default sera asc (si no se recibe en la url).
