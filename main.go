package main

func main() {
	tasks := Tasks{}

	storage := NewStorage("data.json")
	storage.Load(&tasks)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&tasks)

	storage.Save(&tasks)
}
