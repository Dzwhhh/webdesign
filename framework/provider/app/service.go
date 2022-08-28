package app

import (
	"errors"
	"flag"
	"path/filepath"

	"github.com/demian/webdesign/framework"
	"github.com/demian/webdesign/framework/contract"
	"github.com/demian/webdesign/framework/util"
)

type AppService struct {
	contract.App
	container  framework.Container
	baseFolder string
}

func (app *AppService) Version() string {
	return "0.0.1"
}

func (app *AppService) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	// 获取命令行参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder 参数, 默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	// 命令行参数不存在, 使用默认路径
	return util.GetExecDirectory()
}

func (app *AppService) ConfigFolder() string {
	return filepath.Join(app.BaseFolder(), "config")
}

func (app *AppService) StorageFolder() string {
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app *AppService) LogFolder() string {
	return filepath.Join(app.StorageFolder(), "log")
}

func (app *AppService) RuntimeFolder() string {
	return filepath.Join(app.StorageFolder(), "runtime")
}

func (app *AppService) AppFolder() string {
	return filepath.Join(app.BaseFolder(), "app")
}

func (app *AppService) ConsoleFolder() string {
	return filepath.Join(app.AppFolder(), "console")
}

func (app *AppService) HttpFolder() string {
	return filepath.Join(app.AppFolder(), "http")
}

func (app *AppService) ProviderFolder() string {
	return filepath.Join(app.AppFolder(), "provider")
}

func (app *AppService) MiddlewareFolder() string {
	return filepath.Join(app.HttpFolder(), "middleware")
}

func (app *AppService) CommandFolder() string {
	return filepath.Join(app.ConsoleFolder(), "command")
}

func (app *AppService) TestFolder() string {
	return filepath.Join(app.BaseFolder(), "test")
}

func NewAppSevice(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	return &AppService{
		container:  container,
		baseFolder: baseFolder,
	}, nil
}
