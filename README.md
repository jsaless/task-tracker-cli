# task-tracker-cli
A CLI to track all your task where you can add, delete, filter and so on

This project was based on https://roadmap.sh/projects/task-tracker

## How it works?

The CLI basically have the following functionalities:
- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

Here is an example of their commands:
```shell
# Adding a new task
task-cli -add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli -update 1 "Buy groceries and cook dinner"
task-cli -delete 1

# Marking a task as in progress or done
task-cli -mark-in-progress 1
task-cli -mark-done 1

# Listing all tasks
task-cli -list all

# Listing tasks by status
task-cli -list done
task-cli -list todo
task-cli -list in-progress
```

To make it easier to use the CLI you can create an alias to the executable file, otherwise you will have to use `./task-cli`, in the project folder, every time to use it.

> [!IMPORTANT]
> After each operation of `-add`, `-update`, `-delete` commands a `json` file will be created and updated according your actions.
