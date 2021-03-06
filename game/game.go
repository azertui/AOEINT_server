package game

import Carte "git.unistra.fr/AOEINT/server/carte"
import "os"
import "git.unistra.fr/AOEINT/server/joueur"
import "git.unistra.fr/AOEINT/server/ressource"
import "git.unistra.fr/AOEINT/server/batiment"
import "git.unistra.fr/AOEINT/server/npc"
import "time"
import "fmt"
import "encoding/json"
import "git.unistra.fr/AOEINT/server/constants"
import "io/ioutil"

//Game : Structure contenant les donnees principales d'une partie
type Game struct{
	Joueurs []joueur.Joueur
	Carte Carte.Carte
	GameRunning bool
}

//Data :Structure permettant de stocker les informations recuperees sur le fichier json
type Data struct{
	Size int
	Buildings []batiment.Batiment
	Ressources []ressource.Ressource
}

//ExtractData : extract data from a file (ressources, buildings)
func ExtractData() Data{
	getEnvData()
	datafileName:="data/GameData.json"
	if(constants.UseSmallMap){
		datafileName="data/SmallTestMap.json"
	}
	jsonFile, err:= os.Open(datafileName)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	byteValue,_ := ioutil.ReadAll(jsonFile)
	var newGame Data
	err=json.Unmarshal(byteValue, &newGame)
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	err=jsonFile.Close()
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	return newGame
}

func getEnvData(){
	if(len(os.Getenv("GAME_UUID"))==0){
		constants.GameUUID = "DEFAULT"
		fmt.Println("default for GAME_UUID")
	}else{
		constants.GameUUID = constants.GameUUIDDef
	}
	if(len(os.Getenv("API_HOST"))==0){
		constants.APIHost = "DEFAULT"
		fmt.Println("default for API_HOST")
	}else{
		constants.APIHost = constants.APIHostDef
	}
	if(len(os.Getenv("TOKEN"))==0){
		constants.Token = "DEFAULT"
		fmt.Println("default for TOKEN")
	}else{
		constants.Token = constants.TOKENDef
	}
	if(len(os.Getenv("TOKEN_SECRET"))==0){
		constants.TokenSecret = "DEFAULT"
		fmt.Println("default for TOKEN_SECRET")
	}else{
		constants.TokenSecret = constants.TOKENSecretDef
	}
}

//GetPlayerFromUID : Permet de recuperer l'instance d'un joueur à partir de son uid
func (g Game)GetPlayerFromUID(uid string) *joueur.Joueur{
	for i:=0;i<len(g.Joueurs);i++{
		if(g.Joueurs[i].UID==uid){
			return &(g.Joueurs[i])
		}
	}
	return nil
}
//EndOfGame : Interromps la boucle principale du jeu
func (g *Game)EndOfGame(){
	(*g).GameRunning=false
}
//GameLoop : fonction contenant la boucle principale du jeu
func (g *Game)GameLoop(){
	for (*g).GameRunning{
		time.Sleep(time.Duration(1000000000))
	}

}
//GenerateMap : Permet de generer la Carte a partir d'une structure data
func (g *Game)GenerateMap(data Data){
	(*g).Carte =Carte.New(data.Size)
	//On attribue les auberges
	if(len((*g).Joueurs)==2){//Si Seulement 2 Joueurs fournis, fait en sorte de leur donner des bases adverses
		//ajout des npc de base
		pnj,id:=npc.Create("villager",10,10,g.Joueurs[0].GetFaction(),(&(g.Joueurs[0])).GetChannel())
		pnj.Transmit(id)
		g.Joueurs[0].AddNpc(&pnj)
		(*g).Joueurs[0].AddBuilding(&data.Buildings[0])
		if((*g).Carte.AddNewBuilding(&(data.Buildings[0]))==false){
			fmt.Println("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}
		(*g).Joueurs[1].AddBuilding(&data.Buildings[2])
		if((*g).Carte.AddNewBuilding(&(data.Buildings[2]))==false){
			fmt.Println("Erreur lors du placement d'une auberge")
			os.Exit(1)
		}

	}else{//sinon 4 Joueurs classiques dans l'ordre des bases fournies (blue blue red red)
		for i:=0;i<4;i++{
			(*g).Joueurs[i].AddBuilding(&data.Buildings[i])
			if((*g).Carte.AddNewBuilding(&(data.Buildings[i]))==false){
				fmt.Println("Erreur lors du placement d'une auberge")
				os.Exit(1)
			}
		}
	}
	for i:=0;i<len(data.Ressources);i++{
		(&data.Ressources[i]).InitiatePV()
		if((*g).Carte.AddNewRessource(&(data.Ressources[i]))==false){
			fmt.Println("Erreur lors du placement d'une ressource")
			os.Exit(1)
		}
	}
}
/*GetPlayerData : Recupere les donnes des Joueurs entree en parametre du programme
Modification: Changement pour des valeurs statiques (temporaire)
*/
func (g *Game)GetPlayerData(){
	(*g).Joueurs=make([]joueur.Joueur,2)
	(*g).Joueurs[0]=joueur.Create(false,"Bob","V8F1238VF")
	(*g).Joueurs[1]=joueur.Create(true,"Alice","1982N19N2")
	fmt.Println("joueurs:",(*g).Joueurs[0].GetNom(),"",(*g).Joueurs[1].GetNom())
}
