# Practicando
print "Practicando con Python:", __file__
matriz = []
x = 0
y = 0
file = open("./texto.txt","r")
for line in file:
    y = y + 1
    x = 0
    arlin = line.split(",")
    for variable in arlin:
        x = x + 1
        matriz[y][x] = int(variable)

print matriz
