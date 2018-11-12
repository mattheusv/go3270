package go3270

//NewEmulator return a instance of emulator struct
//host is the host that you connect
//port is the port of your host
//scriptPort is the port that you can run x3270 emulator - default is 5000
func NewEmulator(host string, port int, scriptPort string) Emulator {
	return Emulator{
		Host:       host,
		Port:       port,
		ScriptPort: scriptPort,
	}
}
