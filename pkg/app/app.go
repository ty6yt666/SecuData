package app

type App struct {
	basename    string
	name        string
	description string
	//options				CliOptions
	runFunc   RunFunc
	silence   bool
	noVersion bool
	noConfig  bool
}

type RunFunc func(basename string) error

type Option func(*App)

func (a *App) Run() {}

func NewApp(name string, basename string, opts ...Option) {
	a := &App{
		name:     name,
		basename: basename,
	}

	for _, o := range opts {
		o(a)
	}

}
