package command

import (
	"flag"
	"fmt"
	"log/slog"
	"sync"

	"github.com/SUT-technology/download-manager-golang/internal/application/services"
	"github.com/SUT-technology/download-manager-golang/internal/infrastructure/db"
	"github.com/SUT-technology/download-manager-golang/internal/interface/config"
	"github.com/SUT-technology/download-manager-golang/internal/interface/handlers"
	"github.com/SUT-technology/download-manager-golang/ui"
	"github.com/SUT-technology/download-manager-golang/ui/tabs"
)

func Run() error {
	tabs.ClearScreen()
	var configPath string
	flag.StringVar(&configPath, "cfg", "assets/config/development.yaml", "Configuration File")
	flag.Parse()
	c, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("reading config: %w", err)
	}

	//logger := slogger.NewJSONLogger(c.Logger.Level, os.Stdout)
	//slog.SetDefault(logger)

	db, err := db.New(c.DB)
	if err != nil {
		return fmt.Errorf("connecting to database: %w", err)
	}
	slog.Debug("initialized json database")

	srvcs := services.New(db)

	// SAMPLE USE HANDLERS

	hndlrs := handlers.New(srvcs)

	// progChan := make(chan int)
	// var err1 error

	// // Create a goroutine for the download
	// go func() {
	// 	err1 = hndlrs.DownloadHndlr.CreateDownload(dto.DownloadDto{
	// 		URL:      "https://dl.nakaman-music.ir/Music/BAHRAM/Forsat/Bahram%20-%20Gear%20Box.mp3",
	// 		QueueID:  "a187f448-12f2-4203-9197-fc49425f05e8",
	// 		FileName: "testisaa",
	// 	}, progChan)
	// }()

	// if err1 != nil {
	// 	return err1
	// }

	// ticker := time.NewTicker(1 * time.Second)
	// defer ticker.Stop()

	// // Loop to read progress updates from the channel
	// for {
	// 	select {
	// 	case val, ok := <-progChan:
	// 		if !ok {
	// 			// Channel is closed, exit the loop
	// 			fmt.Println("Channel closed, exiting.")
	// 			return nil
	// 		}
	// 		// Print the received progress value
	// 		fmt.Println("Received value:", val)
	// 	case <-ticker.C:
	// 		// You could check for timeout or progress every second
	// 		fmt.Println("Waiting for progress update...")
	// 	}
	// }

	// RUN UI AND USE IT

	// wg := sync.WaitGroup{}

	var wg sync.WaitGroup
	wg.Add(1)

	var err1 error

	go func() {
		err1 = ui.Run(&wg, &hndlrs)
	}()

	if err1 != nil {
		return fmt.Errorf("running ui: %w", err)
	}

	// downloads, err := hndlrs.DownloadHndlr.GetDownloads()

	// if err != nil {
	// 	return err
	// }

	// fmt.Println(downloads)
	//
	//hndlrs.QueueHndlr.CreateQueue(dto.QueueDto{
	//	"tst22",
	//	"tmp/tmp2",
	//	0,
	//	0,
	//	entity.TimeInterval{},
	//})
	//queue, err := hndlrs.QueueHndlr.GetQueueById("2")
	//if err != nil {
	//	return fmt.Errorf("getting queue: %w", err)
	//}
	//
	//hndlrs.DownloadHndlr.CreateDownload(dto.DownloadDto{
	//	URL:      "https://dl.nakaman-music.ir/Music/BAHRAM/Forsat/Bahram%20-%20Gear%20Box.mp3",
	//	QueueID:  queue.ID,
	//	FileName: "bahram.mp3",
	//})

	wg.Wait()

	return nil
}
