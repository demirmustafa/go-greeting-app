package greeting

func GetGreeting() interface{} {
	return &Greeting{Message:"Hello World!"}
}
