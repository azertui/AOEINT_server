package joueur

import npc "server/npc"
import batiment "server/batiment"
import constants "server/constants"

type Joueur struct{
	faction bool //true: faction 1, false: faction 2
	nom string
	nbats int
	batiments[] batiment.Batiment
	nelems int
	entities[] npc.Npc
	id byte
	stone int
	wood int
	food int
}
var model byte =0//Permet d'obtenir des id uniques lors d'une partie

//Crée un joueur
func Create(faction bool,nom string) Joueur{
	res :=Joueur{faction,nom,0,make([]batiment.Batiment,constants.MaxBuildings),0,make([]npc.Npc,constants.MaxEntities),model,constants.StartingStone,constants.StartingWood,constants.StartingFood}
	model++
	return res
}
//Retourne la faction
func (j Joueur) GetFaction() bool{
	return j.faction
}
//Retourne le Nom
func (j Joueur) GetNom() string{
	return j.nom
}
//Retourne l'id jouer
func (j Joueur) GetId() byte{
	return j.id
}

//Retourne la quantité de d'une ressource d'un joueur
func (j Joueur) GetStone() int{
	return j.stone
}
func (j Joueur) GetWood() int{
	return j.wood
}
func (j Joueur) GetFood() int{
	return j.food
}

//ajout de ressources
func (j *Joueur) AddStone(s int){
	(*j).stone +=s
}
func (j *Joueur) AddWood(w int){
	(*j).wood +=w
}
func (j *Joueur) AddFood(f int){
	(*j).food+= f
}
