import json

# Practicando
print "Practicando con Python:", __name__
matriz = []
file = open("/Users/javiercastillocalvo/Documents/gopath/src/github.com/jcastillo1980/Curso/practica_python/texto.txt","r")
for line in file:
    arlin = line.split(",")
    semimatriz = []
    for variable in arlin:
        semimatriz.append(int(variable))
    matriz.append(semimatriz)

print matriz,".....",matriz[1][0]
file.close()

cosa = [0,11,22,33,{
    'valor1':"aqui hay algo",
    'valor2':2.44,
    'valor3':True
}]
sss = json.dumps(cosa,indent=4)

print "->" , sss , "<-"
print type(json.loads(sss))

