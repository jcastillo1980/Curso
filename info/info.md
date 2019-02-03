# Ejemplos de iptable

- sudo su
- hay 3 tablas filter,nat,mangle. cada table tienes cadeneas, y cada cadena tiene reglas
- filter: INPUT, OUTPUT,FORWARD
- nat: OUTPUT,PREROUTING,POSTROUTING
- mangel: INPUT,OUTPUT,FORWARD,PREROUTING, POSTROUTING
- iptables -L  : lista tabla filter
- iptables -t nat -L  : lista tabla nat
- iptables -F  :  borra todo
- iptables -A INPUT -p tcp --dport 22 -j DROP : el puerto entrada 22 de tcp no entra
- iptables -A INPUT -p tcp --dport 1:1024 -j DROP : los puertos entrada 1..1024 de tcp no entra a los procesos
- iptables -P INPUT DROP:  rechaza por defecto filter:INPUT

- iptables -L --lines-numbers : lista numerada
- iptables -R INPUT 1  .... : modifica la linea 1 de la ccadena INPUT

- iptables -N LOGGING : nueva cadena
- iptables -Z : borra contadores

- iptables -A INPUT -p tcp -s 192.168.0.9 --dport 22 -j DROP
- iptables -A INPUT ! -s 192.168.0.9 -p tcp  --dport 22 -j DROP : menos un pc concreto
- iptables -A INPUT -p tcp -m multiport --dport 21,33,777 -j DROP : puertos aleatorios

- iptables -A INPUT -m mac --mac-source 00:00:00:00:00:00 -j DROP

- service rsyslog status : los log del sitema , haber si estan funcionando
- watch tail /var/log/message
- iptables -A INPUT -p icmp -j LOG --log-prefix "JAVIERCASTILLO"
- grep "JAVIERCASTILLO" /var/log/message

## añadir solo si he iniciado yo
- iptables -P OUTPUT DROP
- iptables -P INPUT DROP
- iptables -A OUTPUT -m state --state NEW,ESTABLISHED -j ACCEPT : comprueba por estado (tcp) y lo acepta
- iptables -A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT 

## copias y restauracion
- iptables-save > file.dat
- iptables-restore < file.dat

## NAT
- SNAT: para salir internet, DNAT: para abrir puertos
- echo 1 > /proc/sys/net/ipv4/ip_forward
- iptables -t nat -A POSTROUTING -o eth0 -s 192.168.0.0/24 -j  MASQUERADE :  es SNAT --to IP_PUBLIC en modo automatico
- iptables -t nat -A PREROUTING -p tcp --dport 80 -i eth1 -j DNAT --to 5.6.7.8:8080

# Ejemplos de route

- route -e : lista
- netstat -rn : lista tambien
- route add -net 192.168.0.0 netmask 255.255.255.0 gw 192.168.1.1 dev eth0 : añade 192.168.0.0/24 que salga por 192.168.1.1 que esta en eth0
- si añade -p : se queda en permamente !! pero no en todos las las distro (Mejor añadir a /etc/rc.local)
- route del -net 192.168.0.0 netmask 255.255.255.0 gw 192.168.1.1 dev eth0 :  borra la regla
- route add -net 169.255.0.0/16 dev eth0  : forma simplificada.
- sudo route add default gw 192.168.2.1 dev eth0: poner el de por defecto


# algo de configuracion
- ifconfig eth0 up  : para apagar poner down
- ifconfig eth0 172.168.0.22 netmask 255.255.255.0
- route add default gw 172.168.0.1
