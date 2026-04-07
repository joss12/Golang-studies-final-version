package main

import (
	"context"
	"log"
	"time"

	pb "github.com/joss12/grpc-notes/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewNoteServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// --- 1. Create three notes ---
	log.Println("=== Creating Notes ===")
	ids := []string{}
	for _, title := range []string{"Buy groceries", "Learn gRPC", "Read a book"} {
		resp, err := client.CreateNote(ctx, &pb.CreateNoteRequest{
			Title:   title,
			Content: title + "- content here",
		})
		if err != nil {
			log.Fatalf("CreateNote failed: %v", err)
		}
		log.Printf("Created: [%s] %s", resp.Note.Id, resp.Note.Title)
		ids = append(ids, resp.Note.Id)
	}

	// --- 2. List all notes ---
	log.Println("\n=== Listing all notes ===")
	listResp, err := client.ListNotes(ctx, &pb.ListNotesRequest{})
	if err != nil {
		log.Fatalf("ListNotes failed: %v", err)
	}
	for _, n := range listResp.Notes {
		log.Printf(" - [%s]%s", n.Id, n.Title)
	}

	// --- 3. Get a single note ---
	log.Println("\n=== Getting first note ===")
	getResp, err := client.GetNote(ctx, &pb.GetNoteRequest{Id: ids[0]})
	if err != nil {
		log.Fatalf("GetNote failed: %v", err)
	}
	log.Printf("Fetched: %s — %s", getResp.Note.GetTitle(), getResp.Note.GetContent())

	// --- 4. Delete the first note ---
	log.Println("\n=== Deleting first note ===")
	delResp, err := client.DeleteNote(ctx, &pb.DeleteNoteRequest{Id: ids[0]})
	if err != nil {
		log.Fatalf("DeleteNote failed: %v", err)
	}
	log.Printf("Deleted: %s", delResp.GetMessage())

	// --- 5. Try to get deleted note - expect NotFound ---
	log.Println("\n=== Featching deleted note (expect error) ===")
	_, err = client.GetNote(ctx, &pb.GetNoteRequest{Id: ids[0]})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			log.Printf("Got expected error — code: %s, message: %s", st.Code(), st.Message())
		} else {
			log.Fatalf("Unexpected error: %v", err)
		}
	}

	// --- 6. Try to create a note with empty title — expect InvalidArgument ---
	log.Println("\n=== Creating note with empty title (expect error) ===")
	_, err = client.CreateNote(ctx, &pb.CreateNoteRequest{Title: ""})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.InvalidArgument {
			log.Printf("Got expected error — code: %s, message: %s", st.Code(), st.Message())
		} else {
			log.Fatalf("Unexpected error: %v", err)
		}
	}
}
