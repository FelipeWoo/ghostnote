package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"ghostnote/cmd/cli"
	"ghostnote/internal/bootstrap"
	"ghostnote/internal/config"
	"ghostnote/internal/logger"

	s "ghostnote/internal/services"
	st "ghostnote/internal/storage"
	t "ghostnote/internal/transport/cli"
)

func main() {
	bootstrap.InitAll()

	logger.SetModule("main")
	logger.Info("Starting %s app", config.Config.AppName)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Canal de señales
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	logger.Info("Ready to process requests ...")

	go func() {
		<-sigChan
		logger.Warn("Canceled by Ctrl+C")
		cancel()
	}()

	dir := "./vault"

	if err := ensureDir(dir); err != nil {
		logger.Fatal("%v", err)
		cancel()
	}

	repo := st.NewNoteRepository(dir)
	noteService := s.NewNoteService(repo)
	noteCommand := t.NewNoteCommand(noteService)

	err := cli.Menu(ctx, noteCommand)
	if err != nil {
		logger.Fail("app error: %v", err)
	}

	logger.Info("Shutting down...")
}

func ensureDir(path string) error {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}

	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("path exists but is not a directory: %s", path)
	}

	return nil
}
