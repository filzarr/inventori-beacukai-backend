package infrastructure

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

// var errSkipEvent = errors.New("skip")

// InitializeLogger will set logging format.
func InitializeLogger(stage string, filename string, logLevel zerolog.Level) {
	// pr, pw := io.Pipe()

	var (
		lumberjackLogger = &lumberjack.Logger{
			MaxSize:  100, // megabytes
			MaxAge:   1,   // days
			Filename: filename,
		}
		writers = []io.Writer{zerolog.ConsoleWriter{Out: os.Stderr}, lumberjackLogger}
		mw      = io.MultiWriter(writers...)
	)

	// messageQueue := make(chan []byte, 1000) // 1000 messages buffer (to prevent blocking the main thread)

	// // Create TelegramHook with rate limiting
	// telegramHook := &TelegramHook{
	// 	bot:      adapter.Adapters.VenamonGolog,
	// 	chatID:   config.Envs.VenamonGolog.ChatId,
	// 	threadID: config.Envs.VenamonGolog.ThreadId,
	// 	ticker:   time.NewTicker(3 * time.Second), // 20 messages per minute => 1 message every 3 seconds
	// 	stop:     make(chan struct{}),
	// }

	// go telegramHook.startConsume(messageQueue)
	// go telegramHook.SendMessages(pr, messageQueue)

	// using json format for production
	var logger zerolog.Logger
	if stage == "production" {
		logger = zerolog.New(lumberjackLogger).With().Timestamp().Caller().Logger().Level(zerolog.InfoLevel)
	} else {
		logger = zerolog.New(mw).With().Timestamp().Caller().Logger().Level(logLevel)
	}
	log.Logger = logger

	q := make(chan os.Signal, 1)
	c := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for {
			<-q
			lumberjackLogger.Close()
			log.Info().Msg("Closing logs ...")
			// telegramHook.Stop()
			log.Info().Msg("Closing telegram hook ...")
		}
	}()
	go func() {
		for {
			<-c
			if err := lumberjackLogger.Rotate(); err != nil {
				log.Error().Err(err).Msg("Error while rotating logs")
			}
			log.Info().Msg("Rotating logs ...")
		}
	}()
}

// type TelegramHook struct {
// 	bot      *telebot.Bot
// 	chatID   int64
// 	threadID int
// 	ticker   *time.Ticker
// 	stop     chan struct{}
// }

// func (h *TelegramHook) startConsume(messageQueue <-chan []byte) {
// 	for {
// 		select {
// 		case <-h.ticker.C:
// 			select {
// 			case msg := <-messageQueue:
// 				msgStr := string(msg)
// 				msgFix := `<pre language="json">` + msgStr + "</pre>"

// 				_, err := h.bot.Send(&telebot.Chat{ID: h.chatID, Type: telebot.ChatGroup}, msgFix, &telebot.SendOptions{
// 					ThreadID:  h.threadID,
// 					ParseMode: telebot.ModeHTML,
// 				})
// 				if err != nil {
// 					slog.Error("failed to send message to telegram", "error", err)
// 				}
// 			default:
// 				// No message in queue
// 			}
// 		case <-h.stop:
// 			log.Info().Msg("Stopping message queue processing")
// 			return
// 		}
// 	}
// }

// func (h *TelegramHook) SendMessages(pr *io.PipeReader, messageQueue chan<- []byte) {
// 	dec := json.NewDecoder(pr)

// 	for {
// 		var message = make(map[string]any)
// 		err := dec.Decode(&message)
// 		if err == io.EOF {
// 			return
// 		}

// 		if err == errSkipEvent {
// 			slog.Info("Skip event")
// 			continue
// 		}

// 		if err != nil {
// 			slog.Error("failed to decode message for json decoder", "error", err)
// 			continue
// 		}

// 		msgBytes, err := json.MarshalIndent(message, "", "  ")
// 		if err != nil {
// 			slog.Error("failed to marshal message", "error", err)
// 			continue
// 		}

// 		go func() { // send message to channel in goroutine, so it won't block the main thread if the channel is full
// 			messageQueue <- msgBytes
// 		}()
// 	}
// }

// func (h *TelegramHook) Stop() {
// 	log.Info().Msg("Stopping telegram hook ...")
// 	close(h.stop)
// 	h.ticker.Stop()
// }
