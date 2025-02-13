# Gray
## Todo/Ideas
- [ ] Command to copy the client into a project by providing the path.
## How to use?
- First start the server.
- Copy the gray folder into your porject. You may want to add it to your .gitignore to prevent it from being uploaded to your repository. This is because the gray folder is a development tool and not a part of your project. Also tests will fail if you didn't remove gray commands from your code.

To send something to the server, you can use the following code:
```go

import gray

func main() {
    gray.Send("Hello, World!")
}

```

Some examples of what you can send to the server are located in the examples folder. You can run them by using the following command:
```bash
go run examples/your_example.go
```

## Screenshot
![Screenshot](assets/Screenshot.png)
