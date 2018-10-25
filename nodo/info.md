# Funcionamiento

## modulos necesarios

- go get golang.org/x/text/encoding/unicode
- go get github.com/go-sql-driver/mysql
- go get github.com/denisenkom/go-mssqldb
- go install github.com/denisenkom/go-mssqldb

## nuevos campos BLX2

- Añadimos campos, para determinar en que nodo hay que conectarse para acceder. 
- el dispositivo se conecta a un nodo: el nodo comprueba que su mac es la ip: #MAC
- el programa lectura detecta que ip es #MAC y saca los campos que dan la información a que nodo ip/purto hay que conectarse
- el programa de lectura se conecta al nodo mandadno $MAC
- Los campos a añadir:
  - cliente_mac_ts
  - cliente_nodo_ip
  - cliente_nodo_port
  - cliente_nodo_id