package carte

import tuile "server/carte/tuile"
import "fmt"

type Carte struct{
	size int
	matrice[][] tuile.Tuile
}
//Création de la Carte
func New(size int) Carte{
	var mat Carte
	mat.size=size
	mat.matrice=make([][]tuile.Tuile,size)
	for i := 0; i < size; i++ {
		mat.matrice[i]= make([]tuile.Tuile,size)
		mat.matrice[i][i%size]=tuile.New()
	}
	return mat
}
// Verifie si la cade est vide
func (c Carte) IsEmpty(x int, y int) bool{
		return tuile.GetType(c.matrice[x][y]) ==0
}

//Renvoie la tuile pour la position demandee
func (c Carte) GetTile(x int, y int) tuile.Tuile{
	return c.matrice[x][y]
}

//Affichage sur terminal
func Debug(mat Carte){
	for i:=0;i<mat.size;i++{
		for j:=0;j<mat.size;j++{
			fmt.Print(tuile.GetType(mat.matrice[i][j]))
			fmt.Print(" ");
		}
		fmt.Println("")
	}
}

type Case struct{
	x int
	y int
	tile tuile.Tuile
}

func validCoords(x int,y int, size int) bool{
	return ((x>=0 && x<size) && (y>=0 && y<size))
}

const UNVISITED=-1 //case non parcourue
const OBSTACLE=-2
func attribuerPoids(weightMatrix [][]int,x int,y int, step int) [][]int{
	if(validCoords(x,y,len(weightMatrix))){
		if(weightMatrix[x][y]==UNVISITED && weightMatrix[x][y]!=OBSTACLE){
			weightMatrix[x][y]=step+1
		}
	}
	return weightMatrix
}

func shortestPathAux(weightMatrix [][]int,c Carte, x int, y int, currX *int, currY *int, step int, path []Case) bool{
	if(weightMatrix[x][y]==step-1){
		path[step-1]=(Case{x,y,c.GetTile(x,y)})
		*currX=x
		*currY=y
		if(step==1){
			return true
		}
	}
	return false
}

//Selectionne le chemin le plus court a partir de la matrice des poids
func shortestPath(weightMatrix [][]int,destx int, desty int,c Carte,path []Case) []Case{
	if(weightMatrix[destx][desty]==UNVISITED){
		return (nil)
	}
	path=make([]Case,weightMatrix[destx][desty]+1)
	path[len(path)-1]=(Case{destx,desty,c.GetTile(destx,desty)})
	currX:=destx
	currY:=desty
	for step:=weightMatrix[destx][desty]+1;step>0;step--{
		if(shortestPathAux(weightMatrix,c, currX, currY+1, &currX, &currY, step,path) || shortestPathAux(weightMatrix,c, currX, currY-1, &currX, &currY, step,path) ||shortestPathAux(weightMatrix,c, currX+1, currY, &currX, &currY, step,path) || shortestPathAux(weightMatrix,c, currX-1, currY, &currX, &currY, step,path)){
			break
		}
	}
	return path
}
//Renvoie le chemin le plus court entre deux cases ou nil si inatteignable
func (c Carte) GetPathFromTo(x int, y int, destx int, desty int) []Case{
	var path []Case
	var weightMatrix [][]int
	weightMatrix=make([][]int,c.size)
	for i := 0; i < c.size; i++ {
		weightMatrix[i]= make([]int,c.size)
		for j:=0;j<c.size;j++{
			if(c.IsEmpty(i,j)){
				weightMatrix[i][j]=UNVISITED
			}else{
				weightMatrix[i][j]=OBSTACLE
			}
		}
	}
	weightMatrix[x][y]=0 //Case source visitee
	//Fin de l'initialisation
	for step:=0;step<c.size*c.size;step++{
		for i:=0;i<c.size;i++{
			for j:=0;j<c.size;j++{
				if(weightMatrix[i][j]==step){
					weightMatrix=attribuerPoids(weightMatrix,x,y+1,step) //Haut
					weightMatrix=attribuerPoids(weightMatrix,x+1,y,step) //Droite
					weightMatrix=attribuerPoids(weightMatrix,x,y-1,step) //Bas
					weightMatrix=attribuerPoids(weightMatrix,x-1,y,step) //Gauche
				}
			}
		}
	}
	//Une matrice de poids est cree
	return shortestPath(weightMatrix,destx, desty,c,path)
}