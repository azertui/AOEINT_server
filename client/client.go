//Contient toutes les fonctions pour les echanges de donnees client/serveur
package client

import (
	"git.unistra.fr/AOEINT/server/npc"
	/*"git.unistra.fr/AOEINT/server/carte"
	"git.unistra.fr/AOEINT/server/ressource"
	"git.unistra.fr/AOEINT/server/joueur"
	"git.unistra.fr/AOEINT/server/batiment"*/

	"fmt"
	"context"
	"log"
	"net"
	"google.golang.org/grpc"
	pb "git.unistra.fr/AOEINT/server/serveur"
)

///////////////////////////////////////////////////////////////////////////////
// Général
///////////////////////////////////////////////////////////////////////////////

var server *grpc.Server

type Server struct {}

// Fonction demarrant la gestion des intéractions gRPC
// Fonction bloquante, à lancer en concurrence
func InitListenerServer(adress string) {

	// Initialisation du socket d'écoute réseau
	lis, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialisation du serveur gRPC
	server = grpc.NewServer()

	// Enregistement des services Hello, Map et Interactions sur le serveur
	pb.RegisterHelloServer(server, &Server{})
	pb.RegisterMapServer(server, &Server{})
	pb.RegisterInteractionsServer(server, &Server{})

	// Mise en écoute du serveur
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}

// Fonction arrêtant la gestion des intéractions gRPC (arrêt propre)
func StopListenerServer() {
	server.GracefulStop()
}

// Fonction arrêtant la gestion des intéractions gRPC (arrêt brutal)
func KillListenerServer() {
	server.Stop()
}


///////////////////////////////////////////////////////////////////////////////
// Serveur -> Client
///////////////////////////////////////////////////////////////////////////////

//Envoie toutes les donnees necessaires à la mise en place de la partie en debut de jeu
//A envoyer: donnees des joueurs, structure data(map), entites de depart..
func InitGameState() {}

//Maj les ressources du joueur à partir de l'uid correspondant
func updatePlayerRessources(playerUID string,stone int,wood int,food int){}

//Maj: Indique la destruction d'un Batiment au client pour qu'il soit retire
func BuildingDestroyed(playerUID string,x int, y int){

}

//Permet de Maj la liste des npcs visibles en indiquant leur mort au client
func PlayerNpcsDied(playerUID string,npc []npc.Npc){

}

///////////////////////////////////////////////////////////////////////////////
// Client -> Serveur
///////////////////////////////////////////////////////////////////////////////

// Fonction du service Hello: SayHello
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("Reception d'un HelloRequest et envoie d'un HelloReply")
	return &pb.HelloReply{}, nil
}

// Fonction du service Map: GetMap
func (s *Server) GetMap(ctx context.Context, in *pb.GetMapRequest) (*pb.GetMapReply, error) {
	fmt.Println("Reception d'un GetMapRequest et envoie d'un GetMapReply")
	return &pb.GetMapReply{}, nil
}

// Fonction du service Map: SetMap
func (s *Server) SetMap(ctx context.Context, in *pb.SetMapRequest) (*pb.SetMapReply, error) {
	fmt.Println("Reception d'un SetMapRequest et envoie d'un SetMapReply")
	return &pb.SetMapReply{}, nil
}

// Fonction du service Map: UpdateMap
func (s *Server) UpdateMap(ctx context.Context, in *pb.UpdateMapRequest) (*pb.UpdateMapReply, error) {
	fmt.Println("Reception d'un UpdateMapRequest et envoie d'un UpdateMapReply")
	return &pb.UpdateMapReply{}, nil
}


// Fonction du service Interactions: RightClick
func (s *Server) RightClick(ctx context.Context, in *pb.RightClickRequest) (*pb.RightClickReply, error) {
	fmt.Println("Reception d'un RightClickRequest et envoie d'un RightClickReply")
	return &pb.RightClickReply{}, nil
}

// Fonction du service Map: MoveTo
func (s *Server) MoveTo(ctx context.Context, in *pb.MoveToRequest) (*pb.MoveToReply, error) {
	fmt.Println("Reception d'un MoveToRequest et envoie d'un MoveToReply")
	return &pb.MoveToReply{}, nil
}

//demande la creation d'un batiment à partir de l'uid du joueur, une position et un type de batiment
//class: "auberge","caserne","etabli"
func TryToBuild(playerUID string, x int, y int, class string) bool{
	return false
}

//Demande le deplacement des npc selectionnes
func MoveSelectedNpc(playerUID string, liste []npc.Npc, x int, y int){

}

//Demande la suppression par le joueur de l'un de ses batiments
func EraseBuilding(playerUID string, x int, y int){

}

//Averti le serveur de la creation d'une entite: verification des ressources necessaires
func AddNewNpc(playerUID string, x int, y int, typ int) bool{
	return false
}
//Enleve des Pv a un batiment
func  DamageBuilding(playerUID string, x int, y int, attack int){

}