package server

type Auth interface {
	CheckPasswd(Driver, string, string) (bool, error)
}
