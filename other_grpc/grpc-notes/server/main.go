package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	pb "github.com/joss12/grpc-notes/proto"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type notesServer struct {
	pb.UnimplementedNoteServiceServer
	mu    sync.RWMutex
	notes map[string]*pb.Note
}

func (s *notesServer) CreateNote(ctx context.Context, req *pb.CreateNoteRequest) (*pb.CreateNoteResponse, error) {
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title cannot be empty")
	}
	note := &pb.Note{
		Id:      uuid.NewString(),
		Title:   req.Title,
		Content: req.Content,
	}

	s.mu.Lock()
	s.notes[note.Id] = note
	s.mu.Unlock()

	log.Printf("Created note: %s", note.Id)
	return &pb.CreateNoteResponse{Note: note}, nil
}

func (s *notesServer) GetNote(ctx context.Context, req *pb.GetNoteRequest) (*pb.GetNoteResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	s.mu.RLock()
	note, ok := s.notes[req.Id]
	s.mu.RUnlock()

	if !ok {
		return nil, status.Errorf(codes.NotFound, "note with id %q not found", req.Id)
	}

	return &pb.GetNoteResponse{Note: note}, nil
}

func (s *notesServer) ListNotes(ctx context.Context, req *pb.ListNotesRequest) (*pb.ListNotesResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notes := make([]*pb.Note, 0, len(s.notes))
	for _, n := range s.notes {
		notes = append(notes, n)
	}

	return &pb.ListNotesResponse{Notes: notes}, nil
}

func (s *notesServer) DeleteNote(ctx context.Context, req *pb.DeleteNoteRequest) (*pb.DeleteNoteResponse, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.notes[req.Id]; !ok {
		return nil, status.Errorf(codes.NotFound, "note with id %q not found", req.Id)
	}

	delete(s.notes, req.Id)
	log.Printf("Deleted note: %s", req.Id)
	return &pb.DeleteNoteResponse{Message: "note deleted successfully"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNoteServiceServer(s, &notesServer{
		notes: make(map[string]*pb.Note),
	})

	fmt.Println("Notes server listening on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
