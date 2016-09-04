package iface

type API interface {
	Post(path string, data interface{}) error
}
