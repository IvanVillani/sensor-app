package flagsmanager

// IFlagsManager interface: defines method ParseFlags()
type IFlagsManager interface {
	ParseFlags() ([]string, error)
}
