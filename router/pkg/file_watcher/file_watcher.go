package file_watcher

import (
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type ConfigFileWatcher struct {
	path    string
	watcher *fsnotify.Watcher
	logger  *zap.Logger
}

func FileWatcher(logger *zap.Logger, path string) (*ConfigFileWatcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &ConfigFileWatcher{
		logger: logger.With(
			zap.String("component", "file_watcher"),
			zap.String("path", path),
		),
		path:    path,
		watcher: watcher,
	}, nil
}

func (c *ConfigFileWatcher) Watch(cb func()) error {
	go func() {
		for {
			select {
			case event, ok := <-c.watcher.Events:
				if !ok {
					return
				}
				c.logger.Debug("notify event", zap.Any("event", event))
				if event.Has(fsnotify.Write) {
					cb()
				}
			case err, ok := <-c.watcher.Errors:
				if !ok {
					return
				}
				c.logger.Error("error", zap.Error(err))
			}
		}
	}()

	err := c.watcher.Add(c.path)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigFileWatcher) Close() error {
	return c.watcher.Close()
}
